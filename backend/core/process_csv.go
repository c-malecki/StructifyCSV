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
	go writeJsonFromResults(resultCh, &wg)
	go buildReportSuccesses(successCh, &report, &wg)
	go buildReportFailures(pErrCh, &report, &wg)
	wg.Wait()

	return report
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

		// need to check required properties
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
	jsonFunc := func(resultMap map[string]any) string {
		jsonData, _ := json.MarshalIndent(resultMap, entity.Indent, entity.Indent)
		return entity.Indent + string(jsonData)
	}
	writeString := CreateStringWriter("/home/meeps/Documents/ProductsModel.json")
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

func processCsvRowToMap(rowSchema entity.CsvRowSchema, jsonMap map[string]interface{}, pErrCh chan<- entity.CsvProcessingError) map[string]any {
	for key, propSchema := range rowSchema.Properties {
		if propSchema.Type != "object" && propSchema.CsvHeaderIndex == nil {
			continue
		}

		switch propSchema.Type {
		case "string":
			i64 := propSchema.CsvHeaderIndex.(float64)
			colNum := int(i64)
			v := rowSchema.RowData[colNum]
			isValid := ValidateStringProperty(*propSchema, rowSchema.RowNum, colNum, v, pErrCh)
			if isValid {
				jsonMap[key] = v
			}
		case "number":
			i64 := propSchema.CsvHeaderIndex.(float64)
			colNum := int(i64)
			v, ok := ConvertToFloat(*propSchema, rowSchema, colNum, pErrCh)
			if ok {
				isValid := ValidateNumberProperty(*propSchema, rowSchema.RowNum, colNum, v, pErrCh)
				if isValid {
					jsonMap[key] = v
				}
			}
		case "integer":
			i64 := propSchema.CsvHeaderIndex.(float64)
			colNum := int(i64)
			v, ok := ConvertToInt(*propSchema, rowSchema, colNum, pErrCh)
			if ok {
				isValid := ValidateIntegerProperty(*propSchema, rowSchema.RowNum, colNum, v, pErrCh)
				if isValid {
					jsonMap[key] = v
				}
			}
		case "boolean":
			i64 := propSchema.CsvHeaderIndex.(float64)
			colNum := int(i64)
			v, ok := ConvertToBool(*propSchema, rowSchema, colNum, pErrCh)
			if ok {
				jsonMap[key] = v
			}
		case "null":
			jsonMap[key] = nil
		case "array":
			colNums := ConvertIndexArrToInts(propSchema.CsvHeaderIndex)

			arrItems := make([]any, 0, len(colNums))
			for _, colNum := range colNums {
				v, ok := ConvertArrayItemType(*propSchema, rowSchema, colNum, pErrCh)
				if ok {
					arrItems = append(arrItems, v)
				}
			}

			isValid := ValidateArrayProperty(*propSchema, rowSchema.RowNum, arrItems, pErrCh)
			if isValid {
				jsonMap[key] = arrItems
			}
		case "object":
			nextMap := make(map[string]interface{})
			// if map len > 0
			rowSchema.Properties = propSchema.Properties
			resultMap := processCsvRowToMap(rowSchema, nextMap, pErrCh)

			isValid := ValidateObjectProperty(*propSchema, "", rowSchema.RowNum, pErrCh)
			if isValid {
				jsonMap[key] = resultMap
			}
		}
	}
	return jsonMap
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
