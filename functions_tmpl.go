package gcg

var functionsTemplateSrc = `package {{.Package}}
// generated by github.com/emicklei/graphql-client-gen/cmd/gcg version: {{.BuildVersion }}
// DO NOT EDIT

import (
	"encoding/json"
	"time"
)

var (
	_ = time.Now
	_ = json.Unmarshal
)

{{- range .Functions}}

// {{.Comment}}
type {{.Name}} struct { 
	{{- range .Fields}}
	{{- if gt (len .Comment) 0}}
	// {{.Comment}}{{- end}}
	{{.Name}} {{if .Optional}} *{{else}} {{end}}{{if .IsArray}}[]{{end}}{{.Type}} {{.Tag}}
	{{- end}} 
	// Result captures the query response part of this function.
	Result {{if .IsArray}}[]{{end}}{{.ReturnType}} {{.ResultTag}}
}

func (f *{{.Name}}) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &f.Result)
}
{{- end}}
`
