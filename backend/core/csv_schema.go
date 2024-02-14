package core

import (
	"context"
	"csvtoschema/backend/entity"
	"csvtoschema/backend/ui"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
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

func ProcessCsvWithSchema(c context.Context, schema entity.JsonSchema) {
	lineMapCh := make(chan map[string]any)
	doneCh := make(chan bool)

	go processCsvModel(schema, lineMapCh)
	go writeModelJson(lineMapCh, doneCh)
	<-doneCh
}

func processCsvModel(schema entity.JsonSchema, lineMapCh chan<- map[string]any) {
	file, err := os.Open("/home/meeps/Documents/Products.csv")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	var headers, line []string
	headers, err = reader.Read()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	for {
		line, err = reader.Read()

		if err == io.EOF {
			close(lineMapCh)
			break
		} else if err != nil {
			fmt.Printf("Line: %sError: %s\n", line, err)
		}

		if len(line) != len(headers) {
			err := errors.New("line doesn't match headers format. skipping")
			fmt.Printf("Error: %s\n", err)
		}

		jsonMap := make(map[string]interface{})
		record := processCsvLineToMap(line, schema.Properties, jsonMap)

		if err != nil {
			fmt.Printf("Line: %sError: %s\n", line, err)
			continue
		}
		lineMapCh <- record
	}
}

func processCsvLineToMap(lineData []string, properties entity.Properties, jsonMap map[string]interface{}) map[string]any {
	for key, schema := range properties {
		if schema.Type != "object" && schema.CsvHeaderIndex == nil {
			continue
		}

		switch schema.Type {
		case "array":
			arr := schema.CsvHeaderIndex.([]interface{})
			indexes := make([]int, 0, len(arr))

			for _, v := range arr {
				i64 := v.(float64)
				index := int(i64)
				indexes = append(indexes, index)
			}

			slice := make([]any, 0, len(indexes))
			for _, i := range indexes {
				v := lineData[i]
				slice = append(slice, v)
			}

			jsonMap[key] = slice
		case "object":
			nextMap := make(map[string]interface{})
			jsonMap[key] = processCsvLineToMap(lineData, schema.Properties, nextMap)
		default:
			i64 := schema.CsvHeaderIndex.(float64)
			index := int(i64)
			v, _ := convertLineDataType(lineData[index], *schema)
			jsonMap[key] = v
		}
	}
	return jsonMap
}

// func createSliceOfType( propType string, len int) []any {
// 	switch propType {
// 	case "string":
// 		slice := make([]string, len)
// 		return slice
// 	}
// }

func convertLineDataType(lineItem string, schema entity.SchemaProperty) (any, error) {
	switch schema.Type {
	case "string":
		//  string
		return lineItem, nil
	case "number":
		// handle error for improper value conversion result
		v, err := strconv.ParseFloat(lineItem, 64)
		return v, err
	case "integer":
		// handle error for improper value conversion result
		v, err := strconv.Atoi(lineItem)
		return v, err
	// case "object":
	//  map ??
	// case "array":
	//  slice ??
	case "boolean":
		// handle error for improper value conversion result
		v, err := strconv.ParseBool(lineItem)
		return v, err
	case "null":
		return nil, nil
	default:
		return lineItem, errors.New("schema property value data type doesn't match accepted data types")
	}
}

func writeModelJson(lineMapCh <-chan map[string]any, done chan<- bool) {
	jsonFunc := func(lineMap map[string]any) string {
		jsonData, _ := json.MarshalIndent(lineMap, entity.Indent, entity.Indent)
		return entity.Indent + string(jsonData)
	}
	writeString := CreateStringWriter("/home/meeps/Documents/ProductsModel.json")
	writeString("[\n", false)

	first := true
	for {
		lineMap, more := <-lineMapCh
		if more {
			if !first {
				writeString(",\n", false)
			} else {
				first = false
			}

			jsonData := jsonFunc(lineMap)
			writeString(jsonData, false)
		} else {
			writeString("\n]", true)
			done <- true
			break
		}
	}
}
