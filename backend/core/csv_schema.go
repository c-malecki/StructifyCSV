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

	"github.com/mitchellh/mapstructure"
)

func ImportCsvFileData(c context.Context) entity.CsvFileData {
	filePath := ui.PrepareOpenFileDialog(c, "csv", "")
	file, err := os.Open(filePath)
	if err != nil {
		print(err)
	}
	defer file.Close()

	name := filepath.Base(filePath)

	reader := csv.NewReader(file)
	headers, err := reader.Read()
	if err != nil {
		print(err)
	}
	return entity.CsvFileData{FileName: name, Location: filePath, Headers: headers}
}

func WriteJsonFromCsvModelMap(c context.Context, modelMap entity.CsvModelMap) {
	lineMapCh := make(chan map[string]any)
	doneCh := make(chan bool)

	go processCsvModel(modelMap, lineMapCh)
	go writeModelJson(lineMapCh, doneCh)
	<-doneCh
}

func processCsvModel(modelMap entity.CsvModelMap, lineMapCh chan<- map[string]any) {
	file, err := os.Open("/home/meeps/Documents/Products.csv")
	if err != nil {
		print(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	var headers, line []string
	headers, err = reader.Read()
	if err != nil {
		print(err)
	}

	for {
		line, err = reader.Read()

		if err == io.EOF {
			close(lineMapCh)
			break
		} else if err != nil {
			print(err)
		}

		record, err := processCsvLineToMap(headers, line, modelMap)

		if err != nil {
			fmt.Printf("Line: %sError: %s\n", line, err)
			continue
		}
		lineMapCh <- record
	}
}

func processCsvLineToMap(headers []string, lineData []string, modelMap entity.CsvModelMap) (map[string]any, error) {
	if len(lineData) != len(headers) {
		return nil, errors.New("line doesn't match headers format. skipping")
	}

	jsonMap := make(map[string]interface{})

	for key, value := range modelMap {
		var nodeStruct entity.CsvSchemaProperty
		nodeMap := value.(map[string]interface{})

		_, ok := nodeMap["headerIndexes"]
		if ok {
			err := mapstructure.Decode(nodeMap, &nodeStruct)
			if err != nil {
				print(err)
			}

			if len(nodeStruct.HeaderIndexes) > 0 {
				if nodeStruct.SchemaPropertyType == "array" {
					slice := make([]any, 0, len(nodeStruct.HeaderIndexes))
					for _, n := range nodeStruct.HeaderIndexes {
						v := lineData[n]
						slice = append(slice, v)
					}
					jsonMap[key] = slice
				} else {
					v, _ := convertLineDataType(lineData[nodeStruct.HeaderIndexes[0]], nodeStruct)
					jsonMap[key] = v
				}
			}

		} else {
			parseModelChildMap(nodeMap, lineData, jsonMap, key)
		}
	}

	return jsonMap, nil
}

// todo: bug - if nested maps are empty, json output is an empty object
func parseModelChildMap(child map[string]interface{}, lineData []string, jsonMap map[string]interface{}, lastKey string) {
	nextMap := make(map[string]interface{})

	for key, value := range child {
		var nodeStruct entity.CsvSchemaProperty
		grandChildMap := value.(map[string]interface{})

		_, ok := grandChildMap["headerIndexes"]
		if ok {
			err := mapstructure.Decode(grandChildMap, &nodeStruct)
			if err != nil {
				print(err)
			}

			if len(nodeStruct.HeaderIndexes) > 0 {
				if nodeStruct.SchemaPropertyType == "array" {
					slice := make([]any, 0, len(nodeStruct.HeaderIndexes))
					for _, n := range nodeStruct.HeaderIndexes {
						v := lineData[n]
						slice = append(slice, v)
					}
					nextMap[key] = slice
					jsonMap[key] = nextMap
				} else {
					v, _ := convertLineDataType(lineData[nodeStruct.HeaderIndexes[0]], nodeStruct)
					nextMap[key] = v
					jsonMap[lastKey] = nextMap
				}
			}

		} else {
			jsonMap[lastKey] = nextMap
			parseModelChildMap(grandChildMap, lineData, jsonMap[lastKey].(map[string]interface{}), key)
		}
	}
}

// func createSliceOfType( propType string, len int) []any {
// 	switch propType {
// 	case "string":
// 		slice := make([]string, len)
// 		return slice
// 	}
// }

func convertLineDataType(lineItem string, node entity.CsvSchemaProperty) (any, error) {
	switch node.SchemaPropertyType {
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
	// case "null":
	//  nil ??
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
