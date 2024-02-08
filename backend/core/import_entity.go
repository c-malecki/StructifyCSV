package core

import (
	"context"
	"csvtoschema/backend/entity"
	"csvtoschema/backend/ui"
	"encoding/csv"
	"encoding/json"
	"os"
	"path/filepath"
)

func ImportJsonSchema(c context.Context) entity.JsonSchema {
	filePath := ui.PrepareOpenFileDialog(c, "json", "schema")
	var schema entity.JsonSchema
	b, err := os.ReadFile(filePath)
	if err != nil {
		print(err)
	}
	json.Unmarshal(b, &schema)
	return schema
}

func ImportCsvData(c context.Context) entity.CsvData {
	filePath := ui.PrepareOpenFileDialog(c, "csv", "")
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
