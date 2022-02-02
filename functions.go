package gcg

import (
	"fmt"
	"os"
	"text/template"

	"github.com/vektah/gqlparser/ast"
)

type Function struct {
	Signature   string
	Type        string
	Description string
	Arguments   ast.ArgumentDefinitionList
	IsArray     bool
	ReturnType  string
	Tag         string
}

func (g *Generator) handleFunctions() error {
	out, err := os.Create("functions.go")
	if err != nil {
		return err
	}
	defer out.Close()
	fd := FileData{
		Package:      g.packageName,
		BuildVersion: g.mainVersion,
	}
	tmpl, err := template.New("tt").Parse(functionsTemplateSrc)
	if err != nil {
		return err
	}
	for _, each := range g.functions {
		fnd := FunctionData{
			Comment:    fmt.Sprintf("%s\n// %s", formatComment(each.Description), each.Signature),
			Name:       each.Type,
			IsArray:    each.IsArray,
			ReturnType: each.ReturnType,
			ResultTag:  "`graphql:\"inline\"`",
		}
		for _, other := range each.Arguments {
			// todo refactor this, is now copy of types.go occurrence
			fnd.Fields = append(fnd.Fields, FieldData{
				Comment:  formatComment(other.Description),
				Optional: !other.Type.NonNull,
				Name:     fieldName(other.Name),
				Type:     g.mapScalar(other.Type.Name()),
				IsArray:  isArray(other.Type),
				Tag:      fmt.Sprintf("`graphql-function-arg:\"%s\" graphql-function-type:\"%s\"`", other.Name, other.Type.String()),
			})
		}
		fd.Functions = append(fd.Functions, fnd)
	}
	return tmpl.Execute(out, fd)
}
