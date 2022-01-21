package gcg

import (
	"encoding/json"
	"testing"
)

type MyEnum string

const MyEnum_VAL1 = MyEnum("VAL1")

func TestMyEnumJSON(t *testing.T) {
	data, _ := json.Marshal(MyEnum_VAL1)
	if got, want := string(data), `"VAL1"`; got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
	if got, want := MyEnum("VAL1") == MyEnum_VAL1, true; got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
}
