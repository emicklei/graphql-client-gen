package gcg

import (
	"fmt"
	"os"
	"text/template"

	"github.com/vektah/gqlparser/ast"
)

func (g *Generator) handleScalars(all []*ast.Definition) error {
	out, err := os.Create(fmt.Sprintf("%sscalars.go", g.outputFolder))
	if err != nil {
		return err
	}
	defer out.Close()
	fd := FileData{
		Package:      g.packageName,
		BuildVersion: g.mainVersion,
	}
	tmpl, err := template.New("tt").Parse(scalarsTemplateSrc)
	if err != nil {
		return err
	}
	// unless it is defined by a binding
	fd.IncludeScalarID = true
	for _, each := range g.scalarBindings {
		if each.GraphQLScalar == "ID" && each.GoTypeName != "interface{}" {
			fd.IncludeScalarID = false
			break
		}
	}
	for _, each := range all {
		sd := ScalarData{
			Comment: formatComment(each.Description),
			Name:    each.Name,
		}
		fd.Scalars = append(fd.Scalars, sd)
	}
	return tmpl.Execute(out, fd)
}
