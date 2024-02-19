package core

import (
	"StructifyCSV/backend/entity"
	"StructifyCSV/backend/ui"
	"context"
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
	rowErrCh := make(chan entity.RowError)
	resultCh := make(chan map[string]any)

	wg.Add(4)
	go processCsvWithJsonSchema(jsonSchema, resultCh, rowErrCh, successCh, &wg)
	go writeJsonFromResults(resultCh, &wg)
	go aggregateRowSuccesses(successCh, &report, &wg)
	go aggregateRowErrors(rowErrCh, &report, &wg)
	wg.Wait()

	return report
}

func processCsvWithJsonSchema(jsonSchema entity.JsonSchema, resultCh chan<- map[string]any, rowErrCh chan<- entity.RowError, successCh chan<- int, wg *sync.WaitGroup) {
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

		// need to check required properties
		// jsonMap := make(map[string]interface{})
		// rowSchema := entity.CsvRowSchema{RowNum: rowNum, RowData: row, Properties: schema.Properties}
		// rowMap := processCsvRowToMap(rowSchema, jsonMap, rowErrCh)

		// if err != nil {
		// 	fmt.Printf("Line: %sError: %s\n", row, err)
		// 	continue
		// }

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
