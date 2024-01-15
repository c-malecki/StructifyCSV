package core

import (
	"context"
	"csvtoschema/backend/entity"
	"encoding/json"
	"os"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func buildSaveDialogOpts(entityName string, entityType string) runtime.SaveDialogOptions {
	fileName := strings.ReplaceAll(entityName, " ", "_")
	return runtime.SaveDialogOptions{
		DefaultDirectory: ".",
		DefaultFilename:  fileName + "_" + entityType + ".json",
		Title:            "Save " + entityName + " " + entityType,
	}
}

func prepareSaveFileDialog(c context.Context, entityName string, entityType string) string {
	opts := buildSaveDialogOpts(entityName, entityType)
	filePath, err := runtime.SaveFileDialog(c, opts)
	if err != nil {
		print(err)
	}
	return filePath
}

func ExportSchemaToJson(c context.Context, schema entity.Schema) {
	filePath := prepareSaveFileDialog(c, schema.Name, "Schema")
	if len(filePath) > 0 {
		b, _ := json.MarshalIndent(schema, "", "  ")
		err := os.WriteFile(filePath, b, 0666)
		if err != nil {
			print(err)
		}
	}
}

func ExportModelToJson(c context.Context, model entity.Model) {
	filePath := prepareSaveFileDialog(c, model.Name, "Model")
	if len(filePath) > 0 {
		b, _ := json.MarshalIndent(model, "", "  ")
		err := os.WriteFile(filePath, b, 0666)
		if err != nil {
			print(err)
		}
	}
}

func ExportCsvDescriptor(c context.Context, model entity.Model, hd []entity.HeaderDescriptor, csvPath string) {
	filePath := prepareSaveFileDialog(c, model.Name, "Batch")
	writerCh := make(chan map[string]string)
	doneCh := make(chan bool)

	go ProcessCsvFile(csvPath, hd, writerCh)
	go WriteDescriptorJson(filePath, csvPath, writerCh, doneCh)
	<-doneCh
}
