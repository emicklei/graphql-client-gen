package gcg

import (
	"fmt"
	"os"
	"text/template"
	"time"

	"github.com/iancoleman/strcase"
	"github.com/vektah/gqlparser/ast"
	"github.com/vektah/gqlparser/parser"
)

type Generator struct {
	schemaSource string
	packageName  string
}

func NewGenerator(schemaSource string, options ...Option) *Generator {
	g := &Generator{schemaSource: schemaSource}
	for _, each := range options {
		each(g)
	}
	return g
}

func (g *Generator) Generate() error {
	doc, perr := parser.ParseSchema(&ast.Source{Input: g.schemaSource, Name: "test"})
	if perr != nil {
		return perr
	}
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
			for _, other := range each.Fields {
				od := OperationData{
					Name:         other.Name,
					FunctionName: strcase.ToCamel(other.Name),
					IsArray:      isArray(other.Type),
					ReturnType:   other.Type.Name(),
				}
				fd.Operations = append(fd.Operations, od)
			}
			continue
		}
		if each.Kind == ast.Enum {
			ed := EnumData{
				Name: each.Name,
			}
			for _, other := range each.EnumValues {
				ed.Values = append(ed.Values, other.Name)
			}
			fd.Enums = append(fd.Enums, ed)
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
					Tag:      fmt.Sprintf("`json:\"%s\"`", other.Name),
				})
			}
			fd.Types = append(fd.Types, td)
		}
	}
	return tmpl.Execute(os.Stdout, fd)
}

func fieldName(s string) string {
	if s == "id" {
		return "ID"
	}
	return strcase.ToCamel(s)
}

func isArray(t *ast.Type) bool {
	return t.NamedType == ""
}
