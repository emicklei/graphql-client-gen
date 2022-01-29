package gcg

import (
	"os"
	"text/template"

	"github.com/vektah/gqlparser/ast"
)

/**
All of a union's included types
must be object types (not scalars, input types, etc.)
and that the included types do not need to share any fields.
**/

func (g *Generator) handleUnions(doc *ast.SchemaDocument, all []*ast.Definition) error {
	out, err := os.Create("unions.go")
	if err != nil {
		return err
	}
	defer out.Close()
	fd := FileData{
		Package:      g.packageName,
		BuildVersion: g.mainVersion,
	}
	tmpl, err := template.New("tt").Parse(unionTemplateSrc)
	if err != nil {
		return err
	}
	for _, each := range all {
		ud := UnionData{
			Comment: formatComment(each.Description),
			Name:    each.Name,
			Types:   each.Types,
		}
		fd.Unions = append(fd.Unions, ud)
	}
	return tmpl.Execute(out, fd)
}
