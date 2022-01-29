package tweet

import (
	"testing"
)

// Tweet(id: ID!): Tweet

type TweetQuery2 struct {
	// Operation is the operationName and cannot be empty
	Operation string
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
	id interface{},
) GraphQLRequest {
	_q.Operation = operationName
	return GraphQLRequest{
		Query:         BuildQuery(_q),
		OperationName: operationName,
		Variables: map[string]interface{}{
			"id": id,
		},
	}
}

func TestTweetQueryGen(t *testing.T) {
	q := TweetQuery{}
	q.Data.Tweet.Author = &User{Name: &Get.String}
	q.Data.Stats = &Stat{Likes: &Get.Int32}
	s := q.Build("test", 101)
	t.Log("\n", s)
}

func TestTweetQuery2Gen(t *testing.T) {
	q := TweetQuery2{}
	q.Data.Tweet.Author = &User{Name: &Get.String}
	q.Data.Stats = &Stat{Likes: &Get.Int32}
	s := q.Build("test", 101)
	t.Log("\n", s)
}
