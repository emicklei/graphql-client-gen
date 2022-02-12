package tests

import (
	"strings"
	"testing"
	"time"
)

func TestExplanatioQuery(t *testing.T) {
	q := AllResultsQuery{}
	q.Data = []AllResultsQueryData{
		{Result: Result{
			Explanation: &ResultExplanationFunction{
				Language: "en_us",
			},
		}},
	}
	r := q.Build("operation", CustomDate(time.Now()))
	if got, want := tabless(r.Query), tabless(`query operation($before:Date,$language1:String!) {
	allResults(before: $before) {
		explanation(language:$language1)
	}
}`); got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
}

func tabless(s string) string {
	return strings.ReplaceAll(s, "\t", "")
}
