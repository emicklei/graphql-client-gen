package tests

import "testing"

func TestBuildFilterOpMutation(t *testing.T) {
	m := FilterOpMutation{}
	m.Data = append(m.Data, struct {
		Result "graphql:\"filterOp(sort: $sort)\" json:\"filterOp\""
	}{
		Result{ID: "?", GraphQLTypename: "?"},
	})
	req := m.Build("srt")
	t.Log(req.Query)
	t.Log(req.OperationName)
	t.Log(req.Variables)
}
