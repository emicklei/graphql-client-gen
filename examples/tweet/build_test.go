package tweet

import (
	"strings"
	"testing"
)

type Child struct {
	ID string `graphql:"id"`
}
type Shared struct {
	Arg   int  `graphql-function-arg:"arg"`
	Valid bool `graphql:"valid"`
}
type Root struct {
	Shared  `graphql:"shared"` // embedded
	Name    string             `graphql:"name"`
	Array1  []Child            // no tag
	Array2  []Child            `graphql:"array2"`
	NoValue int
}

func TestBuildQueryRoot(t *testing.T) {
	r := Root{
		Name: "?",
		Array2: []Child{
			{ID: "?"},
		},
		Shared: Shared{Arg: 42, Valid: true},
	}
	q := BuildQuery(r)
	t.Log("\n", q)
	sani := strings.ReplaceAll(strings.ReplaceAll(q, "\n", " "), "\t", "")
	if got, want := sani, "shared(arg:42) { valid } name array2 { id } "; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
}
