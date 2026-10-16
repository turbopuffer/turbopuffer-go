// Package telemetry reports the client's request timings to turbopuffer on later requests.
//
// When enabled, the client buffers a small timing record for each request and attaches buffered
// records to a later request in the X-turbopuffer-Client-Telemetry header. The server controls how
// many requests are recorded by advertising a sample rate in the
// X-turbopuffer-Client-Telemetry-Sample-Rate response header.
package telemetry

import (
	"encoding/json"
	"log/slog"
	"math"
	"math/rand/v2"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

const (
	EnvVar           = "TURBOPUFFER_TELEMETRY"
	Header           = "X-turbopuffer-Client-Telemetry"
	SampleRateHeader = "X-turbopuffer-Client-Telemetry-Sample-Rate"

	version = 1

	// Keeps the header well under the 8 KB per-line limit of common proxies.
	maxRecordsPerHeader = 20
	maxBufferedRecords  = 1000
)

var (
	namespacePattern = regexp.MustCompile(`^(/v\d+/namespaces/)[^/?]+`)
	operationPattern = regexp.MustCompile(`(/operations/)[^/?]+`)
)

// Indirection so tests can control sampling.
var random = rand.Float64

// enabledByEnv reports whether TURBOPUFFER_TELEMETRY enables telemetry.
func enabledByEnv() bool {
	raw := os.Getenv(EnvVar)
	switch strings.ToLower(strings.TrimSpace(raw)) {
	case "1", "true", "yes", "on":
		return true
	case "", "0", "false", "no", "off":
		return false
	}
	slog.Warn("ignoring invalid "+EnvVar+", expected 1 or 0", "value", raw)
	return false
}

// routeTemplate returns the method and path of a request with identifiers like the namespace
// replaced.
func routeTemplate(method, rawURL string) string {
	path := rawURL
	if strings.Contains(path, "://") {
		if u, err := url.Parse(path); err == nil {
			path = u.RequestURI()
		}
	}
	path = namespacePattern.ReplaceAllString(path, "${1}{ns}")
	path = operationPattern.ReplaceAllString(path, "${1}{token}")
	return strings.ToUpper(method) + " " + path
}

// record is the telemetry of one request. Durations are in microseconds, and are nil when their
// phase wasn't observed.
type record struct {
	Route         string  `json:"route"`
	Status        int     `json:"status"`
	Error         string  `json:"error,omitempty"`
	Retries       int     `json:"retries,omitempty"`
	SampleRate    float64 `json:"sample_rate,omitempty"`
	TotalUs       *int64  `json:"total_us,omitempty"`
	SerializeUs   *int64  `json:"serialize_us,omitempty"`
	CompressUs    *int64  `json:"compress_us,omitempty"`
	PoolUs        *int64  `json:"pool_us,omitempty"`
	ConnectUs     *int64  `json:"connect_us,omitempty"`
	SendUs        *int64  `json:"send_us,omitempty"`
	WaitUs        *int64  `json:"wait_us,omitempty"`
	BodyReadUs    *int64  `json:"body_read_us,omitempty"`
	DeserializeUs *int64  `json:"deserialize_us,omitempty"`
	ServerTotalUs *int64  `json:"server_total_us,omitempty"`
	Reused        *bool   `json:"reused,omitempty"`
	// AgeMs is set when the record is taken to be sent.
	AgeMs int64 `json:"age_ms"`
}

type entry struct {
	recordedAt time.Time
	record     record
}

type envelope struct {
	Version int      `json:"version"`
	Records []record `json:"records"`
	Dropped int      `json:"dropped,omitempty"`
}

// batch is records taken from a buffer, encoded as a header value.
type batch struct {
	entries []entry
	dropped int
	header  string
}

func newBatch(entries []entry, dropped int) *batch {
	now := time.Now()
	records := make([]record, len(entries))
	for i, e := range entries {
		records[i] = e.record
		records[i].AgeMs = now.Sub(e.recordedAt).Round(time.Millisecond).Milliseconds()
	}
	// Can't fail, since records hold no infinite or NaN floats.
	header, _ := json.Marshal(envelope{Version: version, Records: records, Dropped: dropped})
	return &batch{entries: entries, dropped: dropped, header: string(header)}
}

// buffer holds telemetry records until a later request can carry them.
type buffer struct {
	mu      sync.Mutex
	entries []entry
	dropped int
	// sampleRate holds the bits of a float64, which is 1 until the server advertises a rate.
	sampleRate atomic.Uint64
}

func newBuffer() *buffer {
	b := &buffer{}
	b.sampleRate.Store(math.Float64bits(1))
	return b
}

func (b *buffer) rate() float64 {
	return math.Float64frombits(b.sampleRate.Load())
}

// observeResponse adopts the sample rate advertised by the server, ignoring invalid values.
func (b *buffer) observeResponse(header http.Header) {
	rate, err := strconv.ParseFloat(header.Get(SampleRateHeader), 64)
	// Also rejects NaN, since every comparison with NaN is false.
	if err == nil && rate >= 0 && rate <= 1 {
		b.sampleRate.Store(math.Float64bits(rate))
	}
}

// sample reports whether to record a request, and the rate it's sampled at.
func (b *buffer) sample() (float64, bool) {
	rate := b.rate()
	if rate <= 0 || (rate < 1 && random() >= rate) {
		return rate, false
	}
	return rate, true
}

// add buffers a record, dropping the oldest when full.
func (b *buffer) add(rec record) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if len(b.entries) >= maxBufferedRecords {
		b.entries = slices.Delete(b.entries, 0, 1)
		b.dropped++
	}
	b.entries = append(b.entries, entry{recordedAt: time.Now(), record: rec})
}

// take removes up to maxRecordsPerHeader of the oldest records.
func (b *buffer) take() *batch {
	b.mu.Lock()
	n := min(len(b.entries), maxRecordsPerHeader)
	if n == 0 {
		b.mu.Unlock()
		return nil
	}
	entries := slices.Clone(b.entries[:n])
	b.entries = slices.Delete(b.entries, 0, n)
	dropped := b.dropped
	b.dropped = 0
	b.mu.Unlock()
	return newBatch(entries, dropped)
}

// restore returns a batch that never reached the server to the front of the buffer.
func (b *buffer) restore(bt *batch) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.dropped += bt.dropped
	// The batch holds the oldest records, so they're the ones to drop when full.
	keep := min(len(bt.entries), maxBufferedRecords-len(b.entries))
	b.dropped += len(bt.entries) - keep
	b.entries = slices.Insert(b.entries, 0, bt.entries[len(bt.entries)-keep:]...)
}

type bufferKey struct {
	apiKey  string
	baseURL string
}

// sharedBuffers maps a bufferKey to its buffer. Buffers are never removed, but an idle one only
// holds the records of its last few requests.
var sharedBuffers sync.Map

// sharedBuffer returns the buffer for an API key and base URL. Sharing lets records ride on any
// request, but never to another org or region.
func sharedBuffer(apiKey, baseURL string) *buffer {
	key := bufferKey{apiKey: apiKey, baseURL: baseURL}
	if b, ok := sharedBuffers.Load(key); ok {
		return b.(*buffer)
	}
	b, _ := sharedBuffers.LoadOrStore(key, newBuffer())
	return b.(*buffer)
}
