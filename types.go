package gcg

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"text/template"

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
		Package:      g.packageName,
		BuildVersion: g.mainVersion,
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
		if each.Kind == ast.Object || each.Kind == ast.Interface {
			td := TypeData{
				Comment: formatComment(each.Description),
				Kind:    string(each.Kind),
				Name:    each.Name,
			}
			for _, other := range withInheritedFields(doc, each) {
				// is direct field or query
				if len(other.Arguments) > 0 {
					functionType := each.Name + fieldName(other.Name) + "Function"
					if !g.hasFunctionDefinition(functionType) {
						fnc := Function{
							Signature:  composeFunctionSignature(other),
							Type:       functionType,
							Arguments:  other.Arguments,
							IsArray:    isArray(other.Type),
							ReturnType: g.mapScalar(other.Type.Name()),
						}
						g.functions = append(g.functions, fnc)
					}
					td.Fields = append(td.Fields, FieldData{
						Comment: formatComment(other.Description),
						Name:    fieldName(other.Name),
						Type:    "*" + functionType, // result is optional so use pointer
						Tag:     fmt.Sprintf("`graphql:\"%s\" json:\"%s,omitempty\"`", other.Name, other.Name),
					})
				} else {
					td.Fields = append(td.Fields, g.buildFieldData(other))
				}
			}
			fd.Types = append(fd.Types, td)
		}
	}
	return tmpl.Execute(out, fd)
}

func withInheritedFields(schema *ast.SchemaDocument, typeDef *ast.Definition) ast.FieldList {
	all := append(ast.FieldList{}, typeDef.Fields...)
	for _, each := range typeDef.Interfaces {
		if idef := schema.Definitions.ForName(each); idef != nil {
			// prevent duplicate field
			for _, other := range idef.Fields {
				include := true
				for _, included := range all {
					if included.Name == other.Name {
						include = false
						break
					}
				}
				if include {
					all = append(all, other)
				}
			}
		}
	}
	return all
}

func composeFunctionSignature(other *ast.FieldDefinition) string {
	b := new(bytes.Buffer)
	io.WriteString(b, other.Name)
	if len(other.Arguments) > 0 {
		io.WriteString(b, "(")
		for i, each := range other.Arguments {
			if i > 0 {
				io.WriteString(b, ",")
			}
			io.WriteString(b, each.Name)
			io.WriteString(b, ":")
			io.WriteString(b, each.Type.String())
		}
		io.WriteString(b, ")")
	}
	io.WriteString(b, ":")
	io.WriteString(b, other.Type.String())
	return b.String()
}

func (g *Generator) buildFieldData(other *ast.FieldDefinition) FieldData {
	return FieldData{
		Comment:  formatComment(other.Description),
		Optional: !other.Type.NonNull,
		Name:     fieldName(other.Name),
		Type:     g.mapScalar(other.Type.Name()),
		IsArray:  isArray(other.Type),
		Tag:      fmt.Sprintf("`graphql:\"%s\" json:\"%s,omitempty\"`", other.Name, other.Name),
	}
}
