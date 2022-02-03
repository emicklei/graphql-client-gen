package tweet

import (
	"testing"
)

type Child struct {
	ID string `graphql:"id"`
}
type Shared struct {
	Arg   int  `graphql-function-arg:"arg" graphql-function-type:"Int!"`
	Valid bool `graphql:"valid"`
}
type Root struct {
	Shared        `graphql:"shared"` // embedded
	Name          string             `graphql:"name"`
	Array1        []Child            // no tag
	Array2        []Child            `graphql:"array2"`
	FunctionField *ScalarFunction    `graphql:"title"`
	NoValue       int
}

type ScalarFunction struct {
	// input
	AsUppercase bool `graphql-function-arg:"asUppercase" graphql-function-type:"Boolean"`
	// output
	string
}

func TestBuildQueryRoot(t *testing.T) {
	r := Root{
		Name: "?",
		Array2: []Child{
			{ID: "?"},
		},
		Shared: Shared{Arg: 42, Valid: true},
		FunctionField: &ScalarFunction{
			AsUppercase: true,
		},
	}
	tv := map[string]valueAndType{}
	req := buildRequest("query", "op", r, tv)
	t.Log("\n", req.Query)
	t.Log("\n", req.Variables)
}
