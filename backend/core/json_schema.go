package core

import (
	"StructifyCSV/backend/entity"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/santhosh-tekuri/jsonschema"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type ImportSchemaRes struct {
	Schema *entity.JsonSchema `json:"schema"`
	Error  error              `json:"error"`
}

func ImportJsonSchema(c context.Context) ImportSchemaRes {
	opts := runtime.OpenDialogOptions{
		Title: "Import JSON Schema",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "JSON Schema (*.schema.json)",
				Pattern:     "*.schema.json",
			},
		},
	}

	filePath, err := runtime.OpenFileDialog(c, opts)

	if err != nil {
		fmt.Printf("File Path Error: %s\n", err)
	}

	var result ImportSchemaRes
	var schema entity.JsonSchema

	if len(filePath) == 0 {
		result.Schema = nil
		result.Error = nil
		return result
	}

	_, schemaErr := jsonschema.Compile(filePath)

	if schemaErr != nil {
		print("compile schema error")
		result.Schema = nil
		result.Error = schemaErr
		return result
	}

	b, fileErr := os.ReadFile(filePath)
	if fileErr != nil {
		print("read file error")
		result.Schema = nil
		result.Error = fileErr
		return result
	}

	jsonErr := json.Unmarshal(b, &schema)

	if jsonErr != nil {
		fmt.Println(jsonErr)
		print("unmarhsal error")
		result.Schema = nil
		result.Error = schemaErr
	}
	result.Schema = &schema
	result.Error = nil
	return result
}

func ExportJsonSchema(c context.Context, jsonSchema entity.JsonSchema) {
	suggestedName := strings.ReplaceAll(jsonSchema.Title, " ", "_")
	opts := runtime.SaveDialogOptions{
		DefaultDirectory: ".",
		DefaultFilename:  suggestedName + ".schema.json",
		Title:            "Save JSON Schema",
	}

	filePath, err := runtime.SaveFileDialog(c, opts)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	writeString := CreateStringWriter(filePath)
	boiler := createSchemaBoiler(jsonSchema)
	writeString(boiler, false)

	schemaMap := processSchemaProperties(jsonSchema.Properties)

	jsonBytes, err := json.MarshalIndent(schemaMap, entity.Indent, entity.Indent)
	if err != nil {
		print(err)
	}
	writeString(string(jsonBytes), false)

	if jsonSchema.Required != nil && len(jsonSchema.Required) > 0 {
		required := createReqProps(jsonSchema.Required)
		writeString(required, false)
	}

	writeString("\n}", true)
}

func createSchemaBoiler(schema entity.JsonSchema) string {
	title := schema.Title
	description := schema.Description
	// todo: doesn't account for anything other than " " literal
	// replaceInner := strings.ReplaceAll(title, " ", "_")
	// trim := strings.TrimSpace(replaceInner)
	// idText := strings.ToLower(trim)
	// \n%[1]v\"$id\": \"https://example.com/%[3]v.schema.json\",
	return fmt.Sprintf("{\n%[1]v\"$schema\": \"http://json-schema.org/draft-07/schema#\",\n%[1]v\"title\": \"%[2]v\",\n%[1]v\"description\": \"%[3]v\",\n%[1]v\"type\": \"object\",\n%[1]v\"properties\": ", entity.Indent, title, description)
}

func createReqProps(required []string) string {
	j, err := json.Marshal(required)
	if err != nil {
		print(err)
	}
	reqsStr := string(j)
	return fmt.Sprintf(",\n%[1]v\"required\": %[2]v", entity.Indent, reqsStr)
}

func processStringSchema(schema entity.PropertySchema, propertyMap map[string]interface{}) {
	if schema.MinLength != nil {
		propertyMap["minLength"] = schema.MinLength
	}
	if schema.MaxLength != nil {
		propertyMap["maxLength"] = schema.MaxLength
	}
	propertyMap["type"] = "string"
}

func processNumberSchema(schema entity.PropertySchema, propertyMap map[string]interface{}) {
	if schema.NumMinimum != nil {
		propertyMap["minimum"] = schema.NumMinimum
	}
	if schema.NumMaximum != nil {
		propertyMap["maximum"] = schema.NumMaximum
	}
	propertyMap["type"] = "number"
}

func processIntSchema(schema entity.PropertySchema, propertyMap map[string]interface{}) {
	if schema.IntMinimum != nil {
		propertyMap["minimum"] = schema.IntMinimum
	}
	if schema.IntMaximum != nil {
		propertyMap["maximum"] = schema.IntMaximum
	}
	propertyMap["type"] = "integer"
}

func processArraySchema(schema entity.PropertySchema, propertyMap map[string]interface{}) {
	if schema.MinItems != nil {
		propertyMap["minItems"] = schema.MinItems
	}
	if schema.MaxItems != nil {
		propertyMap["maxItems"] = schema.MaxItems
	}
	if schema.Items != nil {
		// ex: { items: { type: "string" } }
		propertyMap["items"] = schema.Items
	}
	propertyMap["type"] = "array"
}

type EmptyStruct struct {
	Empty string `json:"empty,omitempty"`
}

func processObjectSchema(schema entity.PropertySchema, propertyMap map[string]interface{}) {
	if schema.MinProperties != nil {
		propertyMap["minProperties"] = schema.MinProperties
	}
	if schema.MaxProperties != nil {
		propertyMap["maxProperties"] = schema.MaxProperties
	}
	if schema.Required != nil && len(schema.Required) > 0 {
		propertyMap["required"] = schema.Required
	}
	if schema.Properties != nil {
		propertyMap["properties"] = processSchemaProperties(schema.Properties)
	} else {
		s := &EmptyStruct{}
		propertyMap["properties"] = s
	}
	propertyMap["type"] = "object"
}

func processSchemaProperties(properties entity.Properties) map[string]interface{} {
	schemaMap := make(map[string]interface{})
	for k, v := range properties {
		propertyMap := make(map[string]interface{})
		switch v.Type {
		case "string":
			processStringSchema(*v, propertyMap)
		case "number":
			processNumberSchema(*v, propertyMap)
		case "integer":
			processIntSchema(*v, propertyMap)
		case "boolean":
			propertyMap["type"] = "boolean"
		case "null":
			propertyMap["type"] = "null"
		case "array":
			processArraySchema(*v, propertyMap)
		case "object":
			processObjectSchema(*v, propertyMap)
		}
		schemaMap[k] = propertyMap
	}
	return schemaMap
}
