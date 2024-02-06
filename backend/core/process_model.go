package core

import (
	"context"
	"csvtoschema/backend/entity"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/mitchellh/mapstructure"
)

func ExportCsvModelJson(c context.Context, schema entity.JsonSchema, model entity.CsvModel) {
	lineMapCh := make(chan map[string]any)
	doneCh := make(chan bool)

	go processCsvModel(schema, model, lineMapCh)
	go writeModelJson(lineMapCh, doneCh)
	<-doneCh
}

func processCsvModel(schema entity.JsonSchema, model entity.CsvModel, lineMapCh chan<- map[string]any) {
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

		record, err := processCsvLineToMap(headers, line, model)

		if err != nil {
			fmt.Printf("Line: %sError: %s\n", line, err)
			continue
		}
		lineMapCh <- record
	}
}

func parseModelChildMap(child map[string]interface{}, lineData []string, jsonMap map[string]interface{}, lastKey string) {
	nextMap := make(map[string]interface{})

	for key, value := range child {
		var nodeStruct entity.CsvModelNodeValue
		grandChildMap := value.(map[string]interface{})

		_, ok := grandChildMap["headerIdx"]
		if ok {
			err := mapstructure.Decode(grandChildMap, &nodeStruct)
			if err != nil {
				print(err)
			}
			if nodeStruct.HeaderIdx != nil {
				v, _ := convertLineDataType(lineData[*nodeStruct.HeaderIdx], nodeStruct)
				nextMap[key] = v
				jsonMap[lastKey] = nextMap
			}
		} else {
			jsonMap[lastKey] = nextMap
			parseModelChildMap(grandChildMap, lineData, jsonMap[lastKey].(map[string]interface{}), key)
		}
	}
}

func processCsvLineToMap(headers []string, lineData []string, model entity.CsvModel) (map[string]any, error) {
	if len(lineData) != len(headers) {
		return nil, errors.New("line doesn't match headers format. skipping")
	}

	jsonMap := make(map[string]interface{})

	for key, value := range model.Map {
		var nodeStruct entity.CsvModelNodeValue
		nodeMap := value.(map[string]interface{})

		_, ok := nodeMap["headerIdx"]
		if ok {
			err := mapstructure.Decode(nodeMap, &nodeStruct)
			if err != nil {
				print(err)
			}
			if nodeStruct.HeaderIdx != nil {
				v, _ := convertLineDataType(lineData[*nodeStruct.HeaderIdx], nodeStruct)
				jsonMap[key] = v
			}
		} else {
			parseModelChildMap(nodeMap, lineData, jsonMap, key)
		}
	}

	return jsonMap, nil
}

func convertLineDataType(lineItem string, node entity.CsvModelNodeValue) (any, error) {
	switch node.DataType {
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
	writeString := createStringWriter("/home/meeps/Documents/ProductsModel.json")
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
