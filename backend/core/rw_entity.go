package core

import (
	"context"
	"csvtoschema/backend/entity"
	"csvtoschema/backend/ui"
	"encoding/json"
	"os"
)

func ExportSchemaToJson(c context.Context, schema entity.Schema) {
	file := ui.PrepareSaveFileDialog(c, schema.Name, "Schema")
	if len(file) > 0 {
		b, _ := json.MarshalIndent(schema, "", "  ")
		err := os.WriteFile(file, b, 0666)
		if err != nil {
			print(err)
		}
	}
}

func ExportModelToJson(c context.Context, model entity.Model) {
	file := ui.PrepareSaveFileDialog(c, model.Name, "Model")
	if len(file) > 0 {
		b, _ := json.MarshalIndent(model, "", "  ")
		err := os.WriteFile(file, b, 0666)
		if err != nil {
			print(err)
		}
	}
}

func ImportSchemaJson(c context.Context) entity.Schema {
	file := ui.PrepareOpenFileDialog(c, "Schema")
	var schema entity.Schema
	b, err := os.ReadFile(file)
	if err != nil {
		print(err)
	}
	json.Unmarshal(b, &schema)

	return schema
}

func ImportModelJson(c context.Context) entity.Model {
	file := ui.PrepareOpenFileDialog(c, "Model")
	var model entity.Model
	b, err := os.ReadFile(file)
	if err != nil {
		print(err)
	}
	json.Unmarshal(b, &model)

	return model
}
