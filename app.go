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

func (a *App) ImportSchema() entity.Schema {
	return core.ImportSchemaJson(a.ctx)
}

func (a *App) ImportModel() entity.Model {
	return core.ImportModelJson(a.ctx)
}

func (a *App) ImportCsvData() entity.CsvData {
	return core.ImportCsvFileData(a.ctx)
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) ExportSchema(schema entity.Schema) {
	core.ExportSchemaToJson(a.ctx, schema)
}

func (a *App) ExportModel(model entity.Model) {
	core.ExportModelToJson(a.ctx, model)
}

func (a *App) ProcessCsvDescriptor(model entity.Model, hd []entity.HeaderDescriptor, csvPath string) {
	core.ExportCsvDescriptor(a.ctx, model, hd, csvPath)
}
