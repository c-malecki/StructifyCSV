package core

import (
	"context"
	"csvtoschema/backend/entity"
	"csvtoschema/backend/ui"
	"encoding/json"
	"fmt"
	"strings"
)

func createSchemaBoiler(indent string, schema entity.JsonSchema) string {
	title := schema.Title
	description := schema.Description
	// todo: doesn't account for anything other than " " literal
	replaceInner := strings.ReplaceAll(title, " ", "_")
	trim := strings.TrimSpace(replaceInner)
	idText := strings.ToLower(trim)
	return fmt.Sprintf("{\n%[1]v\"$schema\": \"https://json-schema.org/draft/2020-12/schema\",\n%[1]v\"$id\": \"https://example.com/%[3]v.schema.json\",\n%[1]v\"title\": \"%[2]v\",\n%[1]v\"description\": \"%[4]v\",\n%[1]v\"type\": \"object\",\n%[1]v\"properties\": {\n", indent, title, idText, description)
}

func createReqProps(indent string, required []string) string {
	j, err := json.Marshal(required)
	if err != nil {
		print(err)
	}
	reqsStr := string(j)
	return fmt.Sprintf("\n%[1]v},\n%[1]v\"required\": %[2]v", indent, reqsStr)
}

func WriteJsonSchema(c context.Context, schema entity.JsonSchema) {
	outPath := ui.PrepareSaveFileDialog(c, schema.Title)
	const indent = "  "
	writeString := createStringWriter(outPath)
	boiler := createSchemaBoiler(indent, schema)
	writeString(boiler, false)
	reqs := []string{"test", "the", "reqs"}
	required := createReqProps(indent, reqs)
	// write each property in "properties": { }
	writeString(required, false)
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
