package gcg

import (
	"github.com/iancoleman/strcase"
	"github.com/vektah/gqlparser/ast"
	"github.com/vektah/gqlparser/parser"
)

type Generator struct {
	schemaSource string
	packageName  string
}

func NewGenerator(schemaSource string, options ...Option) *Generator {
	g := &Generator{schemaSource: schemaSource, packageName: "generated"}
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
	if each := doc.Definitions.ForName("Mutation"); each != nil {
		if err := g.handleMutations(each); err != nil {
			return err
		}
	}
	enums := []*ast.Definition{}
	for _, each := range doc.Definitions {
		if each.Kind == ast.Enum {
			enums = append(enums, each)
		}
	}
	scalars := []*ast.Definition{}
	for _, each := range doc.Definitions {
		if each.Kind == ast.Scalar {
			// filter standards
			if mapScalar(each.Name) == each.Name {
				scalars = append(scalars, each)
			}
		}
	}
	if err := g.handleScalars(scalars); err != nil {
		return err
	}
	if err := g.handleEnums(enums); err != nil {
		return err
	}
	if err := g.handleTypes(doc); err != nil {
		return err
	}
	return nil
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

// https://github.com/shurcooL/graphql/blob/master/scalar.go
func mapScalar(name string) string {
	switch name {
	case "Boolean":
		return "bool"
	case "Float":
		return "float64"
	case "ID":
		return "interface{}"
	case "Int":
		return "int32"
	case "String":
		return "string"
	}
	return name
}
