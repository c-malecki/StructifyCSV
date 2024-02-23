package main

import (
	"StructifyCSV/backend/core"
	"StructifyCSV/backend/entity"
	"StructifyCSV/backend/ui"
	"context"
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

func (a *App) ImportJsonSchema() core.ImportSchemaRes {
	return core.ImportJsonSchema(a.ctx)
}

func (a *App) ImportCsvFileData() entity.CsvFileData {
	return core.ImportCsvFileData(a.ctx)
}

func (a *App) ExportJsonSchema(jsonSchema entity.JsonSchema) {
	core.ExportJsonSchema(a.ctx, jsonSchema)
}

func (a *App) ProcessCsvWithSchema(csvFile entity.CsvFileData, jsonSchema entity.JsonSchema) entity.CsvProcessingReport {
	report := core.ProcessCsv(a.ctx, csvFile, jsonSchema)
	return report
}

func (a *App) MinimizeWindow() {
	ui.Minimize(a.ctx)
}

func (a *App) MaximizeWindow() {
	ui.Maximize(a.ctx)
}

func (a *App) UnmaximizeWindow() {
	ui.Unmaximize(a.ctx)
}

func (a *App) ExitProgram() {
	ui.Exit(a.ctx)
}
