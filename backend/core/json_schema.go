package core

import (
	"context"
	"csvtoschema/backend/entity"
	"csvtoschema/backend/ui"
	"encoding/json"
	"fmt"
	"os"

	"github.com/santhosh-tekuri/jsonschema"
)

type ImportSchemaRes struct {
	Schema *entity.JsonSchema `json:"schema"`
	Error  error              `json:"error"`
}

func ImportJsonSchema(c context.Context) ImportSchemaRes {
	filePath := ui.PrepareOpenFileDialog(c, "json", "schema")
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

//	func createReqProps(indent string, required []string) string {
//		j, err := json.Marshal(required)
//		if err != nil {
//			print(err)
//		}
//		reqsStr := string(j)
//		return fmt.Sprintf("\n%[1]v},\n%[1]v\"required\": %[2]v", indent, reqsStr)
//	}
func processStringSchema(schema entity.Schema, propertyMap map[string]interface{}) {
	if schema.MinLength != nil {
		propertyMap["minLength"] = schema.MinLength
	}
	if schema.MaxLength != nil {
		propertyMap["maxLength"] = schema.MaxLength
	}
	propertyMap["type"] = "string"
}

func processNumberSchema(schema entity.Schema, propertyMap map[string]interface{}) {
	if schema.NumMinimum != nil {
		propertyMap["minimum"] = schema.NumMinimum
	}
	if schema.NumMaximum != nil {
		propertyMap["maximum"] = schema.NumMaximum
	}
	propertyMap["type"] = "number"
}

func processIntSchema(schema entity.Schema, propertyMap map[string]interface{}) {
	if schema.IntMinimum != nil {
		propertyMap["minimum"] = schema.IntMinimum
	}
	if schema.IntMaximum != nil {
		propertyMap["maximum"] = schema.IntMaximum
	}
	propertyMap["type"] = "integer"
}

type ArrayItems struct {
	Type string `json:"type"`
}

func processArraySchema(schema entity.Schema, propertyMap map[string]interface{}) {
	if schema.MinItems != nil {
		propertyMap["minItems"] = schema.MinItems
	}
	if schema.MaxItems != nil {
		propertyMap["maxItems"] = schema.MaxItems
	}
	if schema.Items != nil {
		propertyMap["items"] = *schema.Items
	}
	propertyMap["type"] = "array"
}

type EmptyStruct struct {
	Empty string `json:"empty,omitempty"`
}

func processObjectSchema(schema entity.Schema, propertyMap map[string]interface{}) {
	if schema.MinProperties != nil {
		propertyMap["minProperties"] = schema.MinProperties
	}
	if schema.MaxProperties != nil {
		propertyMap["maxProperties"] = schema.MaxProperties
	}
	if schema.Required != nil {
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

func WriteJsonSchema(c context.Context, schema entity.JsonSchema) {
	outPath := ui.PrepareSaveFileDialog(c, schema.Title)

	writeString := CreateStringWriter(outPath)
	boiler := createSchemaBoiler(schema)
	writeString(boiler, false)

	schemaMap := processSchemaProperties(schema.Properties)

	jsonBytes, err := json.MarshalIndent(schemaMap, entity.Indent, entity.Indent)
	if err != nil {
		print(err)
	}
	writeString(string(jsonBytes), false)

	// reqs := []string{"test", "the", "reqs"}
	// required := createReqProps(indent, reqs)
	// writeString(required, false)

	writeString("\n}", true)
}
