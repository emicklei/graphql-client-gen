package gcg

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"github.com/iancoleman/strcase"
	"github.com/vektah/gqlparser/ast"
)

func (g *Generator) handleQueries(def *ast.Definition) error {
	out, err := os.Create("queries.go")
	if err != nil {
		return err
	}
	defer out.Close()
	fd := FileData{
		Package:       g.packageName,
		BuildVersion:  g.mainVersion,
		SchemaVersion: g.schemaVersion,
	}
	tmpl, err := template.New("tt").Parse(queriesTemplateSrc)
	if err != nil {
		return err
	}
	for _, each := range def.Fields {
		op := OperationData{
			Comment:      formatComment(each.Description),
			FunctionName: strcase.ToCamel(each.Name),
			ReturnType:   g.mapScalar(each.Type.Name()),
			IsArray:      isArray(each.Type), // refers to the Data field
			ErrorsTag:    "`json:\"errors\"`",
		}
		// build return field tag
		// `graphql:"Tweet(id: $id)" json:"Tweet"`
		tag := new(bytes.Buffer)
		fmt.Fprintf(tag, "`graphql:\"%s", each.Name)
		if len(each.Arguments) > 0 {
			fmt.Fprint(tag, "(")
			for i, arg := range each.Arguments {
				if i > 0 {
					fmt.Fprintf(tag, ",")
				}
				fmt.Fprintf(tag, "%s: $%s", arg.Name, arg.Name)
				op.Arguments = append(op.Arguments, Argument{
					Name:      goArgName(arg.Name),
					JSONName:  arg.Name,
					Type:      g.mapScalar(arg.Type.Name()),
					GraphType: arg.Type.String(),
					IsArray:   isArray(arg.Type)})
			}
			fmt.Fprintf(tag, ")")
		}
		fmt.Fprintf(tag, "\" json:\"%s\"`", each.Name)
		op.ReturnFieldTag = tag.String()
		op.DataTag = "`graphql:\"query\"`"
		fd.Queries = append(fd.Queries, op)
	}
	return tmpl.Execute(out, fd)
}
