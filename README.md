# google-cloud-go-transaction-generator
Generates common transaction functions for google cloude datastore models

### Usage

##### Install
```
go install -o gcgtg github.com/codegp/google-cloud-go-transaction-generator
```

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
