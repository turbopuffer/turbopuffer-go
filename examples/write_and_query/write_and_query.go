// A straightforward example of storing and retrieving documents via vector
// similarity search.
//
// Run this example with: go run examples/write_and_query/write_and_query.go

package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/turbopuffer/turbopuffer-go"
	"github.com/turbopuffer/turbopuffer-go/option"
)

func main() {
	ctx := context.Background()
	client := turbopuffer.NewClient(option.WithRegion("gcp-us-central1"))

	namespace := client.Namespace("turbopuffer-go-write-and-query-test")
	fmt.Printf("Operating on namespace: %s\n", namespace.ID())

	_, err := namespace.DeleteAll(ctx, turbopuffer.NamespaceDeleteAllParams{})
	if err != nil {
		var apiErr *turbopuffer.Error
		if errors.As(err, &apiErr) && apiErr.StatusCode == 404 {
			fmt.Println("Namespace not found, continuing")
		} else {
			panic(err)
		}
	}

	// Upsert some documents.
	{
		res, err := namespace.Write(ctx, turbopuffer.NamespaceWriteParams{
			UpsertRows: []turbopuffer.RowParam{
				{
					"id":     "b3ff34ea-87bb-469c-a854-9cb7e3713fc3",
					"vector": []float32{1.0, 2.0, 3.0},
					"name":   "Luke",
					"age":    32,
				},
				{
					"id":     "580d4471-9a9b-44fb-b59d-637ade604f72",
					"vector": []float32{4.0, 5.0, 6.0},
					"name":   "Leia",
					"age":    28,
				},
			},
			Schema: map[string]turbopuffer.AttributeSchemaConfigParam{
				"id":   {Type: turbopuffer.String("uuid")},
				"name": {Type: turbopuffer.String("string"), Filterable: turbopuffer.Bool(true)},
				"age":  {Type: turbopuffer.String("uint")},
			},
			DistanceMetric: turbopuffer.DistanceMetricCosineDistance,
		})
		if err != nil {
			panic(err)
		}
		fmt.Printf("Upsert rows affected: %d\n", res.RowsAffected)
	}

	// Do a vector query.
	{
		res, err := namespace.Query(ctx, turbopuffer.NamespaceQueryParams{
			RankBy:            turbopuffer.NewRankByVector("vector", []float32{3.0, 4.0, 5.0}),
			TopK:              turbopuffer.Int(10),
			IncludeAttributes: turbopuffer.IncludeAttributesParam{Bool: turbopuffer.Bool(true)},
			Filters: turbopuffer.NewFilterAnd([]turbopuffer.Filter{
				turbopuffer.NewFilterGt("age", 30),
				turbopuffer.NewFilterLt("age", 35),
			}),
		})
		if err != nil {
			panic(err)
		}
		fmt.Println("Query results:")
		for _, row := range res.Rows {
			fmt.Printf("    ID: %s, Name: %s, Age: %s\n", row["id"], row["name"], row["age"].(json.Number))
		}
	}

	// Print the schema.
	{
		res, err := namespace.Schema(ctx, turbopuffer.NamespaceSchemaParams{})
		if err != nil {
			panic(err)
		}
		fmt.Printf("Schema:\n%s\n", turbopuffer.PrettyPrint(res))
	}

	// Patch one document.
	{
		res, err := namespace.Write(ctx, turbopuffer.NamespaceWriteParams{
			PatchRows: []turbopuffer.RowParam{
				{
					"id":   "580d4471-9a9b-44fb-b59d-637ade604f72",
					"name": "Leia",
					"age":  82,
				},
			},
		})
		if err != nil {
			panic(err)
		}
		fmt.Printf("Patch rows affected: %d\n", res.RowsAffected)
	}

	// Do a non-vector query to see the patched results.
	{
		res, err := namespace.Query(ctx, turbopuffer.NamespaceQueryParams{
			RankBy:            turbopuffer.NewRankByAttribute("id", "asc"),
			TopK:              turbopuffer.Int(10),
			IncludeAttributes: turbopuffer.IncludeAttributesParam{Bool: turbopuffer.Bool(true)},
		})
		if err != nil {
			panic(err)
		}
		fmt.Println("Query results:")
		for _, row := range res.Rows {
			fmt.Printf("    ID: %s, Name: %s, Age: %s\n", row["id"], row["name"], row["age"].(json.Number))
		}
	}
}
