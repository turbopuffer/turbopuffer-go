// A simple example that lists all extant namespaces.
//
// Run this example with: go run examples/list_namespaces/list_namespaces.go

package main

import (
	"context"
	"fmt"

	"github.com/turbopuffer/turbopuffer-go"
	"github.com/turbopuffer/turbopuffer-go/option"
)

func main() {
	ctx := context.Background()
	client := turbopuffer.NewClient(option.WithRegion("gcp-us-central1"))

	namespaces := client.NamespacesAutoPaging(ctx, turbopuffer.NamespacesParams{})
	for namespaces.Next() {
		namespace := namespaces.Current()
		fmt.Printf("namespace: %s\n", namespace.ID)
	}
	if err := namespaces.Err(); err != nil {
		panic(err.Error())
	}
}
