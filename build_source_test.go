package gcg

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func ExampleGraphQLRequest() {
	request := GraphQLRequest{}
	requestBytes, _ := json.Marshal(request)
	requestReader := bytes.NewReader(requestBytes)
	_, _ = http.Post("http://your.service/api", "application/json", requestReader)
}
