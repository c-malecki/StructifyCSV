package ui

import (
	"context"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func buildSaveDialogOpts(entityName string) runtime.SaveDialogOptions {
	fileName := strings.ReplaceAll(entityName, " ", "_")
	return runtime.SaveDialogOptions{
		DefaultDirectory: ".",
		DefaultFilename:  fileName + ".schema.json",
		Title:            "Save " + entityName + " Schema",
	}
}

func buildOpenDialogOptions(fileType string, entityType string) runtime.OpenDialogOptions {
	return runtime.OpenDialogOptions{
		Title: "Import " + entityType + "." + fileType,
		Filters: []runtime.FileFilter{
			{
				DisplayName: fileType + "(*." + fileType,
				Pattern:     "*." + fileType,
			},
		},
	}
}

func PrepareSaveFileDialog(c context.Context, entityName string) string {
	opts := buildSaveDialogOpts(entityName)
	filePath, err := runtime.SaveFileDialog(c, opts)
	if err != nil {
		print(err)
	}
	return filePath
}

func PrepareOpenFileDialog(c context.Context, fileType string, entityType string) string {
	opts := buildOpenDialogOptions(fileType, entityType)
	filePath, err := runtime.OpenFileDialog(c, opts)
	if err != nil {
		print(err)
	}
	return filePath
}
