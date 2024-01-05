package gcg

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"text/template"

	"github.com/iancoleman/strcase"
	"github.com/vektah/gqlparser/ast"
)

func (g *Generator) handleMutations(each *ast.Definition) error {
	log.Println("generating mutations.go")
	out, err := os.Create("mutations.go")
	if err != nil {
		return err
	}
	defer out.Close()
	fd := FileData{
		Package:      g.packageName,
		BuildVersion: g.mainVersion,
	}
	tmpl, err := template.New("tt").Parse(operationsTemplateSrc)
	if err != nil {
		return err
	}
	for _, other := range each.Fields {
		rt := other.Type.Name() // g.mapScalar(other.Type.Name())
		od := OperationData{
			Comment:         formatComment(other.Description),
			Definition:      fieldDefinitionFor(other),
			Name:            other.Name,
			FunctionName:    strcase.ToCamel(other.Name),
			IsArray:         isArray(other.Type),
			ReturnFieldName: rt,
			ErrorsTag:       "`json:\"errors\"`",
		}
		// build return field tag
		// `graphql:"createGrouping(input:$input,pritSheetID:$pritSheetID,repositoryID:$repositoryID)" json:"createGrouping"`
		tag := new(bytes.Buffer)
		fmt.Fprintf(tag, "`graphql:\"%s(", other.Name)
		for i, arg := range other.Arguments {
			if i > 0 {
				fmt.Fprintf(tag, ",")
			}
			fmt.Fprintf(tag, "%s: $%s", arg.Name, arg.Name)
			od.Arguments = append(od.Arguments, Argument{
				Name:     goArgName(arg.Name),
				JSONName: arg.Name,
				Type:     g.mapScalar(arg.Type.Name()), IsArray: isArray(arg.Type),
				GraphType: arg.Type.String()})
		}
		fmt.Fprintf(tag, ")\" json:\"%s\"`", other.Name)
		od.ReturnFieldTag = tag.String()
		od.DataTag = "`graphql:\"mutation\"`"
		fd.Mutations = append(fd.Mutations, od)
	}
	return tmpl.Execute(out, fd)
}

func fieldDefinitionFor(f *ast.FieldDefinition) string {
	b := new(bytes.Buffer)
	io.WriteString(b, f.Name)
	io.WriteString(b, "(")
	for i, each := range f.Arguments {
		if i > 0 {
			io.WriteString(b, ",")
		}
		io.WriteString(b, each.Name)
		io.WriteString(b, ":")
		io.WriteString(b, each.Type.String())
	}
	io.WriteString(b, ")")
	io.WriteString(b, ":")
	io.WriteString(b, f.Type.String())
	return b.String()

}
