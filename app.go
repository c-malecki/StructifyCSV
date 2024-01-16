package main

import (
	"context"
	"csvtoschema/backend/core"
	"csvtoschema/backend/entity"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) ImportSchema() entity.JsonSchema {
	return core.ImportSchemaJson(a.ctx)
}

func (a *App) ImportCsvData() entity.CsvData {
	return core.ImportCsvFileData(a.ctx)
}

func (a *App) ExportSchema(schema entity.JsonSchema) {
	core.WriteJsonSchema(a.ctx, schema)
}

func (a *App) ProcessCsvDescriptor(schema entity.JsonSchema, hd []entity.HeaderDescriptor, csvPath string) {
	core.ExportCsvDescriptor(a.ctx, schema, hd, csvPath)
}
