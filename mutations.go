package gcg

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/iancoleman/strcase"
	"github.com/vektah/gqlparser/ast"
)

func (g *Generator) handleMutations(each *ast.Definition) error {
	out, err := os.Create("mutations.go")
	if err != nil {
		return err
	}
	defer out.Close()
	fd := FileData{
		Package: g.packageName,
		Created: time.Now(),
	}
	tmpl, err := template.New("tt").Parse(operationsTemplateSrc)
	if err != nil {
		return err
	}
	for _, other := range each.Fields {
		rt := mapScalar(other.Type.Name())
		rttokens := strings.Split(rt, ".")
		od := OperationData{
			Comment:      other.Description,
			Name:         other.Name,
			FunctionName: strcase.ToCamel(other.Name),
			IsArray:      isArray(other.Type),
			ReturnType:   rt,
			ReturnField:  rttokens[len(rttokens)-1],
		}
		// `graphql:"addEmploymentDocument(employmentID: $employmentID, fileID: $fileID, repositoryID: $repositoryID)"`
		tag := new(bytes.Buffer)
		fmt.Fprintf(tag, "`graphql:\"%s(", other.Name)
		for i, arg := range other.Arguments {
			if i > 0 {
				fmt.Fprintf(tag, ",")
			}
			fmt.Fprintf(tag, "%s: $%s", arg.Name, arg.Name)
			od.Arguments = append(od.Arguments, Argument{
				Name: arg.Name, Type: mapScalar(arg.Type.Name()), IsArray: isArray(arg.Type)})
		}
		fmt.Fprintf(tag, ")\"`")
		od.Tag = tag.String()
		fd.Operations = append(fd.Operations, od)
	}
	return tmpl.Execute(out, fd)
}
