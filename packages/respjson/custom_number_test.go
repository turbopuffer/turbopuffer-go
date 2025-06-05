package respjson_test

import (
	"encoding/json"
	"testing"

	"github.com/turbopuffer/turbopuffer-go/packages/respjson"
)

func TestNumber(t *testing.T) {
	// Use a number that's too large to fit precisely in a float64
	largeNumber := int64(9007199254740993)
	largeNumberStr := "9007199254740993"

	// Test creation.
	num := respjson.Number(largeNumberStr)
	if num.String() != largeNumberStr {
		t.Errorf("expected %s, got %s", largeNumberStr, num.String())
	}

	// Test int parsing.
	if numInt, err := num.Int64(); err != nil {
		t.Fatalf("failed to parse number: %v", err)
	} else if numInt != largeNumber {
		t.Errorf("expected %d, got %d", largeNumber, numInt)
	}

	// Test marshaling.
	if jsonBytes, err := json.Marshal(num); err != nil {
		t.Fatalf("failed to marshal number: %v", err)
	} else if string(jsonBytes) != largeNumberStr {
		t.Errorf("expected %s, got %s", largeNumberStr, string(jsonBytes))
	}

	// Test unmarshaling.
	var numUnmarshaled respjson.Number
	if err := json.Unmarshal([]byte(largeNumberStr), &numUnmarshaled); err != nil {
		t.Fatalf("failed to unmarshal number: %v", err)
	} else if numUnmarshaled != num {
		t.Errorf("expected %s, got %s", num, numUnmarshaled)
	}
}
