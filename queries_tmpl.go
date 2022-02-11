package gcg

var queriesTemplateSrc = `package {{.Package}}
// generated by github.com/emicklei/graphql-client-gen/cmd/gcg version: {{.BuildVersion }}
// DO NOT EDIT

import (
	"time"
)

var (
	_ = time.Now
)

const SchemaVersion = "{{.SchemaVersion}}"

{{- range .Queries}}

// {{.FunctionName}}Query is used for both specifying the query and capturing the response. {{.Comment}}
type {{.FunctionName}}Query struct {
	Errors Errors {{.ErrorsTag}}
	Data      {{if .IsArray}}[]{{end}}{{.FunctionName}}QueryData {{.DataTag}}
}

type {{.FunctionName}}QueryData struct {
	{{.ReturnType}} {{.ReturnFieldTag}}
}

// Build returns a GraphQLRequest with all the parts to send the HTTP request.
func (_q {{.FunctionName}}Query) Build(
	operationName string, // cannot be emtpy
	{{- range .Arguments}}
	{{.Name}} {{if .IsArray}}[]{{end}}{{.Type}}, 
	{{- end }}
) GraphQLRequest {
	_typedVars := map[string]valueAndType{
		{{- range .Arguments}}
		"{{.JSONName}}": {value:{{.Name}},graphType:"{{.GraphType}}"},
		{{- end }}
	}
	return buildRequest("query",operationName, _q.Data, _typedVars)
}

{{- end}}
`
