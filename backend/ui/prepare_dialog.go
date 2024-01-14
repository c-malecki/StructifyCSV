package ui

import (
	"context"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// func buildDialogOptions(fileType string,action string,  entityType string) runtime.OpenDialogOptions {
// 	if action == "Import" {
// 		return runtime.OpenDialogOptions{
// 			Title: action + entityType + "." + fileType,
// 			Filters: [] runtime.FileFilter{
// 				{
// 					DisplayName: fileType + "(*." + fileType,
// 					Pattern: "*." + fileType,
// 				},
// 			},
// 		}
// 	} else {
// 		return runtime.OpenDialogOptions{
// 			DefaultDirectory: ".",
// 			DefaultFilename:
// 		}
// 	}
// }

func PrepareOpenFileDialog(c context.Context, entityType string) string {
	filePath, err := runtime.OpenFileDialog(c, runtime.OpenDialogOptions{
		Title: "Import" + entityType + ".json",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "JSON (*.json)",
				Pattern:     "*.json",
			},
		},
	})
	if err != nil {
		print(err)
	}
	return filePath
}

func PrepareSaveFileDialog(c context.Context, entityName string, entityType string) string {
	fileName := strings.ReplaceAll(entityName, " ", "_")
	file, err := runtime.SaveFileDialog(c, runtime.SaveDialogOptions{
		DefaultDirectory: ".",
		DefaultFilename:  fileName + "_" + entityType + ".json",
		Title:            "Save" + " " + entityName + " " + entityType,
	})
	if err != nil {
		print(err)
	}
	return file
}

func PrepareOpenDirectoryDialog(c context.Context) string {
	dirPath, err := runtime.OpenDirectoryDialog(c, runtime.OpenDialogOptions{
		Title: "Select Directory for Output",
	})
	if err != nil {
		print(err)
	}
	return dirPath
}
