package gcg

import (
	"log"
	"os"
	"text/template"

	"github.com/vektah/gqlparser/ast"
)

func (g *Generator) handleInputs(doc *ast.SchemaDocument, all []*ast.Definition) error {
	log.Println("generating inputs.go")
	out, err := os.Create("inputs.go")
	if err != nil {
		return err
	}
	defer out.Close()
	fd := FileData{
		Package:      g.packageName,
		BuildVersion: g.mainVersion,
	}
	tmpl, err := template.New("tt").Parse(inputTemplateSrc)
	if err != nil {
		return err
	}
	for _, each := range all {
		id := InputData{
			Comment: formatComment(each.Description),
			Name:    each.Name,
		}
		for _, other := range each.Fields {
			id.Fields = append(id.Fields, FieldData{
				StructName: each.Name,
				Comment:    formatComment(other.Description),
				Name:       fieldName(other.Name),
				JSONName:   other.Name,
				IsArray:    isArray(other.Type),
				Optional:   !other.Type.NonNull,
				Type:       g.mapScalar(other.Type.Name()),
				GraphType:  other.Type.String(),
				Deprecated: isDeprecatedField(other),
			})
		}
		fd.Inputs = append(fd.Inputs, id)
	}
	return tmpl.Execute(out, fd)
}

func isDeprecatedField(def *ast.FieldDefinition) bool {
	return def.Directives.ForName("deprecated") != nil
}
