package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"

	"gopkg.in/yaml.v2"
)

/*
  This file generates a file with crud operation functions for each model type
*/
const templatePath = "templates/modeltemplate.go.tmpl"

type config struct {
	OutputPackage string       `yaml:"outputPackage"`
	Receiver      string       `yaml:"receiver"`
	Generators    []*generator `yaml:"generators"`
	OutputDir     string       `yaml:"outputDir"`
}

type generator struct {
	ModelPackageName string   `yaml:"modelPackageName"`
	ModelImportPath  string   `yaml:"modelsImportPath"`
	Models           []string `yaml:"models"`
	OutputPackage    string   `yaml:"outputPackage"`
	OutputDir        string   `yaml:"outputDir"`
}

type templateData struct {
	OutputPackage      string
	Model              string
	ModelPackage       string
	ModelPackageImport string
	Receiver           string
}

var cfg config
var cfgPath string
var tmpl *template.Template

func init() {
	flag.Parse()
	if len(flag.Args()) < 1 {
		fmt.Printf("Usage: google-cloud-go-transaction-generator path/to/config.yaml")
		os.Exit(1)
	}

	cfgPath = flag.Arg(0)
	cfgBytes, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		fmt.Printf("Could not read config file at path %s", cfgPath)
		os.Exit(1)
	}

	cfg = config{}
	err = yaml.Unmarshal(cfgBytes, &cfg)
	if err != nil {
		fmt.Printf("Could not parse configuration yaml file, %v", err)
		os.Exit(1)
	}

	tmpl = template.New("tmpl")
	tmplContent, err := Asset(templatePath)
	if err != nil {
		fmt.Printf("Failed to parse binded template, %v", err)
	}

	tmpl, err = tmpl.Parse(string(tmplContent))
	if err != nil {
		fmt.Printf("Failed to parse template, %v", err)
		os.Exit(1)
	}
}

func main() {
	for i := range cfg.Generators {
		gen(i)
	}
}

func gen(generatorIndex int) {
	generator := cfg.Generators[generatorIndex]
	fmt.Printf("Generating models for %s\n", generator.ModelPackageName)

	for _, model := range generator.Models {
		fmt.Printf("   Generating model %s\n", model)

		outputDir := cfg.OutputDir
		if generator.OutputDir != "" {
			outputDir = generator.OutputDir
		}

		outputPackage := cfg.OutputPackage
		if generator.OutputPackage != "" {
			outputPackage = generator.OutputPackage
		}

		outputFile := fmt.Sprintf("%s/gen_%s.go", outputDir, model)
		f, err := os.Create(outputFile)
		if err != nil {
			fmt.Printf("Failed to create file at %s", outputFile)
			os.Exit(1)
		}

		writer := bufio.NewWriter(f)

		err = tmpl.Execute(writer, &templateData{
			OutputPackage:      outputPackage,
			Model:              model,
			ModelPackage:       generator.ModelPackageName,
			ModelPackageImport: generator.ModelImportPath,
			Receiver:           cfg.Receiver,
		})

		if err != nil {
			fmt.Printf("Failed to execute template: %v", err)
			os.Exit(1)
		}

		writer.Flush()
		f.Close()
	}
}
