package tweet

import (
	"encoding/json"
	"testing"
)

// union FilterValue = ValueFilter | ValuesFilter | RangeFilter

type MyFilterValue struct {
	ValueFilter
	ValuesFilter
	RangeFilter
}

func TestFilterValue(t *testing.T) {
	data := `{
		"value": "a",
		"values": ["b"]
	}`
	f := MyFilterValue{}
	if err := json.Unmarshal([]byte(data), &f); err != nil {
		t.Error(err)
	}
	if got, want := f.ValueFilter.Value, "a"; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
	if got, want := f.ValuesFilter.Values[0], "b"; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
}
