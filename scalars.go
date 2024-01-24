package gcg

import (
	"os"
	"path"
	"text/template"

	"github.com/vektah/gqlparser/ast"
)

func (g *Generator) handleScalars(all []*ast.Definition) error {
	out, err := os.Create(path.Join(g.targetDirectory, "scalars.go"))
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
