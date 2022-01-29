package tweet

// generated by github.com/emicklei/graphql-client-gen/cmd/gcg version: (dev)
// DO NOT EDIT

import (
	"time"
)

var (
	_ = time.Now
)

// CreateTweetMutation is used for both specifying the query and capturing the response.
type CreateTweetMutation struct {
	Data struct {
		Tweet `graphql:"createTweet(body: $body)" json:"createTweet"`
	} `graphql:"mutation createTweet($body: String)"`
}

// Build returns a GraphQLRequest with all the parts to send the HTTP request.
func (_m CreateTweetMutation) Build(
	body string,
) GraphQLRequest {
	return GraphQLRequest{
		Query:         BuildQuery(_m),
		OperationName: "createTweet",
		Variables: map[string]interface{}{
			"body": body,
		},
	}
}

// DeleteTweetMutation is used for both specifying the query and capturing the response.
type DeleteTweetMutation struct {
	Data struct {
		Tweet `graphql:"deleteTweet(id: $id)" json:"deleteTweet"`
	} `graphql:"mutation deleteTweet($id: ID!)"`
}

// Build returns a GraphQLRequest with all the parts to send the HTTP request.
func (_m DeleteTweetMutation) Build(
	id interface{},
) GraphQLRequest {
	return GraphQLRequest{
		Query:         BuildQuery(_m),
		OperationName: "deleteTweet",
		Variables: map[string]interface{}{
			"id": id,
		},
	}
}

// MarkTweetReadMutation is used for both specifying the query and capturing the response.
type MarkTweetReadMutation struct {
	Data struct {
		bool `graphql:"markTweetRead(id: $id)" json:"markTweetRead"`
	} `graphql:"mutation markTweetRead($id: ID!)"`
}

// Build returns a GraphQLRequest with all the parts to send the HTTP request.
func (_m MarkTweetReadMutation) Build(
	id interface{},
) GraphQLRequest {
	return GraphQLRequest{
		Query:         BuildQuery(_m),
		OperationName: "markTweetRead",
		Variables: map[string]interface{}{
			"id": id,
		},
	}
}
