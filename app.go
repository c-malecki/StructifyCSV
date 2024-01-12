package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func exitOnError(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}

func checkForError(err error) {
	if err != nil {
		exitOnError(err)
	}
}

type DataType struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Field struct {
	Name     string   `json:"name"`
	Type     string   `json:"type"`
	DataType DataType `json:"dataType"`
}

type Schema struct {
	Name   string  `json:"name"`
	Type   string  `json:"type"`
	Fields []Field `json:"fields"`
}

type Model struct {
	Name       string   `json:"name"`
	Type       string   `json:"type"`
	Schemas    []Schema `json:"schemas"`
	BaseSchema int      `json:"baseSchema"`
}

type CsvData struct {
	FileName string   `json:"fileName"`
	Headers  []string `json:"headers"`
}

func (a *App) parseCsv(absPath string) CsvData {
	file, err := os.Open(absPath)
	checkForError(err)
	defer file.Close()

	name := filepath.Base(absPath)

	reader := csv.NewReader(file)
	headers, err := reader.Read()
	checkForError(err)
	return CsvData{name, headers}
}

func (a *App) ImportCsvFile() CsvData {
	file, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Import CSV file",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "CSV (*.csv)",
				Pattern:     "*.csv",
			},
		},
	})
	if err != nil {
		print(err)
	}
	if len(file) > 0 {
		return a.parseCsv(file)
	}
	return CsvData{}
}

func (a *App) ExportSchemaToJson(schema Schema) {
	fileName := strings.ReplaceAll(schema.Name, " ", "_")
	file, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		DefaultDirectory: ".",
		DefaultFilename:  fileName + "_Schema.json",
		Title:            "Save" + " " + schema.Name + " " + "Schema",
	})
	if err != nil {
		print(err)
	}
	if len(file) > 0 {
		b, _ := json.MarshalIndent(schema, "", "  ")
		err := os.WriteFile(file, b, 0666)
		if err != nil {
			print(err)
		}
	}
}

func (a *App) ImportSchemaJson() Schema {
	filePath, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Import Schema.json",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "JSON (*.json)",
				Pattern:     "*.json",
			},
		},
	})
	if err != nil {
		print(err)
	}

	var schema Schema
	b, err := os.ReadFile(filePath)
	if err != nil {
		print(err)
	}
	json.Unmarshal(b, &schema)

	return schema
}

func (a *App) ExportModelToJson(model Model) {
	fileName := strings.ReplaceAll(model.Name, " ", "_")
	file, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		DefaultDirectory: ".",
		DefaultFilename:  fileName + "_Model.json",
		Title:            "Save" + " " + model.Name + " " + "Model",
	})
	if err != nil {
		print(err)
	}
	if len(file) > 0 {
		b, _ := json.MarshalIndent(model, "", "  ")
		err := os.WriteFile(file, b, 0666)
		if err != nil {
			print(err)
		}
	}
}

func (a *App) ImportModelJson() Model {
	filePath, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Import Model.json",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "JSON (*.json)",
				Pattern:     "*.json",
			},
		},
	})
	if err != nil {
		print(err)
	}

	var model Model
	b, err := os.ReadFile(filePath)
	if err != nil {
		print(err)
	}
	json.Unmarshal(b, &model)

	return model
}

// save schema
// save model

// need to validate json
// import schema
// import model
