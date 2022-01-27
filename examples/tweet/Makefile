gen:
	rm -f scalars.go 
	rm -f enums.go 
	rm -f mutations.go
	rm -f types.go
	rm -f queries.go
	rm -f build.go
	rm -f unions.go
	go run ../../cmd/gcg/*.go -schema tweet.graphqls -pkg tweet
	gofmt -w enums.go
	gofmt -w mutations.go
	gofmt -w types.go
	gofmt -w scalars.go
	gofmt -w queries.go
	gofmt -w build.go
	gofmt -w unions.go