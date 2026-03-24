package jsonpatchtomongo

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestParsePatches(t *testing.T) {
	want := `{
		"$push": {
			"hello.0.hi": {
				"$each": [
					{"$numberDouble": "4.0"},
					{"$numberDouble": "3.0"},
					{"$numberDouble": "2.0"},
					{"$numberDouble": "1.0"}
				],
				"$position": {"$numberInt": "5"}
			}
		},
		"$set": {
			"hello.0.hi.num": {"$numberDouble": "4.0"}
		}
	}`

	patches := []byte(`[
  		{ "op": "add", "path": "/hello/0/hi/5", "value": 1 },
  		{ "op": "add", "path": "/hello/0/hi/5", "value": 2 },
  		{ "op": "add", "path": "/hello/0/hi/5", "value": 3 },
  		{ "op": "add", "path": "/hello/0/hi/5", "value": 4 },
  		{ "op": "add", "path": "/hello/0/hi/num", "value": 4 }
	]`)
	val, _, err := ParsePatches(patches)
	valStr := fmt.Sprint(val)

	if err != nil || normalizeJSON(t, valStr) != normalizeJSON(t, want) {
		t.Errorf("ParsePatches() = %q, want %q", valStr, want)
	}
}

func normalizeJSON(t *testing.T, s string) string {
	t.Helper()

	var v any
	if err := json.Unmarshal([]byte(s), &v); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("re-marshal failed: %v", err)
	}
	return string(b)
}
