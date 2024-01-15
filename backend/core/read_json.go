package core

import (
	"context"
	"csvtoschema/backend/entity"
	"encoding/csv"
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func buildOpenDialogOptions(fileType string, entityType string) runtime.OpenDialogOptions {
	return runtime.OpenDialogOptions{
		Title: "Import " + entityType + "." + fileType,
		Filters: []runtime.FileFilter{
			{
				DisplayName: fileType + "(*." + fileType,
				Pattern:     "*." + fileType,
			},
		},
	}
}

func prepareOpenFileDialog(c context.Context, fileType string, entityType string) string {
	opts := buildOpenDialogOptions(fileType, entityType)
	filePath, err := runtime.OpenFileDialog(c, opts)
	if err != nil {
		print(err)
	}
	return filePath
}

func ImportSchemaJson(c context.Context) entity.Schema {
	filePath := prepareOpenFileDialog(c, "json", "Schema")
	var schema entity.Schema
	b, err := os.ReadFile(filePath)
	if err != nil {
		print(err)
	}
	json.Unmarshal(b, &schema)
	return schema
}

func ImportModelJson(c context.Context) entity.Model {
	filePath := prepareOpenFileDialog(c, "json", "Model")
	var model entity.Model
	b, err := os.ReadFile(filePath)
	if err != nil {
		print(err)
	}
	json.Unmarshal(b, &model)
	return model
}

func ImportCsvFileData(c context.Context) entity.CsvData {
	filePath := prepareOpenFileDialog(c, "csv", "")
	file, err := os.Open(filePath)
	if err != nil {
		print(err)
	}
	defer file.Close()

	name := filepath.Base(filePath)

	reader := csv.NewReader(file)
	headers, err := reader.Read()
	if err != nil {
		print(err)
	}
	return entity.CsvData{FileName: name, Location: filePath, Headers: headers}
}
