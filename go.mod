module github.com/turbopuffer/turbopuffer-go

go 1.22

require (
	github.com/tidwall/gjson v1.18.0
	github.com/tidwall/pretty v1.2.1
	github.com/tidwall/sjson v1.2.5
)

require github.com/tidwall/match v1.1.1 // indirect

// Versions before v1.12.2 contain a data race in Client.Namespace.
// See https://github.com/turbopuffer/turbopuffer-go/pull/88.
retract [v0.1.0-alpha.1, v1.12.1]
