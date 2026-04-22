package turbopuffer

import (
	"encoding/json"
	"fmt"

	"github.com/tidwall/pretty"
)

// PrettyPrint pretty prints a JSON-encodable value to a string.
func PrettyPrint(v any) string {
	if err, ok := v.(error); ok {
		return fmt.Sprintf("<error: %s>", err.Error())
	}
	bytes, err := json.Marshal(v)
	if err != nil {
		return fmt.Sprintf("<error: %s>", err)
	}
	return string(pretty.Pretty(bytes))
}
