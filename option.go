package gcg

type Option func(*Generator)

func WithPackage(name string) func(*Generator) {
	return func(g *Generator) {
		g.packageName = name
	}
}
