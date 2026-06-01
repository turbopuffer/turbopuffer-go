// Package respondasync implements transparent polling for tpuf APIs that
// accept `prefer: respond-async`.
//
// Every API request is stamped with `prefer: respond-async`. If the server
// applies the preference (i.e. responds with `202 Accepted` +
// `preference-applied: respond-async`) the SDK polls the operation URL until
// it finishes and returns the final result as if the call had been
// synchronous.
package respondasync

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"slices"
	"strings"
	"time"

	"github.com/turbopuffer/turbopuffer-go/v2/internal/requestconfig"
	"github.com/turbopuffer/turbopuffer-go/v2/option"
)

const (
	headerPrefer            = "Prefer"
	headerPreferenceApplied = "Preference-Applied"
	headerLocation          = "Location"
	respondAsync            = "respond-async"
)

var PollInterval = 1 * time.Second

func NewOption(clientOpts []option.RequestOption) option.RequestOption {
	m := &middleware{clientOpts: slices.Clone(clientOpts)}
	return option.WithMiddleware(m.handle)
}

type middleware struct {
	clientOpts []option.RequestOption
}

func (m *middleware) handle(req *http.Request, next option.MiddlewareNext) (*http.Response, error) {
	// Don't overwrite an existing `prefer` header. This allows the caller
	// to disable async mode.
	if _, ok := req.Header[http.CanonicalHeaderKey(headerPrefer)]; !ok {
		req.Header.Set(headerPrefer, respondAsync)
	}

	res, err := next(req)
	if err != nil {
		return res, err
	}
	return m.maybePoll(req, res)
}

func (m *middleware) maybePoll(origReq *http.Request, res *http.Response) (*http.Response, error) {
	if !respondAsyncApplied(res) {
		return res, nil
	}

	// Drain and close the body, to prevent connection leaking.
	if res.Body != nil {
		_, _ = io.Copy(io.Discard, res.Body)
		_ = res.Body.Close()
	}

	location, err := extractLocation(res, origReq.URL)
	if err != nil {
		return res, err
	}

	ctx := origReq.Context()
	for {
		body, err := m.pollOnce(ctx, location)
		if err != nil {
			return res, err
		}
		finalRes, err := resolvePollResponse(body)
		if err != nil {
			return res, err
		}
		if finalRes != nil {
			return finalRes, nil
		}

		select {
		case <-ctx.Done():
			return res, ctx.Err()
		case <-time.After(PollInterval):
		}
	}
}

func extractLocation(res *http.Response, origURL *url.URL) (string, error) {
	rawLocation := res.Header.Get(headerLocation)
	if rawLocation == "" {
		return "", errors.New("server returned async response without a 'Location' header")
	}

	// Resolve the Location against the original request URL.
	resolved, err := origURL.Parse(rawLocation)
	if err != nil {
		return "", fmt.Errorf("malformed 'Location' header: %q", rawLocation)
	}

	// Reject a Location pointing at a different origin, to prevent API key exfiltration.
	if resolved.Scheme != origURL.Scheme || resolved.Host != origURL.Host {
		return "", fmt.Errorf("'Location' origin does not match request origin: %q", rawLocation)
	}

	return resolved.String(), nil
}

func respondAsyncApplied(res *http.Response) bool {
	if res == nil || res.StatusCode != http.StatusAccepted {
		return false
	}
	return strings.EqualFold(res.Header.Get(headerPreferenceApplied), respondAsync)
}

func (m *middleware) pollOnce(ctx context.Context, location string) ([]byte, error) {
	var res *http.Response
	opts := append(slices.Clone(m.clientOpts), option.WithResponseBodyInto(&res))
	if err := requestconfig.ExecuteNewRequest(ctx, http.MethodGet, location, nil, nil, opts...); err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return io.ReadAll(res.Body)
}

type pollBody struct {
	Status string      `json:"status"`
	Result *pollResult `json:"result"`
}

type pollResult struct {
	Success json.RawMessage `json:"success"`
	Error   *pollError      `json:"error"`
}

type pollError struct {
	StatusCode int             `json:"status_code"`
	Detail     json.RawMessage `json:"detail"`
}

var errMalformedPollResponse = errors.New("malformed poll response")

func resolvePollResponse(body []byte) (*http.Response, error) {
	var parsed pollBody

	if err := json.Unmarshal(body, &parsed); err != nil {
		return nil, errMalformedPollResponse
	}
	if parsed.Status == "running" {
		return nil, nil
	}
	if parsed.Status != "finished" || parsed.Result == nil {
		return nil, errMalformedPollResponse
	}

	var statusCode int
	var responseBody []byte
	switch {
	case parsed.Result.Success != nil:
		statusCode = http.StatusOK
		responseBody = parsed.Result.Success
	case parsed.Result.Error != nil:
		if parsed.Result.Error.StatusCode == 0 {
			return nil, errMalformedPollResponse
		}
		statusCode = parsed.Result.Error.StatusCode
		responseBody = parsed.Result.Error.Detail
	default:
		return nil, errMalformedPollResponse
	}

	return &http.Response{
		Status:     fmt.Sprintf("%d %s", statusCode, http.StatusText(statusCode)),
		StatusCode: statusCode,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header: http.Header{
			"Content-Type": {"application/json"},
			// Prevent the outer retry loop from starting another
			// async operation.
			"X-Should-Retry": {"false"},
		},
		Body: io.NopCloser(bytes.NewReader(responseBody)),
	}, nil
}
