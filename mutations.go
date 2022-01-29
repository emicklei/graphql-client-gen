package gcg

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"github.com/iancoleman/strcase"
	"github.com/vektah/gqlparser/ast"
)

func (g *Generator) handleMutations(each *ast.Definition) error {
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
		rt := g.mapScalar(other.Type.Name())
		od := OperationData{
			Comment:      formatComment(other.Description),
			Name:         other.Name,
			FunctionName: strcase.ToCamel(other.Name),
			IsArray:      isArray(other.Type),
			ReturnType:   rt,
			ErrorsTag:    "`json:\"errors\"`",
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
				Type:     g.mapScalar(arg.Type.Name()), IsArray: isArray(arg.Type)})
		}
		fmt.Fprintf(tag, ")\" json:\"%s\"`", other.Name)
		od.ReturnFieldTag = tag.String()
		// build data field tag
		// `graphql:"mutation createGrouping($input: GroupingInput!, $pritSheetID: ID!, $repositoryID: ID!)"`
		tag = new(bytes.Buffer)
		fmt.Fprintf(tag, "`graphql:\"mutation %s(", other.Name)
		for i, arg := range other.Arguments {
			if i > 0 {
				fmt.Fprintf(tag, ",")
			}
			fmt.Fprintf(tag, "$%s: %s", arg.Name, arg.Type.String())
		}
		fmt.Fprintf(tag, ")\"`")
		od.DataTag = tag.String()
		fd.Mutations = append(fd.Mutations, od)
	}
	return tmpl.Execute(out, fd)
}
