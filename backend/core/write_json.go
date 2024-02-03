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
	"strings"
)

func createStringWriter(filePath string) func(string, bool) {
	file, err := os.Create(filePath)
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

func createSchemaBoiler(indent string, schema entity.JsonSchema) string {
	title := schema.Title
	description := schema.Description
	// todo: doesn't account for anything other than " " literal
	replaceInner := strings.ReplaceAll(title, " ", "_")
	trim := strings.TrimSpace(replaceInner)
	idText := strings.ToLower(trim)
	return fmt.Sprintf("{\n%[1]v\"$schema\": \"https://json-schema.org/draft/2020-12/schema\",\n%[1]v\"$id\": \"https://example.com/%[3]v.schema.json\",\n%[1]v\"title\": \"%[2]v\",\n%[1]v\"description\": \"%[4]v\",\n%[1]v\"type\": \"object\",\n%[1]v\"properties\": ", indent, title, idText, description)
}

// func createReqProps(indent string, required []string) string {
// 	j, err := json.Marshal(required)
// 	if err != nil {
// 		print(err)
// 	}
// 	reqsStr := string(j)
// 	return fmt.Sprintf("\n%[1]v},\n%[1]v\"required\": %[2]v", indent, reqsStr)
// }

func WriteJsonSchema(c context.Context, schema entity.JsonSchema) {
	outPath := ui.PrepareSaveFileDialog(c, schema.Title)

	const indent = "    "
	writeString := createStringWriter(outPath)
	boiler := createSchemaBoiler(indent, schema)
	writeString(boiler, false)

	fmt.Printf("%v", schema.Properties)

	jsonBytes, err := json.MarshalIndent(schema.Properties, indent, indent)
	if err != nil {
		print(err)
	}
	writeString(string(jsonBytes), false)

	// reqs := []string{"test", "the", "reqs"}
	// required := createReqProps(indent, reqs)

	// go channel to write each

	// write each property in "properties": { }

	// writeString(required, false)

	writeString("\n}", true)
}

func ExportCsvDescriptor(c context.Context, schema entity.JsonSchema, hd []entity.HeaderDescriptor, csvPath string) {
	filePath := ui.PrepareSaveFileDialog(c, schema.Title)
	writerCh := make(chan map[string]string)
	doneCh := make(chan bool)

	go ProcessCsvFile(csvPath, hd, writerCh)
	go WriteDescriptorJson(filePath, csvPath, writerCh, doneCh)
	<-doneCh
}

// writing csv to json from model
// open csv and get headers again
// only use selected headers
// ignore map properties in model without an assigned header
// rescursively step through property values to write json
func WriteJsonFromModel(c context.Context, selectedHeaders []string, model map[string]interface{}) {
	file, err := os.Open("/home/meeps/Documents/Products.csv")
	if err != nil {
		print(err)
	}
	defer file.Close()

	writerCh := make(chan map[string]string)
	// doneCh := make(chan bool)

	reader := csv.NewReader(file)

	var headers, line []string
	headers, err = reader.Read()
	if err != nil {
		print(err)
	}

	var usedIndexes []int
	for i, h := range headers {
		for _, sh := range selectedHeaders {
			if h == sh {
				usedIndexes = append(usedIndexes, i)
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

		if len(line) != len(headers) {
			// return nil, errors.New("line doesn't match headers format. skipping")
			continue
		}

		record, err := processCsvLine(headers, line, usedIndexes)

		if err != nil {
			fmt.Printf("Line: %sError: %s\n", line, err)
			continue
		}
		writerCh <- record
	}
}

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
