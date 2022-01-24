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
	Name    string  `graphql:"name"`
	Array1  []Child // no tag
	Array2  []Child `graphql:"array2"`
	Shared  `graphql:"shared()"`
	NoValue int
}

func TestBuildQueryRoot(t *testing.T) {
	r := Root{
		Name: "?",
		Array2: []Child{
			{ID: "?"},
		},
		Shared: Shared{Valid: true},
	}
	q := BuildQuery(r)
	sani := strings.ReplaceAll(strings.ReplaceAll(q, "\n", " "), "\t", "")
	if got, want := sani, "name array2 { id } shared() { valid } "; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
}
