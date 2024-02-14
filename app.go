package main

import (
	"context"
	"csvtoschema/backend/core"
	"csvtoschema/backend/entity"
	"csvtoschema/backend/ui"
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

func (a *App) ExportJsonSchema(schema entity.JsonSchema) {
	core.WriteJsonSchema(a.ctx, schema)
}

func (a *App) ProcessCsvWithSchema(schema entity.JsonSchema) {
	core.ProcessCsvWithSchema(a.ctx, schema)
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
