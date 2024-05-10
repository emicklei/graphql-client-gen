package gcg

import (
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/vektah/gqlparser/ast"
)

func (g *Generator) handleEnums(all []*ast.Definition) error {
	log.Println("generating enums.go")
	out, err := os.Create(fmt.Sprintf("%senums.go", g.outputFolder))
	if err != nil {
		return err
	}
	defer out.Close()
	fd := FileData{
		Package:      g.packageName,
		BuildVersion: g.mainVersion,
	}
	tmpl, err := template.New("tt").Parse(enumsTemplateSrc)
	if err != nil {
		return err
	}
	for _, each := range all {
		ed := EnumData{
			Comment: each.Description,
			Name:    each.Name,
		}
		for _, other := range each.EnumValues {
			ed.Values = append(ed.Values, other.Name)
		}
		fd.Enums = append(fd.Enums, ed)
	}
	return tmpl.Execute(out, fd)
}
