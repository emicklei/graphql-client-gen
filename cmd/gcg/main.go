package main

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/emicklei/gcg"
)

var oSchema = flag.String("schema", "", "SDL file")
var oPackage = flag.String("pkg", "generated", "package in which type are generated")

func main() {
	flag.Parse()
	data, err := ioutil.ReadFile(*oSchema)
	if err != nil {
		log.Fatalln(err)
	}
	gen := gcg.NewGenerator(string(data),
		gcg.WithPackage(*oPackage))
	err = gen.Generate()
	if err != nil {
		log.Fatalln(err)
	}
}
