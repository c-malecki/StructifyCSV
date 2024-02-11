package core

import (
	"context"
	"csvtoschema/backend/entity"
	"csvtoschema/backend/ui"
	"encoding/json"
	"fmt"
	"os"
	"strings"

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
		fmt.Printf("%v", schemaErr)
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
		print("unmarhsal error")
	}
	result.Schema = &schema
	result.Error = nil
	return result
}

func createSchemaBoiler(schema entity.JsonSchema) string {
	title := schema.Title
	description := schema.Description
	// todo: doesn't account for anything other than " " literal
	replaceInner := strings.ReplaceAll(title, " ", "_")
	trim := strings.TrimSpace(replaceInner)
	idText := strings.ToLower(trim)
	return fmt.Sprintf("{\n%[1]v\"$schema\": \"https://json-schema.org/draft/2020-12/schema\",\n%[1]v\"$id\": \"https://example.com/%[3]v.schema.json\",\n%[1]v\"title\": \"%[2]v\",\n%[1]v\"description\": \"%[4]v\",\n%[1]v\"type\": \"object\",\n%[1]v\"properties\": ", entity.Indent, title, idText, description)
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

	writeString := CreateStringWriter(outPath)
	boiler := createSchemaBoiler(schema)
	writeString(boiler, false)

	jsonBytes, err := json.MarshalIndent(schema.Properties, entity.Indent, entity.Indent)
	if err != nil {
		print(err)
	}
	writeString(string(jsonBytes), false)

	// reqs := []string{"test", "the", "reqs"}
	// required := createReqProps(indent, reqs)
	// writeString(required, false)

	writeString("\n}", true)
}
