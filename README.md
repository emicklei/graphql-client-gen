# gcg - graphql client generator

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
    	createTweet (
        	body: String
    	): Tweet	
	}

With generated Go code from this schema, you can write queries and mutations.

Create example:

	m := CreateTweetMutation{}
	// set non-zero value to mark the field as part of query
	m.Data.ID = "?"
	// build GraphQLRequest with query,operation and variables
	r := m.Build("hello")

Create query:

	mutation createTweet($body:String) {
		createTweet(body: $body) {
			id
		}
	}

Read example:

	q := TweetsQuery{}
	q.Data = []TweetsQueryData{
		{
			// set non-zero value to mark the field as part of query
			Tweet: Tweet{ID: "?"},
		},
	}
	// build GraphQLRequest with query,operation and variables
	r := q.Build("testTweets", 1, 0, "id", "desc")

Read query 

	query testTweets($limit:Int,$skip:Int,$sort_field:String,$sort_order:String) {
		Tweets(limit: $limit,skip: $skip,sort_field: $sort_field,sort_order: $sort_order) {
			id
		}
	}

## status

work in progress

## convert schema JSON to SDL

The `gcg` program requires a schema in SDL. If you need to convert it from JSON then you can use the npm module.

	npm i graphql-json-to-sdl
	npx graphql-json-to-sdl schema.json schema.graphqls

### todo
 
- make optional arguments for function
- __typename meta field

## limitations

Unsupported GraphQL features are:

- inline fragments
- directives
- default values

© 2022+, [ernestmicklei.com](http://ernestmicklei.com). MIT License. Contributions welcome.