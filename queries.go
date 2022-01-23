package gcg

import (
	"os"
	"text/template"
	"time"

	"github.com/vektah/gqlparser/ast"
)

func (g *Generator) handleQueries(each *ast.Definition) error {
	out, err := os.Create("queries.go")
	if err != nil {
		return err
	}
	defer out.Close()
	fd := FileData{
		Package: g.packageName,
		Created: time.Now(),
	}
	tmpl, err := template.New("tt").Parse(queriesTemplateSrc)
	if err != nil {
		return err
	}
	return tmpl.Execute(out, fd)
}
