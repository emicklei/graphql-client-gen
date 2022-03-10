package tests

import (
	"strings"
	"testing"
)

func TestExplanationQuery(t *testing.T) {
	q := AllResultsQuery{}
	q.Data = []AllResultsQueryData{
		{Result: Result{
			Explanation: &ResultExplanationField{
				Language: "en_us",
			},
		}},
	}
	before := CustomDate("1920-03-12")
	r := q.Build("operation", before)
	if got, want := tabless(r.Query), tabless(`query operation($before:Date,$language1:String!) {
	allResults(before: $before) {
		explanation(language:$language1)
	}
}`); got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
	if got, want := r.Variables["language1"], "en_us"; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
	if got, want := r.Variables["before"], before; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
}

func tabless(s string) string {
	return strings.ReplaceAll(s, "\t", "")
}
