package gcg

import (
	"bytes"
	"fmt"
	"runtime/debug"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/vektah/gqlparser/ast"
	"github.com/vektah/gqlparser/parser"
)

type ScalarBinding struct {
	GraphQLScalar string
	GoTypeName    string
}

type Generator struct {
	sourceFilename string
	schemaSource   string
	packageName    string
	functions      []Function
	scalarBindings []ScalarBinding
	mainVersion    string
	schemaVersion  string
}

func NewGenerator(schemaSource string, options ...Option) *Generator {
	g := &Generator{schemaSource: schemaSource, packageName: "generated"}
	// need version to put in generated files
	bi, ok := debug.ReadBuildInfo()
	if ok && len(bi.Main.Version) > 0 {
		g.mainVersion = bi.Main.Version
	} else {
		g.mainVersion = "(dev)"
	}
	for _, each := range options {
		each(g)
	}

	// add default scalar mappings
	// https://github.com/shurcooL/graphql/blob/master/scalar.go
	g.scalarBindings = append(g.scalarBindings, ScalarBinding{"Boolean", "bool"})
	g.scalarBindings = append(g.scalarBindings, ScalarBinding{"Float", "float64"})
	g.scalarBindings = append(g.scalarBindings, ScalarBinding{"ID", "interface{}"})
	g.scalarBindings = append(g.scalarBindings, ScalarBinding{"Int", "int32"})
	g.scalarBindings = append(g.scalarBindings, ScalarBinding{"String", "string"})

	return g
}

func (g *Generator) Generate() error {
	doc, perr := parser.ParseSchema(&ast.Source{Input: g.schemaSource, Name: g.sourceFilename})
	if perr != nil {
		return perr
	}
	// get version if available
	for _, s := range doc.Schema {
		for _, each := range s.Directives {
			if each.Name == "version" {
				for _, other := range each.Arguments {
					if other.Name == "name" {
						g.schemaVersion = other.Value.Raw
						break
					}
				}
			}
		}
	}

	if each := doc.Definitions.ForName("Mutation"); each != nil {
		if err := g.handleMutations(each); err != nil {
			return err
		}
	}
	if each := doc.Definitions.ForName("Query"); each != nil {
		if err := g.handleQueries(each); err != nil {
			return err
		}
	}
	enums := []*ast.Definition{}
	for _, each := range doc.Definitions {
		if each.Kind == ast.Enum {
			enums = append(enums, each)
		}
	}
	// Find scalars that need code generation
	scalarDefs := []*ast.Definition{}
	for _, each := range doc.Definitions {
		if each.Kind == ast.Scalar {
			// filter standards
			if g.scalarMustBeGenerated(each.Name) {
				scalarDefs = append(scalarDefs, each)
			}
		}
	}
	unions := []*ast.Definition{}
	for _, each := range doc.Definitions {
		if each.Kind == ast.Union {
			unions = append(unions, each)
		}
	}
	if err := g.handleUnions(doc, unions); err != nil {
		return err
	}
	if err := g.handleScalars(scalarDefs); err != nil {
		return err
	}
	if err := g.handleEnums(enums); err != nil {
		return err
	}
	if err := g.handleTypes(doc); err != nil {
		return err
	}
	if err := g.handleFunctions(); err != nil {
		return err
	}
	if err := g.handleBuildTools(); err != nil {
		return err
	}
	return nil
}

func (g *Generator) scalarMustBeGenerated(name string) bool {
	return !g.isScalarTypeProvided(name)
}

func (g *Generator) isScalarTypeProvided(name string) bool {
	for _, each := range g.scalarBindings {
		if each.GraphQLScalar == name {
			return true
		}
	}
	return false
}

func (g *Generator) mapScalar(name string) string {
	for _, each := range g.scalarBindings {
		if each.GraphQLScalar == name {
			return each.GoTypeName
		}
	}

	return name
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

func formatComment(comment string) string {
	lines := strings.Split(comment, "\n")
	if len(lines) <= 1 {
		return comment
	}
	b := new(bytes.Buffer)
	for _, each := range lines {
		fmt.Fprintf(b, "\n// %s", each)
	}
	return b.String()
}

func goArgName(name string) string {
	// check for reserved Go names to prevent syntax errors
	return "_" + name
}

func (g *Generator) hasFunctionDefinition(typeName string) bool {
	for _, each := range g.functions {
		if each.Type == typeName {
			return true
		}
	}
	return false
}
