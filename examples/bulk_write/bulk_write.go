// Upserts several large batches of documents and prints the total time taken.
//
// Run this example with: go run examples/bulk_write/bulk_write.go

package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/turbopuffer/turbopuffer-go"
	"github.com/turbopuffer/turbopuffer-go/option"
)

const (
	numBatches = 10
	vectorDim  = 1024

	// Batch size is tuned to produce 32MB payloads. This achieves the maximum
	// discount (50%) without being close to the limit (256MB).
	batchSize = 1500
)

func main() {
	ctx := context.Background()
	client := turbopuffer.NewClient(option.WithRegion("gcp-us-central1"))

	namespace := client.Namespace("turbopuffer-go-bulk-upsert-test")
	fmt.Printf("Operating on namespace: %s\n", namespace.ID())

	startTime := time.Now()

	for batch := 0; batch < numBatches; batch++ {
		batchStartTime := time.Now()

		documents := make([]turbopuffer.RowParam, 0, batchSize)
		for i := 0; i < batchSize; i++ {
			id := batch*batchSize + i
			vector := make([]float64, vectorDim)
			for j := 0; j < vectorDim; j++ {
				vector[j] = rand.Float64()
			}
			documents = append(documents, turbopuffer.RowParam{
				ID:     turbopuffer.IDParam{Int: turbopuffer.Int(int64(id))},
				Vector: turbopuffer.VectorParam{FloatArray: vector},
			})
		}

		res, err := namespace.Write(ctx, turbopuffer.NamespaceWriteParams{
			UpsertRows: documents,
		})
		if err != nil {
			panic(err)
		}
		fmt.Printf("Batch %d complete, status: %s, time: %.2f seconds\n", batch+1, res.Status, time.Since(batchStartTime).Seconds())
	}

	endTime := time.Now()
	totalTimeSeconds := endTime.Sub(startTime).Seconds()
	fmt.Printf("Total time: %.2f seconds\n", totalTimeSeconds)
}
