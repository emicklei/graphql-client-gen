package gcg

import (
	"os"
	"text/template"
	"time"
)

func (g *Generator) handleBuildTools() error {
	out, err := os.Create("build.go")
	if err != nil {
		return err
	}
	defer out.Close()
	fd := FileData{
		Package: g.packageName,
		Created: time.Now(),
	}
	tmpl, err := template.New("tt").Parse(buildSrcTemplate)
	if err != nil {
		return err
	}
	return tmpl.Execute(out, fd)
}
