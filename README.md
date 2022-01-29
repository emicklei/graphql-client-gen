# gcg - graphql client generator

Generated sources have no dependencies outside the standard Go SDK.

## run

    go run github.com/emicklei/graphql-client-gen/cmd/gcg -schema schema.graphqls -pkg main
	gofmt -w enums.go
	gofmt -w mutations.go
	gofmt -w types.go
	gofmt -w scalars.go
	gofmt -w queries.go
	gofmt -w unions.go
	gofmt -w build.go

## status

work in progress

# convert schema JSON to SDL

	npm i graphql-json-to-sdl
	npx graphql-json-to-sdl schema.json schema.graphqls

# todo

- embed source ipv tmp
- add errors
- make options for function