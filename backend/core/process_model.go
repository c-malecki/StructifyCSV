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
)

func ExportCsvModelJson(c context.Context, schema entity.JsonSchema, model entity.CsvModel) {
	lineMapCh := make(chan map[string]any)
	doneCh := make(chan bool)

	go processCsvModel(schema, model, lineMapCh)
	go writeModelJson(lineMapCh, doneCh)
	<-doneCh
}

// writing csv to json from model
// open csv and get headers again
// only use selected headers
// ignore map properties in model without an assigned header
// rescursively step through property values to write json
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

		record, err := processCsvLineToMap(headers, line, model.HeaderDescriptors)

		if err != nil {
			fmt.Printf("Line: %sError: %s\n", line, err)
			continue
		}
		lineMapCh <- record
	}
}

func processCsvLineToMap(headers []string, lineData []string, headerDescriptors []entity.HeaderDescriptor) (map[string]any, error) {
	if len(lineData) != len(headers) {
		return nil, errors.New("line doesn't match headers format. skipping")
	}

	lineMap := make(map[string]any)

	for _, hd := range headerDescriptors {
		if hd.SchemaProperty == nil {
			continue
		}
		v, err := convertLineDataType(lineData[hd.HeaderIndex], *hd.SchemaProperty)
		if err != nil {
			return nil, err
		}
		lineMap[headers[hd.HeaderIndex]] = v
	}

	return lineMap, nil
}

func convertLineDataType(lineItem string, schemaProperty entity.SchemaProperty) (any, error) {
	switch schemaProperty.Value {
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
		jsonData, _ := json.MarshalIndent(lineMap, "  ", "  ")
		return "  " + string(jsonData)
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

// func processLineMapToModel() {

// }

// func processLineWithModel(model map[string]interface{}) {
// 	for k,v := range model {
// 		header, ok := v["csvheader"]
// 		if ok {
// 			dataType := v["dataType"]
// 		}

// 	}
// }

// func coerceVal(dataType string, lineItem interface{}) {
// 	switch dataType {
// 	case "string":
//  string

// 	case "number":
//  float

// 	case "integer":
//  int

// 	case "object":
//  map ??

// 	case "array":
//  slice ??

// 	case "boolean":
//  boolean

// 	case "null":
//  nil ??

// 	}
// }
// for line in csv
// for [k,v] in model
// if v is map, recur
// else v.get(csvHeader)
// if csvHeader != nil
// assert or convert csv value
