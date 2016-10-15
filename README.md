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
outputPackage: outputpackage # package in which the generated code will live
receiver: MyClient # struct that has a DatastoreClient getter
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
    outputDir: "./customOutputDir" # custom output directory for this set of models
    outputPackage: "customoutputpackage" # custom output package for this set of models
outputDir: "./defaultOutputDir"
```

##### Generate the codez
```
gcgtg ./config.yaml
```


##### Example
 ```
 gcgtg ./example/exampleConfig.yaml
 ```

 Generates:

 ```
 package examplepackage

 import (
 	"fmt"

 	"cloud.google.com/go/datastore"
 	"github.com/codegp/google-cloud-go-transaction-generator/examplemodel"

 	"golang.org/x/net/context"
 )

 // GetExampleModel retrieves a ExampleModel by its ID.
 func (c *ExampleReceiver) GetExampleModel(id int64) (*examplemodel.ExampleModel, error) {
 	ctx := context.Background()
 	k := datastore.NewKey(ctx, "ExampleModel", "", id, nil)
 	ExampleModel := &examplemodel.ExampleModel{}
 	if err := c.DatastoreClient().Get(ctx, k, ExampleModel); err != nil {
 		return nil, fmt.Errorf("datastoredb: could not get ExampleModel: %v", err)
 	}
 	ExampleModel.ID = id
 	return ExampleModel, nil
 }

 // AddExampleModel saves a given ExampleModel, assigning it a new ID.
 func (c *ExampleReceiver) AddExampleModel(b *examplemodel.ExampleModel) (*examplemodel.ExampleModel, error) {
 	ctx := context.Background()
 	k := datastore.NewIncompleteKey(ctx, "ExampleModel", nil)
 	k, err := c.DatastoreClient().Put(ctx, k, b)
 	if err != nil {
 		return nil, fmt.Errorf("datastoredb: could not put ExampleModel: %v", err)
 	}
 	b.ID = k.ID()
 	return b, nil
 }

 // DeleteExampleModel removes a given ExampleModel by its ID.
 func (c *ExampleReceiver) DeleteExampleModel(id int64) error {
 	ctx := context.Background()
 	k := datastore.NewKey(ctx, "ExampleModel", "", id, nil)
 	if err := c.DatastoreClient().Delete(ctx, k); err != nil {
 		return fmt.Errorf("datastoredb: could not delete ExampleModel: %v", err)
 	}
 	return nil
 }

 // UpdateExampleModel updates the entry for a given ExampleModel.
 func (c *ExampleReceiver) UpdateExampleModel(b *examplemodel.ExampleModel) error {
 	ctx := context.Background()
 	k := datastore.NewKey(ctx, "ExampleModel", "", b.ID, nil)
 	if _, err := c.DatastoreClient().Put(ctx, k, b); err != nil {
 		return fmt.Errorf("datastoredb: could not update ExampleModel: %v", err)
 	}
 	return nil
 }

 // ListExampleModels returns a list of ExampleModels
 func (c *ExampleReceiver) ListExampleModels() ([]*examplemodel.ExampleModel, error) {
 	ctx := context.Background()
 	ExampleModels := make([]*examplemodel.ExampleModel, 0)
 	q := datastore.NewQuery("ExampleModel")

 	keys, err := c.DatastoreClient().GetAll(ctx, q, &ExampleModels)

 	if err != nil {
 		return nil, fmt.Errorf("datastoredb: could not list ExampleModels: %v", err)
 	}

 	for i, k := range keys {
 		ExampleModels[i].ID = k.ID()
 	}

 	return ExampleModels, nil
 }

 //  QueryExampleModelsByProp
 func (c *ExampleReceiver) QueryExampleModelsByProp(propName, value string) (*examplemodel.ExampleModel, error) {
 	ctx := context.Background()
 	ExampleModels := make([]*examplemodel.ExampleModel, 0)
 	q := datastore.NewQuery("ExampleModel").Filter(fmt.Sprintf("%s =", propName), value)

 	keys, err := c.DatastoreClient().GetAll(ctx, q, &ExampleModels)

 	if err != nil {
 		return nil, fmt.Errorf("datastoredb: could not list ExampleModels: %v", err)
 	}

 	if len(ExampleModels) == 0 {
 		return nil, nil
 	}

 	ExampleModels[0].ID = keys[0].ID()
 	return ExampleModels[0], nil
 }

 ```
