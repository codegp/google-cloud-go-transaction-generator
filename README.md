# google-cloud-go-transaction-generator
Generates common transaction functions for google cloude datastore models

### Usage

##### Install
```
go install github.com/codegp/google-cloud-go-transaction-generator
```

##### Make an alias (Optional, Recommended)
This package name is ridiculously long. I recommend you add an alias to your bash profile so you don't have to type google-cloud-go-transaction-generator everytime you want to run this thing. I use gcgtg, which is used in the examples below.

##### Define a config.yaml

```
receiver: MyClient
generators:
  - modelPackageName: exampleModels
    modelsImportPath: github.com/example/exampleModels
    models:
      - ModelA
      - ModelB
  - modelPackageName: demoModels
    modelsImportPath: github.com/demo/demoModels
    models:
      - ModelC
      - ModelD
    outputDir: "./customOutputDir"
outputDir: "./defaultOutputDir"
```

##### Generate the codez
```
gcgtg ./config.yaml
```
