![Travis (.com)](https://img.shields.io/travis/com/emicklei/graphql-client-gen)
![GitHub](https://img.shields.io/github/license/emicklei/graphql-client-gen)
![OSS Lifecycle](https://img.shields.io/osslifecycle/emicklei/graphql-client-gen)

# gcg - graphql client generator


This tool takes a GraphQL schema in SDL and generates Go sources for all entities (enums, mutations, types, queries, unions, functions).

The generated types can be used both for composing GraphQL queries and populating response data.
Queries and mutations are composed by setting field values in your (nested) structure to a non-zero value.
If you need more complex queries than can not be expressed using the types then you can choose to use the types only for capturing the response.

Generated sources have no dependencies outside the standard Go SDK.

## prepare

Create a configuration file `schema-generate.yaml` with the following contents:

	# the Go package name for the generated files
	package: tweet

	# the GraphQL schema in SDL (use a converter if you have JSON)
	schema: schema.gql
	
	# optionally, map Scalar to your own implementation
	bindings:
  		Date: CustomDate
	
	# optional: provide a folder to output the generated files to.
	# note - this folder must already exist, path is relative to config
	# file location
	# 
	# trailing slash is required for formatting
	output_folder: pkg/generated/

## run

    go run github.com/emicklei/graphql-client-gen/cmd/gcg
	gofmt -w enums.go
	gofmt -w mutations.go
	gofmt -w types.go
	gofmt -w scalars.go
	gofmt -w queries.go
	gofmt -w unions.go
	gofmt -w functions.go
	gofmt -w build.go

## usage

Schema example (SDL):

	type Tweet {
		id: ID!
		body: String
		Responders(limit:Int!):[User!]
	}
	type Query {
    	Tweets(limit: Int, skip: Int, sort_field: String, sort_order: String): [Tweet]
	}
	type Mutation {
    	createTweet (body: String): Tweet	
	}

With generated Go code from this schema, you can write queries and mutations.

Create example:

	m := CreateTweetMutation{}

	// set non-zero value to mark the field as part of query
	m.Data.ID = "?"

	// build GraphQLRequest with query,operation and variables
	r := m.Build("hello")

Results in query:

	mutation createTweet($body:String) {
		createTweet(body: $body) {
			id
		}
	}

Read Tweet ID from response

	// use the CreateTweetMutation to capture the data
	json.Unmarshal(responseBytes, &m)
	id := m.Data.ID

Read example:

	q := TweetsQuery{}
	q.Data = []TweetsQueryData{
		{
			Tweet: Tweet{ID: "?"},
		},
	}
	r := q.Build("testTweets", 1, 0, "id", "desc")

Results in query 

	query testTweets($limit:Int,$skip:Int,$sort_field:String,$sort_order:String) {
		Tweets(limit: $limit,skip: $skip,sort_field: $sort_field,sort_order: $sort_order) {
			id
		}
	}

Read Tweet ID from response

	// use the TweetsQuery to capture the data
	json.Unmarshal(responseBytes, &q)
	id := q.Data.Tweets[0].ID

## how to post a GraphQLRequest

	request := aQueryOrMutation.Build(....)
	requestBytes, err := json.Marshal(request)
	requestReader := bytes.NewReader(requestBytes)

	// if you need to pass headers then use http.NewRequest instead
	resp , err = http.Post("http://your.service/api", "application/json", requestReader)

## inject the schema version

Using a directive, you can provide version information about the schema.
The following example shows how to define this

	directive @version(name:String="dev") on SCHEMA

	schema @version(name:"v1.0.0") {
		query: Query
		mutation: Mutation
	}

Then the generator will initialize a constant in queries.go such as:

	const SchemaVersion = "v1.0.0"

## limitations

Unsupported GraphQL features are:

- inline fragments
- directives
- default values

© 2022+, [ernestmicklei.com](http://ernestmicklei.com). MIT License. Contributions welcome.
