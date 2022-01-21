package gcg

import (
	"fmt"
	"os"
	"text/template"
	"time"

	"github.com/vektah/gqlparser/ast"
)

func (g *Generator) handleTypes(doc *ast.SchemaDocument) error {
	out, err := os.Create("types.go")
	if err != nil {
		return err
	}
	defer out.Close()
	tmpl, err := template.New("tt").Parse(typeTemplateSrc)
	if err != nil {
		return err
	}
	fd := FileData{
		Package: g.packageName,
		Created: time.Now(),
	}
	for _, each := range doc.Definitions {
		if each.Name == "Mutation" {
			continue
		}
		if each.Kind == ast.Enum {
			continue
		}
		if each.Kind == ast.Object || each.Kind == ast.InputObject || each.Kind == ast.Interface {
			td := TypeData{
				Kind:          string(each.Kind),
				EmbeddedTypes: each.Interfaces,
				Name:          each.Name,
			}
			for _, other := range each.Fields {
				td.Fields = append(td.Fields, FieldData{
					Optional: !other.Type.NonNull,
					Name:     fieldName(other.Name),
					Type:     other.Type.Name(),
					IsArray:  isArray(other.Type),
					Tag:      fmt.Sprintf("`graphql:\"%s\" json:\"%s\"`", other.Name, other.Name),
				})
			}
			fd.Types = append(fd.Types, td)
		}
	}
	return tmpl.Execute(out, fd)
}
