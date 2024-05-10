package gcg

type Option func(*Generator)

func WithScalarBindings(bindings map[string]string) func(*Generator) {
	return func(g *Generator) {
		for k, v := range bindings {
			g.scalarBindings = append(g.scalarBindings, ScalarBinding{k, v})
		}

	}
}

func WithPackage(name string) func(*Generator) {
	return func(g *Generator) {
		g.packageName = name
	}
}

func WithSource(filename string) func(*Generator) {
	return func(g *Generator) {
		g.sourceFilename = filename
	}
}

func WithOutputFolder(path string) func(*Generator) {
	return func(g *Generator) {
		g.outputFolder = path
	}
}
