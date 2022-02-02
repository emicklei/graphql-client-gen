package gcg

import (
	"encoding/json"
	"testing"
)

type FunctionWithArg struct {
	Arg      int `json:"arg"`
	Returned bool
}

func (f *FunctionWithArg) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &f.Returned)
}
func TestBuildFunctionWithBoolReturn(t *testing.T) {
	type FBU struct {
		FieldFunction *FunctionWithArg `json:"field"`
	}
	fbu := new(FBU)
	data := `{"field":true}`
	if err := json.Unmarshal([]byte(data), fbu); err != nil {
		t.Error(err)
	}
	if got, want := fbu.FieldFunction.Returned, true; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
}
