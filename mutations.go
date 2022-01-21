package gcg

import (
	"os"
	"text/template"
	"time"

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
		Package: g.packageName,
		Created: time.Now(),
	}
	tmpl, err := template.New("tt").Parse(operationsTemplateSrc)
	if err != nil {
		return err
	}
	for _, other := range each.Fields {
		od := OperationData{
			Name:         other.Name,
			FunctionName: strcase.ToCamel(other.Name),
			IsArray:      isArray(other.Type),
			ReturnType:   other.Type.Name(),
		}
		for _, arg := range other.Arguments {
			od.Arguments = append(od.Arguments, Argument{
				Name: arg.Name, Type: arg.Type.Name(), IsArray: isArray(arg.Type)})
		}
		fd.Operations = append(fd.Operations, od)
	}
	return tmpl.Execute(out, fd)
}
