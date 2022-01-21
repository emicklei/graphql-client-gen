# gcg - graphql client generator

Generated sources have no foreign dependencies.

## run

    go run github.com/emicklei/graphql-client-gen/cmd/gcg -schema test.graphqls -pkg main
	gofmt -w enums.go
	gofmt -w mutations.go
	gofmt -w types.go
	gofmt -w scalars.go

## status

experiment