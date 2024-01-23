package gcg

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"

	_ "embed"
)

//go:embed build_source.go
var buildSource string

func (g *Generator) handleBuildTools() error {
	log.Println("generating build.go")
	out, err := os.Create(path.Join(g.targetDirectory, "build.go"))
	if err != nil {
		return err
	}
	defer out.Close()
	replaced := strings.Replace(buildSource, "package gcg", fmt.Sprintf(`package %s
	// copied from build_source.go by github.com/emicklei/graphql-client-gen/cmd/gcg version: %s
	// DO NOT EDIT`, g.packageName, g.mainVersion), -1)
	_, err = io.WriteString(out, replaced)
	return err
}
