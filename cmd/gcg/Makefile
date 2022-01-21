run:
	rm -f scalars.go 
	rm -f enums.go 
	rm -f mutations.go
	rm -f types.go
	go run *.go -schema test.graphqls -pkg main
	gofmt -w enums.go
	gofmt -w mutations.go
	gofmt -w types.go
	gofmt -w scalars.go