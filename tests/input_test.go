package tests

import (
	"encoding/json"
	"testing"
)

func TestResultInput(t *testing.T) {
	i := ResultInput{}
	i.Unused(nil)
	i.When(CustomDate("2022-12-09"))
	data, err := json.Marshal(i)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := string(data), `{"unused":null,"when":"2022-12-09"}`; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
}
