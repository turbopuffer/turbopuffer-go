package turbopuffer

import (
	"context"
	"fmt"
	"slices"

	"github.com/turbopuffer/turbopuffer-go/v2/internal"
	"github.com/turbopuffer/turbopuffer-go/v2/internal/requestconfig"
	"github.com/turbopuffer/turbopuffer-go/v2/option"
)

// Represents a turbopuffer namespace.
type Namespace struct {
	NamespaceService
}

// Namespace creates a new namespace resource.
func (r *Client) Namespace(namespace string) Namespace {
	if len(r.Options) != cap(r.Options) {
		// If the options slice is not at its capacity, past versions of the Go
		// SDK would fail to create a copy below.
		r.shareCount.Add(1)
	}
	sendShareCount := requestconfig.RequestOptionFunc(func(c *requestconfig.RequestConfig) error {
		// Embed the share count in the User-Agent header to help determine
		// impact of past bug.
		shareCount := r.shareCount.Load()
		userAgent := fmt.Sprintf("Turbopuffer/Go %s (share-count=%d)", internal.PackageVersion, shareCount)
		c.Request.Header.Set("User-Agent", userAgent)
		return nil
	})
	// IMPORTANT: must append to a copy of r.Options. slices.Clip forces this
	// with only a single heap allocation. (Most other patterns for this result
	// in a clone operation and then a resize operation.)
	opts := slices.Clip(r.Options)
	opts = append(opts, option.WithDefaultNamespace(namespace), sendShareCount)
	return Namespace{NamespaceService: newNamespaceService(opts...)}
}

func (r *NamespaceService) ID() string {
	requestConfig, err := requestconfig.NewRequestConfig(context.Background(), "GET", "/", nil, nil, r.Options...)
	if err != nil {
		panic(fmt.Errorf("failed to create request config: %w", err))
	}
	return *requestConfig.DefaultNamespace
}
