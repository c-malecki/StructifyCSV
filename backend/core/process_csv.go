package core

import (
	"context"
	"csvtoschema/backend/entity"
	"csvtoschema/backend/ui"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
)

func ImportCsvFileData(c context.Context) entity.CsvFileData {
	filePath := ui.PrepareOpenFileDialog(c, "csv", "")
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	defer file.Close()

	name := filepath.Base(filePath)

	reader := csv.NewReader(file)
	headers, err := reader.Read()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	return entity.CsvFileData{FileName: name, Location: filePath, Headers: headers}
}

func ProcessCsv(jsonSchema entity.JsonSchema) entity.CsvProcessingReport {
	var wg sync.WaitGroup
	var report entity.CsvProcessingReport

	successCh := make(chan int)
	pErrCh := make(chan entity.CsvProcessingError)
	resultCh := make(chan map[string]any)

	wg.Add(4)
	go processCsvWithJsonSchema(jsonSchema, resultCh, pErrCh, successCh, &wg)
	go buildReportSuccesses(successCh, &report, &wg)
	go buildReportFailures(pErrCh, &report, &wg)
	go writeJsonFromResults(resultCh, &wg)
	wg.Wait()

	return report
}

func buildReportSuccesses(successCh <-chan int, report *entity.CsvProcessingReport, wg *sync.WaitGroup) {
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

func buildReportFailures(pErrCh <-chan entity.CsvProcessingError, report *entity.CsvProcessingReport, wg *sync.WaitGroup) {
	var csvProcessingErrs []entity.CsvProcessingError
	for {
		pErr, more := <-pErrCh
		if more {
			csvProcessingErrs = append(csvProcessingErrs, pErr)
			report.Errors = csvProcessingErrs
		} else {
			wg.Done()
			break
		}
	}
}

func processCsvWithJsonSchema(schema entity.JsonSchema, resultCh chan<- map[string]any, pErrCh chan<- entity.CsvProcessingError, successCh chan<- int, wg *sync.WaitGroup) {
	file, err := os.Open("/home/meeps/Documents/Products.csv")
	if err != nil {
		// return error that file could not be opened
		fmt.Printf("Error: %s\n", err)
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
			close(pErrCh)
			close(successCh)
			wg.Done()
			break
		} else if err != nil {
			// return some sort of error for not being able to read csv line
			fmt.Printf("Row: %sError: %s\n", row, err)
		}

		rowNum++

		if len(row) != len(headers) {
			pErr := CreateHeaderMismatchErr(rowNum, len(row), len(headers))
			pErrCh <- pErr
			continue
		}

		jsonMap := make(map[string]interface{})
		rowSchema := entity.CsvRowSchema{RowNum: rowNum, RowData: row, Properties: schema.Properties}
		rowMap := processCsvRowToMap(rowSchema, jsonMap, pErrCh)

		// if err != nil {
		// 	fmt.Printf("Line: %sError: %s\n", row, err)
		// 	continue
		// }
		successCh <- rowNum
		resultCh <- rowMap
	}
}

func writeJsonFromResults(resultCh <-chan map[string]any, wg *sync.WaitGroup) {
	jsonFunc := func(rowMap map[string]any) string {
		jsonData, _ := json.MarshalIndent(rowMap, entity.Indent, entity.Indent)
		return entity.Indent + string(jsonData)
	}
	writeString := CreateStringWriter("/home/meeps/Documents/ProductsModel.json")
	writeString("[\n", false)

	first := true
	for {
		rowMap, more := <-resultCh
		if more {
			if !first {
				writeString(",\n", false)
			} else {
				first = false
			}

			jsonData := jsonFunc(rowMap)
			writeString(jsonData, false)
		} else {
			writeString("\n]", true)
			wg.Done()
			break
		}
	}
}

func processCsvRowToMap(lineSchema entity.CsvRowSchema, jsonMap map[string]interface{}, pErrCh chan<- entity.CsvProcessingError) map[string]any {
	for key, propSchema := range lineSchema.Properties {
		if propSchema.Type != "object" && propSchema.CsvHeaderIndex == nil {
			continue
		}

		switch propSchema.Type {
		case "array":
			colNums := getArrayItemsColNums(propSchema.CsvHeaderIndex)

			arrItems := make([]any, 0, len(colNums))
			for _, num := range colNums {
				v, isValid := ConvertArrayItemType(*propSchema, lineSchema, num, pErrCh)
				if isValid {
					arrItems = append(arrItems, v)
				}
			}
			// validate array schema min max
			jsonMap[key] = arrItems
		case "object":
			nextMap := make(map[string]interface{})
			// if map len > 0
			lineSchema.Properties = propSchema.Properties
			resultMap := processCsvRowToMap(lineSchema, nextMap, pErrCh)

			jsonMap[key] = resultMap
		default:
			i64 := propSchema.CsvHeaderIndex.(float64)
			colNum := int(i64)
			v, isValid := ConvertColumnValueType(*propSchema, lineSchema, colNum, pErrCh)
			if isValid {
				// validate schema type min max etc...
				jsonMap[key] = v
			}
		}
	}
	return jsonMap
}

func getArrayItemsColNums(csvHeaderIndex interface{}) []int {
	indexes := csvHeaderIndex.([]interface{})
	colNums := make([]int, 0, len(indexes))
	for _, v := range indexes {
		i64 := v.(float64)
		index := int(i64)
		colNums = append(colNums, index)
	}
	return colNums
}
