package tweet

import (
	"testing"
)

// Tweet(id: ID!): Tweet

type TweetQuery2 struct {
	// Operation is the operationName and cannot be empty
	Operation string
	Errors2   `json:"errors,omitempty"`
	Data      TweetQuery2Data `graphql:"query OperationName($id: ID!)"`
}

type TweetQuery2Data struct {
	Tweet `graphql:"Tweet(id: $id)" json:"Tweet"`
}

// OperationName returns the actual query operation name that is used to replace "OperationName"
func (q TweetQuery2) OperationName() string {
	return q.Operation
}

func (_q TweetQuery2) Build(
	operationName string,
	args TweetQuery2Args,
) GraphQLRequest {
	_q.Operation = operationName
	return GraphQLRequest{
		Query:         BuildQuery(_q),
		OperationName: operationName,
		Variables: map[string]interface{}{
			"id": args.id,
		},
	}
}

type TweetQuery2Args struct {
	id interface{}
}

func TestTweetQuery2Gen(t *testing.T) {
	q := TweetQuery2{}
	q.Data.Tweet.Author = &User{Name: &Get.String}
	q.Data.Stats = &Stat{Likes: &Get.Int32}
	s := q.Build("test", TweetQuery2Args{id: 1})
	t.Log("\n", s)
}

type Errors2 struct {
	Message   string `json:"message,omitempty"`
	Locations []struct {
		Line   int `json:"line"`
		Column int `json:"column"`
	} `json:"locations,omitempty"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`
}
