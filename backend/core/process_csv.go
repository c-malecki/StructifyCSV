package core

import (
	"StructifyCSV/backend/entity"
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func ImportCsvFileData(c context.Context) entity.CsvFileData {
	opts := runtime.OpenDialogOptions{
		Title: "Import CSV",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "CSV (*.csv)",
				Pattern:     "*.csv",
			},
		},
	}

	filePath, err := runtime.OpenFileDialog(c, opts)
	if err != nil {
		fmt.Printf("File Path Error: %s\n", err)
	}

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Open File Error: %s\n", err)
		// canceled import
		return entity.CsvFileData{}
	}

	defer file.Close()

	name := filepath.Base(filePath)

	reader := csv.NewReader(file)

	csvHeaders := make([]entity.CsvHeader, 0)

	var headers, row []string
	headers, err = reader.Read()

	for i, h := range headers {
		col, _ := IndexToColumn(i + 1)
		ch := entity.CsvHeader{Column: col, Header: h}
		csvHeaders = append(csvHeaders, ch)
	}

	firstFiveRows := make([][]string, 0, 5)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	rowNum := 0

	for {
		if rowNum >= 6 {
			file.Close()
			break
		}

		row, err = reader.Read()

		if err == io.EOF {
			file.Close()
			break
		} else if err != nil {
			// return some sort of error for not being able to read csv line
			fmt.Printf("Row: %sError: %s\n", row, err)
		}

		if len(row) != len(headers) {
			continue
		}

		firstFiveRows = append(firstFiveRows, row)

		rowNum++
	}

	return entity.CsvFileData{FileName: name, Location: filePath, Headers: csvHeaders, ReferenceRows: firstFiveRows}
}

func ProcessCsv(c context.Context, csvFile entity.CsvFileData, jsonSchema entity.JsonSchema) entity.CsvProcessingReport {
	fileName := strings.TrimSuffix(csvFile.FileName, ".csv")

	opts := runtime.SaveDialogOptions{
		DefaultDirectory: ".",
		DefaultFilename:  fileName + ".output.json",
		Title:            "Save " + fileName + " output",
	}

	filePath, err := runtime.SaveFileDialog(c, opts)

	if err != nil {
		fmt.Printf("File Path Error: %s\n", err)
	}

	var wg sync.WaitGroup
	var report entity.CsvProcessingReport

	successCh := make(chan int)
	rowErrCh := make(chan entity.RowError)
	resultCh := make(chan map[string]any)

	wg.Add(4)
	go processCsvWithJsonSchema(csvFile.Location, jsonSchema, resultCh, rowErrCh, successCh, &wg)
	go writeJsonFromResults(filePath, resultCh, &wg)
	go aggregateRowSuccesses(successCh, &report, &wg)
	go aggregateRowErrors(rowErrCh, &report, &wg)
	wg.Wait()

	return report
}

func processCsvWithJsonSchema(csvLocation string, jsonSchema entity.JsonSchema, resultCh chan<- map[string]any, rowErrCh chan<- entity.RowError, successCh chan<- int, wg *sync.WaitGroup) {
	file, err := os.Open(csvLocation)
	if err != nil {
		// return error that file could not be opened
		fmt.Printf("Open File Error: %s\n", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	var headers, row []string
	headers, err = reader.Read()
	if err != nil {
		// return some sort of error for not being able to read csv line
		fmt.Printf("Error: %s\n", err)
	}

	rowNum := 0

	for {
		row, err = reader.Read()

		if err == io.EOF {
			close(resultCh)
			close(rowErrCh)
			close(successCh)
			wg.Done()
			break
		} else if err != nil {
			// return some sort of error for not being able to read csv line
			fmt.Printf("Row: %sError: %s\n", row, err)
		}

		rowNum++

		if len(row) != len(headers) {
			rowErr := CreateHeaderMismatchErr(rowNum, len(row), len(headers))
			rowErrCh <- rowErr
			continue
		}

		rowMap := make(map[string]interface{})
		for key, propSchema := range jsonSchema.Properties {
			if propSchema.Type != "object" && propSchema.CsvHeaderIndex == nil {
				continue
			}
			processCsvRowToMap(key, *propSchema, rowNum, row, rowMap, rowErrCh)
		}
		successCh <- rowNum
		resultCh <- rowMap
	}
}

func processCsvRowToMap(propKey string, propSchema entity.PropertySchema, rowNum int, rowData []string, propMap map[string]interface{}, rowErrCh chan<- entity.RowError) {

	switch propSchema.Type {
	case "string":
		i64 := propSchema.CsvHeaderIndex.(float64)
		colNum := int(i64)
		v := rowData[colNum]
		isValid := ValidateStringProperty(v, propKey, propSchema, rowNum, colNum, rowErrCh)
		if isValid {
			propMap[propKey] = v
		}
	case "number":
		i64 := propSchema.CsvHeaderIndex.(float64)
		colNum := int(i64)
		v, ok := ConvertToFloat(rowData[colNum], propKey, propSchema, rowNum, colNum, rowErrCh)
		if ok {
			isValid := ValidateNumberProperty(v, propKey, propSchema, rowNum, colNum, rowErrCh)
			if isValid {
				propMap[propKey] = v
			}
		}
	case "integer":
		i64 := propSchema.CsvHeaderIndex.(float64)
		colNum := int(i64)
		v, ok := ConvertToInt(rowData[colNum], propKey, propSchema, rowNum, colNum, rowErrCh)
		if ok {
			isValid := ValidateIntegerProperty(v, propKey, propSchema, rowNum, colNum, rowErrCh)
			if isValid {
				propMap[propKey] = v
			}
		}
	case "boolean":
		i64 := propSchema.CsvHeaderIndex.(float64)
		colNum := int(i64)
		v, ok := ConvertToBool(rowData[colNum], propKey, propSchema, rowNum, colNum, rowErrCh)
		if ok {
			propMap[propKey] = v
		}
	case "null":
		propMap[propKey] = nil
	case "array":
		colNums := ConvertIndexArrToInts(propSchema.CsvHeaderIndex)

		arrItems := make([]any, 0, len(colNums))
		for _, colNum := range colNums {
			v, ok := ConvertArrayItemType(rowData[colNum], propKey, propSchema, rowNum, colNum, rowErrCh)
			if ok {
				arrItems = append(arrItems, v)
			}
		}

		isValid := ValidateArrayProperty(arrItems, propKey, propSchema, rowNum, rowErrCh)
		if isValid {
			propMap[propKey] = arrItems
		}
	case "object":
		nextMap := make(map[string]interface{})
		for key, propSchema := range propSchema.Properties {
			processCsvRowToMap(key, *propSchema, rowNum, rowData, nextMap, rowErrCh)
			isValid := ValidateObjectProperty(nextMap, key, *propSchema, rowNum, rowErrCh)
			if isValid {
				propMap[propKey] = nextMap
			}
		}
		// if map len > 0
		// rowSchema.Properties = propSchema.Properties
		// resultMap := processCsvRowToMap(rowSchema, nextMap, rowErrCh)

		// isValid := ValidateObjectProperty(propSchema, "", rowSchema.RowNum, rowErrCh)
		// if isValid {
		// 	propMap[propKey] = resultMap
		// }
	}

}

func aggregateRowSuccesses(successCh <-chan int, report *entity.CsvProcessingReport, wg *sync.WaitGroup) {
	var successfulRowNums []int
	for {
		colNum, more := <-successCh
		if more {
			successfulRowNums = append(successfulRowNums, colNum)
			report.Successes = successfulRowNums
		} else {
			wg.Done()
			break
		}
	}
}

func aggregateRowErrors(rowErrCh <-chan entity.RowError, report *entity.CsvProcessingReport, wg *sync.WaitGroup) {
	var rowErrors []entity.RowError
	for {
		rowErr, more := <-rowErrCh
		if more {
			rowErrors = append(rowErrors, rowErr)
			report.RowErrors = rowErrors
		} else {
			wg.Done()
			break
		}
	}
}

func writeJsonFromResults(filePath string, resultCh <-chan map[string]any, wg *sync.WaitGroup) {
	jsonFunc := func(resultMap map[string]any) string {
		jsonData, _ := json.MarshalIndent(resultMap, entity.Indent, entity.Indent)
		return entity.Indent + string(jsonData)
	}
	writeString := CreateStringWriter(filePath)
	writeString("[\n", false)

	first := true
	for {
		resultMap, more := <-resultCh
		if more {
			if !first {
				writeString(",\n", false)
			} else {
				first = false
			}

			jsonData := jsonFunc(resultMap)
			writeString(jsonData, false)
		} else {
			writeString("\n]", true)
			wg.Done()
			break
		}
	}
}
