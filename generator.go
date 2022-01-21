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
