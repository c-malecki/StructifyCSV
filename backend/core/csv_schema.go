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
	"strconv"
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

func ProcessCsvWithJsonSchema(jsonSchema entity.JsonSchema) entity.CsvProcessingReport {
	var wg sync.WaitGroup
	var csvProcessingErrs []entity.CsvProcessingError
	var successfulRowNums []int

	successCh := make(chan int)
	pErrMapCh := make(chan entity.CsvProcessingError)
	rowMapCh := make(chan map[string]any)

	wg.Add(2)
	go processCsv(jsonSchema, rowMapCh, pErrMapCh, successCh, &wg)
	go writeJsonFromCsv(rowMapCh, &wg)
	wg.Wait()

	for pErr := range pErrMapCh {
		csvProcessingErrs = append(csvProcessingErrs, pErr)
	}
	for rowNum := range successCh {
		successfulRowNums = append(successfulRowNums, rowNum)
	}
	return entity.CsvProcessingReport{Successes: successfulRowNums, Errors: csvProcessingErrs}
}

func processCsv(schema entity.JsonSchema, rowMapCh chan<- map[string]any, pErrMapCh chan<- entity.CsvProcessingError, successCh chan<- int, wg *sync.WaitGroup) {
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
			fmt.Printf("\nRow: %v end\n", rowNum)
			close(rowMapCh)
			close(pErrMapCh)
			close(successCh)
			wg.Done()
			break
		} else if err != nil {
			// return some sort of error for not being able to read csv line
			fmt.Printf("Row: %sError: %s\n", row, err)
		}

		rowNum++

		if len(row) != len(headers) {
			print("row len != header len")
			msg := fmt.Sprintf("row %v doesn't match headers format", rowNum)
			pErr := CreatePErr(msg, rowNum)
			pErrMapCh <- pErr
			continue
		}

		jsonMap := make(map[string]interface{})
		rowSchema := entity.CsvRowSchema{RowNum: rowNum, RowData: row, Properties: schema.Properties}
		rowMap := processRowToMap(rowSchema, jsonMap, pErrMapCh)
		fmt.Printf("\n%v: %+v\n", rowNum, rowMap)

		// if err != nil {
		// 	fmt.Printf("Line: %sError: %s\n", row, err)
		// 	continue
		// }
		successCh <- rowNum
		rowMapCh <- rowMap
	}
}

func writeJsonFromCsv(rowMapCh <-chan map[string]any, wg *sync.WaitGroup) {
	jsonFunc := func(rowMap map[string]any) string {
		jsonData, _ := json.MarshalIndent(rowMap, entity.Indent, entity.Indent)
		return entity.Indent + string(jsonData)
	}
	writeString := CreateStringWriter("/home/meeps/Documents/ProductsModel.json")
	writeString("[\n", false)
	print("write")

	first := true
	for {
		rowMap, more := <-rowMapCh
		fmt.Printf("\n%+v\n", rowMap)
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

func processRowToMap(lineSchema entity.CsvRowSchema, jsonMap map[string]interface{}, pErrMapCh chan<- entity.CsvProcessingError) map[string]any {
	for key, propSchema := range lineSchema.Properties {
		if propSchema.Type != "object" && propSchema.CsvHeaderIndex == nil {
			continue
		}

		switch propSchema.Type {
		case "array":
			indexes := propSchema.CsvHeaderIndex.([]interface{})
			colNums := make([]int, 0, len(indexes))

			for _, v := range indexes {
				i64 := v.(float64)
				index := int(i64)
				colNums = append(colNums, index)
			}

			arr := make([]any, 0, len(colNums))
			for _, num := range colNums {
				v := convertArrayItemType(*propSchema, lineSchema, num, pErrMapCh)
				arr = append(arr, v)
			}
			jsonMap[key] = arr
		case "object":
			nextMap := make(map[string]interface{})
			// if map len > 0
			lineSchema.Properties = propSchema.Properties
			resultMap := processRowToMap(lineSchema, nextMap, pErrMapCh)

			jsonMap[key] = resultMap
		default:
			i64 := propSchema.CsvHeaderIndex.(float64)
			index := int(i64)
			// convert then validate
			v := convertLineItemType(*propSchema, lineSchema, index, pErrMapCh)
			// isValid := validateColumnValue(*propSchema, lineSchema.RowNum, index, v, pErrMapCh)
			// if !isValid {
			// 	continue
			// }
			jsonMap[key] = v
		}
	}
	return jsonMap
}

func validateStringSchema(propSchema entity.SchemaProperty, rowNum int, colNum int, str string, pErrMapCh chan<- entity.CsvProcessingError) bool {
	isValid := true
	if propSchema.MinLength != nil {
		min := int(*propSchema.MinLength)
		if len(str) != min {
			msg := fmt.Sprintf("row %v column %v: string value \"%v\"'s length is less than %v", rowNum, colNum, str, min)
			pErr := CreatePErr(msg, rowNum)
			pErrMapCh <- pErr
			isValid = false
		}
	}
	if propSchema.MaxLength != nil {
		max := int(*propSchema.MaxLength)
		if len(str) != max {
			msg := fmt.Sprintf("row %v column %v: string value \"%v\"'s length is greater than %v", rowNum, colNum, str, max)
			pErr := CreatePErr(msg, rowNum)
			pErrMapCh <- pErr
			isValid = false
		}
	}
	return isValid
}

func validateColumnValue(propSchema entity.SchemaProperty, rowNum int, colNum int, value any, pErrMapCh chan<- entity.CsvProcessingError) bool {
	switch propSchema.Type {
	case "string":
		isValid := validateStringSchema(propSchema, rowNum, colNum, value.(string), pErrMapCh)
		return isValid
	// case "number":
	// handle error for improper value conversion result
	// 	v, err := strconv.ParseFloat(lineSchema.rowData[colNum], 64)
	// 	return v, err
	// case "integer":
	// handle error for improper value conversion result
	// v, err := strconv.Atoi(lineSchema.rowData[colNum])
	// return v, err
	// case "object":
	//  map ??
	// case "array":
	//  slice ??
	// case "boolean":
	// handle error for improper value conversion result
	// 	v, err := strconv.ParseBool(lineSchema.rowData[colNum])
	// 	return v, err
	// case "null":
	// 	return nil, nil
	default:
		msg := fmt.Sprintf("row %v column %v: column value \"%v\" doesn't match accepted data types", rowNum, colNum, value)
		pErr := CreatePErr(msg, rowNum)
		pErr.ColNum = &colNum
		pErrMapCh <- pErr
		return false
	}
}

// add error for conversion
func convertArrayItemType(propSchema entity.SchemaProperty, lineSchema entity.CsvRowSchema, colNum int, pErrMapCh chan<- entity.CsvProcessingError) any {
	var value any
	if propSchema.Items != nil {

		arrayItem := propSchema.Items.(map[string]interface{})
		itemType := arrayItem["type"]
		// itemType := propSchema.Items.(entity.ArrayItems).Type
		switch itemType {
		case "string":
			value = lineSchema.RowData[colNum]
		case "number":
			value, _ = strconv.ParseFloat(lineSchema.RowData[colNum], 64)
		case "integer":
			value, _ = strconv.Atoi(lineSchema.RowData[colNum])
		default:
			value = lineSchema.RowData[colNum]
		}
	}
	return value
}

// add error for conversion
func convertLineItemType(propSchema entity.SchemaProperty, lineSchema entity.CsvRowSchema, colNum int, pErrMapCh chan<- entity.CsvProcessingError) any {
	switch propSchema.Type {
	case "string":
		return lineSchema.RowData[colNum]
	case "number":
		v, _ := strconv.ParseFloat(lineSchema.RowData[colNum], 64)
		return v
	case "integer":
		v, _ := strconv.Atoi(lineSchema.RowData[colNum])
		return v
	case "boolean":
		v, _ := strconv.ParseBool(lineSchema.RowData[colNum])
		return v
	case "null":
		return nil
	default:
		return lineSchema.RowData[colNum]
	}
}
