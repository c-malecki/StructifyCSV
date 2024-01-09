package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"

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

func (a *App) getCsvHeaders(absPath string) []string {
	file, err := os.Open(absPath)
	checkForError(err)
	defer file.Close()

	reader := csv.NewReader(file)
	headers, err := reader.Read()
	checkForError(err)
	return headers
}

type SelectFileRes struct {
	Headers []string `json:"headers"`
	Err     bool     `json:"error"`
}

type FieldType struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Field struct {
	FieldName string    `json:"fieldName"`
	FieldType FieldType `json:"fieldType"`
	Parent    string    `json:"parent"`
}

type Schema struct {
	Name   string  `json:"name"`
	Fields []Field `json:"fields"`
}

func (a *App) SelectFile() SelectFileRes {
	file, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select a CSV file",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "CSV (*.csv)",
				Pattern:     "*.csv",
			},
		},
	})

	return SelectFileRes{a.getCsvHeaders(file), err != nil}
}
