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

func (a *App) ImportJsonSchema() entity.JsonSchema {
	return core.ImportJsonSchema(a.ctx)
}

func (a *App) ImportCsvData() entity.CsvData {
	return core.ImportCsvData(a.ctx)
}

func (a *App) ExportJsonSchema(schema entity.JsonSchema) {
	core.WriteJsonSchema(a.ctx, schema)
}

func (a *App) WriteJsonFromCsvModelMap(modelMap entity.CsvModelMap) {
	core.WriteJsonFromCsvModelMap(a.ctx, modelMap)
}
