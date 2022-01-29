package gcg

import (
	"os"
	"text/template"
)

func (g *Generator) handleBuildTools() error {
	out, err := os.Create("build.go")
	if err != nil {
		return err
	}
	defer out.Close()
	fd := FileData{
		Package:      g.packageName,
		BuildVersion: g.mainVersion,
	}
	tmpl, err := template.New("tt").Parse(buildSrcTemplate)
	if err != nil {
		return err
	}
	fd.Build = BuildData{
		QueryTag:         "`json:\"query\"`",
		OperationNameTag: "`json:\"operationName\"`",
		VariablesTag:     "`json:\"variables\"`",
	}
	return tmpl.Execute(out, fd)
}
