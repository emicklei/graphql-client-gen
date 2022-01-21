package gcg

import (
	"os"
	"text/template"
	"time"

	"github.com/vektah/gqlparser/ast"
)

func (g *Generator) handleEnums(all []*ast.Definition) error {
	out, err := os.Create("enums.go")
	if err != nil {
		return err
	}
	defer out.Close()
	fd := FileData{
		Package: g.packageName,
		Created: time.Now(),
	}
	tmpl, err := template.New("tt").Parse(enumsTemplateSrc)
	if err != nil {
		return err
	}
	for _, each := range all {
		ed := EnumData{
			Name: each.Name,
		}
		for _, other := range each.EnumValues {
			ed.Values = append(ed.Values, other.Name)
		}
		fd.Enums = append(fd.Enums, ed)
	}
	return tmpl.Execute(out, fd)
}
