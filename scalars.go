package gcg

import (
	"os"
	"text/template"
	"time"

	"github.com/vektah/gqlparser/ast"
)

func (g *Generator) handleScalars(all []*ast.Definition) error {
	out, err := os.Create("scalars.go")
	if err != nil {
		return err
	}
	defer out.Close()
	fd := FileData{
		Package: g.packageName,
		Created: time.Now(),
	}
	tmpl, err := template.New("tt").Parse(scalarsTemplateSrc)
	if err != nil {
		return err
	}
	for _, each := range all {
		sd := ScalarData{
			Comment: each.Description,
			Name:    each.Name,
		}
		fd.Scalars = append(fd.Scalars, sd)
	}
	return tmpl.Execute(out, fd)
}
