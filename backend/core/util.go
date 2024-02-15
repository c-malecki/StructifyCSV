package core

import (
	"csvtoschema/backend/entity"
	"errors"
	"fmt"
	"os"
)

func CreateStringWriter(filePath string) func(string, bool) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	return func(data string, close bool) {
		_, err := file.WriteString(data)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}

		if close {
			file.Close()
		}
	}
}

func CreatePErr(msg string, rowNum int) entity.CsvProcessingError {
	err := errors.New(msg)
	pErr := entity.CsvProcessingError{RowNum: rowNum, Error: err}
	fmt.Printf("Error: %s\n", err)
	return pErr
}
