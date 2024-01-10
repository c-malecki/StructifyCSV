package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
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

type SelectFileRes struct {
	Headers []string `json:"headers"`
	Err     bool     `json:"error"`
}

func (a *App) getCsvHeaders(absPath string) []string {
	file, err := os.Open(absPath)
	checkForError(err)
	defer file.Close()

	reader := csv.NewReader(file)
	headers, err := reader.Read()
	checkForError(err)
	return headers
}

func (a *App) ImportCsv() SelectFileRes {
	file, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Import CSV file",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "CSV (*.csv)",
				Pattern:     "*.csv",
			},
		},
	})

	return SelectFileRes{a.getCsvHeaders(file), err != nil}
}

func (a *App) WriteSchema(schema Schema) {
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

func (a *App) WriteModel(model Model) {
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

// save schema
// save model

// need to validate json
// import schema
// import model
