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

## status

work in progress

## convert schema JSON to SDL

The `gcg` requires a schema in SDL. If you need to convert it from JSON then you can use the npm module.

	npm i graphql-json-to-sdl
	npx graphql-json-to-sdl schema.json schema.graphqls

### todo
 
- make options for function
- __typename

## limitations

Unsupported GraphQL features are:

- inline fragments
- directives
- default values