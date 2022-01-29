package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"

	gcg "github.com/emicklei/graphql-client-gen"
)

const DEFAULT_CONFIG_FILE = "schema-generate.yaml"

var oSchema = flag.String("schema", "", "GraphQL schema SDL file (default 'schema.gql')")
var oPackage = flag.String("pkg", "", "package in which types are generated (default 'generated')")
var oConfig = flag.String("config", "", fmt.Sprintf("a YAML config file containing scalar binding mapping (default %s)", DEFAULT_CONFIG_FILE))

type Config struct {
	SchemaFile     string            `yaml:"schema"`
	PackageName    string            `yaml:"package"`
	ScalarBindings map[string]string `yaml:"bindings"`
}

func readConfig() *Config {
	flag.Parse()

	var configFile string

	if *oConfig != "" {
		if _, err := os.Stat(*oConfig); errors.Is(err, os.ErrNotExist) {
			log.Fatalf("The config file `%s` does not seem to exist\n", *oConfig)
		}
		configFile = *oConfig
	} else {
		if _, err := os.Stat(DEFAULT_CONFIG_FILE); err == nil {
			configFile = DEFAULT_CONFIG_FILE
		}
	}

	config := Config{"schema.gql", "generated", map[string]string{}}

	if configFile != "" {
		// Parse the configuration
		data, err := ioutil.ReadFile(configFile)
		if err != nil {
			log.Fatalln(err)
		}

		err = yaml.Unmarshal(data, &config)

		if err != nil {
			log.Fatalln(err)
		}
	}

	if *oSchema != "" {
		config.SchemaFile = *oSchema
	}

	if *oPackage != "" {
		config.PackageName = *oPackage
	}

	return &config
}

func main() {
	config := readConfig()

	data, err := ioutil.ReadFile(config.SchemaFile)
	if err != nil {
		log.Fatalf("Cannot find schema file `%s`", config.SchemaFile)
	}

	gen := gcg.NewGenerator(string(data),
		gcg.WithScalarBindings(config.ScalarBindings),
		gcg.WithPackage(config.PackageName))

	err = gen.Generate()
	if err != nil {
		log.Fatalln(err)
	}
}
