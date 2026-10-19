package option

import (
	"github.com/turbopuffer/turbopuffer-go/v2/internal/telemetry"
)

// WithTelemetry returns a RequestOption that enables or disables client telemetry, overriding the
// TURBOPUFFER_TELEMETRY environment variable. Telemetry is disabled by default.
//
// When enabled, the client reports its own request timings to turbopuffer to help diagnose
// latency between your application and the API. Timings of completed requests are attached to
// later requests, so no additional requests are made.
func WithTelemetry(enabled bool) RequestOption {
	return telemetry.Option(enabled)
}
