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
		if each.Name == "Query" {
			continue
		}
		if each.Kind == ast.Enum {
			continue
		}
		if each.Kind == ast.Object || each.Kind == ast.InputObject || each.Kind == ast.Interface {
			td := TypeData{
				Comment: each.Description,
				Kind:    string(each.Kind),
				Name:    each.Name,
			}
			for _, other := range each.Fields {
				// is direct field or query
				if len(other.Arguments) > 0 {
					functionType := each.Name + fieldName(other.Name) + "Function"
					fnc := Function{
						Type:       functionType,
						Arguments:  other.Arguments,
						IsArray:    isArray(other.Type),
						ReturnType: mapScalar(other.Type.Name()),
					}
					g.functions = append(g.functions, fnc)
					td.Fields = append(td.Fields, FieldData{
						Comment: other.Description,
						Name:    fieldName(other.Name),
						Type:    "*" + functionType, // result is optional so use pointer
						Tag:     fmt.Sprintf("`graphql:\"%s()\" json:\"-,inline\"`", other.Name),
					})
				} else {
					td.Fields = append(td.Fields, FieldData{
						Comment:  other.Description,
						Optional: !other.Type.NonNull,
						Name:     fieldName(other.Name),
						Type:     mapScalar(other.Type.Name()),
						IsArray:  isArray(other.Type),
						Tag:      fmt.Sprintf("`graphql:\"%s\" json:\"%s,omitempty\"`", other.Name, other.Name),
					})
				}
			}
			fd.Types = append(fd.Types, td)
		}
	}
	return tmpl.Execute(out, fd)
}
