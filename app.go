package main

import (
	"context"
	"csvtoschema/backend/core"
	"csvtoschema/backend/entity"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

func (a *App) parseCsv(absPath string) entity.CsvData {
	file, err := os.Open(absPath)
	checkForError(err)
	defer file.Close()

	name := filepath.Base(absPath)

	reader := csv.NewReader(file)
	headers, err := reader.Read()
	checkForError(err)
	return entity.CsvData{FileName: name, Location: absPath, Headers: headers}
}

func (a *App) ImportCsvFile() entity.CsvData {
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
	return entity.CsvData{}
}

func (a *App) ExportSchema(schema entity.Schema) {
	core.ExportSchemaToJson(a.ctx, schema)
}

func (a *App) ExportModel(model entity.Model) {
	core.ExportModelToJson(a.ctx, model)
}

func (a *App) ImportSchema() entity.Schema {
	return core.ImportSchemaJson(a.ctx)
}

func (a *App) ImportModel() entity.Model {
	return core.ImportModelJson(a.ctx)
}

func prepareOpenDirectoryDialog(c context.Context) string {
	dirPath, err := runtime.OpenDirectoryDialog(c, runtime.OpenDialogOptions{
		Title: "Select Directory for Output",
	})
	if err != nil {
		print(err)
	}
	return dirPath
}

func (a *App) ProcessCsvDescriptor(model entity.Model, headerDescriptors []entity.HeaderDescriptor, csvLocation string) {
	outDir := prepareOpenDirectoryDialog(a.ctx)

	writerCh := make(chan map[string]string)
	doneCh := make(chan bool)

	go processCsvFile(csvLocation, headerDescriptors, writerCh)
	go writeJsonFile(outDir, csvLocation, writerCh, doneCh)
	<-doneCh
}

func processCsvFile(csvLocation string, headerDescriptors []entity.HeaderDescriptor, writerCh chan<- map[string]string) {
	file, err := os.Open(csvLocation)
	if err != nil {
		print(err)
	}
	defer file.Close()

	var headers, line []string
	reader := csv.NewReader(file)
	headers, err = reader.Read()
	if err != nil {
		print(err)
	}
	var headerIndexes []int
	for i, h := range headers {
		for _, hd := range headerDescriptors {
			if h == hd.Header {
				headerIndexes = append(headerIndexes, i)
			}
		}
	}

	for {
		line, err = reader.Read()

		if err == io.EOF {
			close(writerCh)
			break
		} else if err != nil {
			print(err)
		}

		record, err := processLine(headers, line, headerIndexes)

		if err != nil {
			// need to return array of lines which were not processed to frontend
			fmt.Printf("Line: %sError: %s\n", line, err)
			continue
		}
		writerCh <- record
	}
}

func writeJsonFile(outDir string, csvLocation string, writerCh <-chan map[string]string, done chan<- bool) {
	jsonFunc := func(record map[string]string) string {
		jsonData, _ := json.MarshalIndent(record, "  ", "  ")
		return "  " + string(jsonData)
	}
	writeString := createStringWriter(outDir, csvLocation)
	writeString("[\n", false)

	first := true
	for {
		record, more := <-writerCh
		if more {
			if !first {
				writeString(",\n", false)
			} else {
				first = false
			}

			jsonData := jsonFunc(record)
			writeString(jsonData, false)
		} else {
			writeString("\n]", true)
			done <- true
			break
		}
	}
}

func createStringWriter(outDir string, csvLocation string) func(string, bool) {
	jsonName := fmt.Sprintf("%s_Output.json", strings.TrimSuffix(filepath.Base(csvLocation), ".csv"))
	finalLocation := fmt.Sprintf("%s/%s", outDir, jsonName)

	file, err := os.Create(finalLocation)
	if err != nil {
		print(err)
	}

	return func(data string, close bool) {
		_, err := file.WriteString(data)
		if err != nil {
			print(err)
		}

		if close {
			file.Close()
		}
	}
}
func processLine(headers []string, lineData []string, headerIndexes []int) (map[string]string, error) {
	if len(lineData) != len(headers) {
		return nil, errors.New("line doesn't match headers format. skipping")
	}

	recordMap := make(map[string]string)

	for _, index := range headerIndexes {
		recordMap[headers[index]] = lineData[index]
	}

	return recordMap, nil
}

// 	for each headerDescriptor, find idx in headers
// 	when parsing lines, skip all index that aren't included in header indexes
// 	at value of index in line, check against descriptor and convert strings to data type
// 	write line
