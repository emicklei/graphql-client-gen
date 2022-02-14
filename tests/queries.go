package tests

// generated by github.com/emicklei/graphql-client-gen/cmd/gcg version: (devel)
// DO NOT EDIT

import (
	"time"
)

var (
	_ = time.Now
)

const SchemaVersion = "v1.0.0"

// NoArgOpQuery is used for both specifying the query and capturing the response.
type NoArgOpQuery struct {
	Errors Errors           `json:"errors"`
	Data   NoArgOpQueryData `graphql:"query"`
}

type NoArgOpQueryData struct {
	int32 `graphql:"noArgOp" json:"noArgOp"`
}

// Build returns a GraphQLRequest with all the parts to send the HTTP request.
func (_q NoArgOpQuery) Build(
	operationName string, // cannot be emtpy
) GraphQLRequest {
	_typedVars := map[string]valueAndType{}
	return buildRequest("query", operationName, _q.Data, _typedVars)
}

// OneArgOpQuery is used for both specifying the query and capturing the response.
type OneArgOpQuery struct {
	Errors Errors            `json:"errors"`
	Data   OneArgOpQueryData `graphql:"query"`
}

type OneArgOpQueryData struct {
	string `graphql:"oneArgOp(required: $required)" json:"oneArgOp"`
}

// Build returns a GraphQLRequest with all the parts to send the HTTP request.
func (_q OneArgOpQuery) Build(
	operationName string, // cannot be emtpy
	_required bool,
) GraphQLRequest {
	_typedVars := map[string]valueAndType{
		"required": {value: _required, graphType: "Boolean!"},
	}
	return buildRequest("query", operationName, _q.Data, _typedVars)
}

// FilterOpQuery is used for both specifying the query and capturing the response.
type FilterOpQuery struct {
	Errors Errors              `json:"errors"`
	Data   []FilterOpQueryData `graphql:"query"`
}

type FilterOpQueryData struct {
	Result `graphql:"filterOp(sort: $sort)" json:"filterOp"`
}

// Build returns a GraphQLRequest with all the parts to send the HTTP request.
func (_q FilterOpQuery) Build(
	operationName string, // cannot be emtpy
	_sort string,
) GraphQLRequest {
	_typedVars := map[string]valueAndType{
		"sort": {value: _sort, graphType: "String!"},
	}
	return buildRequest("query", operationName, _q.Data, _typedVars)
}

// ListOpQuery is used for both specifying the query and capturing the response.
type ListOpQuery struct {
	Errors Errors            `json:"errors"`
	Data   []ListOpQueryData `graphql:"query"`
}

type ListOpQueryData struct {
	Result `graphql:"ListOp(limit: $limit,prefix: $prefix)" json:"ListOp"`
}

// Build returns a GraphQLRequest with all the parts to send the HTTP request.
func (_q ListOpQuery) Build(
	operationName string, // cannot be emtpy
	_limit int32,
	_prefix string,
) GraphQLRequest {
	_typedVars := map[string]valueAndType{
		"limit":  {value: _limit, graphType: "Int"},
		"prefix": {value: _prefix, graphType: "String!"},
	}
	return buildRequest("query", operationName, _q.Data, _typedVars)
}

// PlusOpQuery is used for both specifying the query and capturing the response.
type PlusOpQuery struct {
	Errors Errors          `json:"errors"`
	Data   PlusOpQueryData `graphql:"query"`
}

type PlusOpQueryData struct {
	int32 `graphql:"plusOp(a: $a,b: $b)" json:"plusOp"`
}

// Build returns a GraphQLRequest with all the parts to send the HTTP request.
func (_q PlusOpQuery) Build(
	operationName string, // cannot be emtpy
	_a int32,
	_b int32,
) GraphQLRequest {
	_typedVars := map[string]valueAndType{
		"a": {value: _a, graphType: "Int!"},
		"b": {value: _b, graphType: "Int!"},
	}
	return buildRequest("query", operationName, _q.Data, _typedVars)
}

// PlusArrayOpQuery is used for both specifying the query and capturing the response.
type PlusArrayOpQuery struct {
	Errors Errors                 `json:"errors"`
	Data   []PlusArrayOpQueryData `graphql:"query"`
}

type PlusArrayOpQueryData struct {
	int32 `graphql:"plusArrayOp(as: $as,bs: $bs)" json:"plusArrayOp"`
}

// Build returns a GraphQLRequest with all the parts to send the HTTP request.
func (_q PlusArrayOpQuery) Build(
	operationName string, // cannot be emtpy
	_as []int32,
	_bs []int32,
) GraphQLRequest {
	_typedVars := map[string]valueAndType{
		"as": {value: _as, graphType: "[Int]!"},
		"bs": {value: _bs, graphType: "[Int!]"},
	}
	return buildRequest("query", operationName, _q.Data, _typedVars)
}

// AllResultsQuery is used for both specifying the query and capturing the response.
type AllResultsQuery struct {
	Errors Errors                `json:"errors"`
	Data   []AllResultsQueryData `graphql:"query"`
}

type AllResultsQueryData struct {
	Result `graphql:"allResults(before: $before)" json:"allResults"`
}

// Build returns a GraphQLRequest with all the parts to send the HTTP request.
func (_q AllResultsQuery) Build(
	operationName string, // cannot be emtpy
	_before CustomDate,
) GraphQLRequest {
	_typedVars := map[string]valueAndType{
		"before": {value: _before, graphType: "Date"},
	}
	return buildRequest("query", operationName, _q.Data, _typedVars)
}