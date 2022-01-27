package gcg

import (
	"bytes"
	"fmt"
	"os"
	"text/template"
	"time"

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
		Package: g.packageName,
		Created: time.Now(),
	}
	tmpl, err := template.New("tt").Parse(typeTemplateSrc)
	if err != nil {
		return err
	}
	for _, each := range all {
		td := TypeData{
			Comment: formatComment(each.Description),
			Kind:    string(each.Kind),
			Name:    unionTypeName(each),
		}
		//fields := []*ast.FieldDefinition{}
		for _, other := range each.Types {
			typeDef := doc.Definitions.ForName(other)
			// assume never nil
			for _, his := range typeDef.Fields {
				fd := g.buildFieldData(his)
				td.Fields = append(td.Fields, fd)
			}
		}
		fd.Types = append(fd.Types, td)
	}
	return tmpl.Execute(out, fd)
}

func unionTypeName(union *ast.Definition) string {
	b := new(bytes.Buffer)
	for _, each := range union.Types {
		fmt.Fprint(b, each)
	}
	fmt.Fprint(b, "Union")
	return b.String()
}
