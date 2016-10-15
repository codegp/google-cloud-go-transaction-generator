package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"

	yaml "gopkg.in/yaml.v2"
)

/*
  This file generates a file with crud operation functions for each model type
*/
const templatePath = "../templates/modeltemplate.go.tmpl"

type config struct {
	receiver   string
	generators []*generator
	outputDir  string
}

type generator struct {
	modelPackageName string
	modelsImportPath string
	models           []string
	outputDir        string
}

type templateData struct {
	Model              string
	ModelPackage       string
	ModelPackageImport string
	Receiver           string
}

var cfg *config
var cfgPath string
var tmpl *template.Template

func init() {
	flag.Parse()
	if len(flag.Args()) < 1 {
		fmt.Printf("Usage: google-cloud-go-transaction-generator path/to/config.yaml")
		os.Exit(1)
	}

	cfgPath = flag.Arg(1)
	cfgBytes, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		fmt.Printf("Could not locate config file at path %s", cfgPath)
		os.Exit(1)
	}

	err = yaml.Unmarshal(cfgBytes, cfg)
	if err != nil {
		fmt.Printf("Could not parse configuration yaml file, %v", err)
		os.Exit(1)
	}

	tmpl, err = template.ParseFiles(templatePath)
	if err != nil {
		fmt.Printf("Failed to parse template, %v", err)
		os.Exit(1)
	}
}

func main() {
	for i := range cfg.generators {
		gen(i)
	}
}

func gen(generatorIndex int) {
	generator := cfg.generators[generatorIndex]
	fmt.Printf("Generating models for %s", generator.modelPackageName)

	for _, model := range generator.models {
		fmt.Printf("   Generating model %s", model)

		outputDir := cfg.outputDir
		if generator.outputDir != "" {
			outputDir = generator.outputDir
		}

		outputFile := fmt.Sprintf("%s/gen_%s.go", outputDir, model)
		f, err := os.Create(outputFile)
		if err != nil {
			fmt.Printf("Failed to create file at %s", outputFile)
		}

		writer := bufio.NewWriter(f)

		err = tmpl.Execute(writer, &templateData{
			Model:              model,
			ModelPackage:       generator.modelPackageName,
			ModelPackageImport: generator.modelsImportPath,
			Receiver:           cfg.receiver,
		})

		if err != nil {
			fmt.Printf("Failed to execute template: %v", err)
		}

		writer.Flush()
		f.Close()
	}
}
