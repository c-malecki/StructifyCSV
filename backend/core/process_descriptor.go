package core

import (
	"csvtoschema/backend/entity"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

func ProcessCsvFile(csvLocation string, headerDescriptors []entity.HeaderDescriptor, writerCh chan<- map[string]string) {
	file, err := os.Open(csvLocation)
	if err != nil {
		print(err)
	}
	defer file.Close()

	var headers, line []string
	reader := csv.NewReader(file)
	headers, err = reader.Read()
	if err != nil {
		print(err)
	}
	var headerIndexes []int
	for i, h := range headers {
		for _, hd := range headerDescriptors {
			if h == hd.Header {
				headerIndexes = append(headerIndexes, i)
			}
		}
	}

	for {
		line, err = reader.Read()

		if err == io.EOF {
			close(writerCh)
			break
		} else if err != nil {
			print(err)
		}

		record, err := processCsvLine(headers, line, headerIndexes)

		if err != nil {
			// need to return array of lines which were not processed to frontend
			fmt.Printf("Line: %sError: %s\n", line, err)
			continue
		}
		writerCh <- record
	}
}

func WriteDescriptorJson(filePath string, csvLocation string, writerCh <-chan map[string]string, done chan<- bool) {
	jsonFunc := func(record map[string]string) string {
		jsonData, _ := json.MarshalIndent(record, "  ", "  ")
		return "  " + string(jsonData)
	}
	writeString := createStringWriter(filePath)
	writeString("[\n", false)

	first := true
	for {
		record, more := <-writerCh
		if more {
			if !first {
				writeString(",\n", false)
			} else {
				first = false
			}

			jsonData := jsonFunc(record)
			writeString(jsonData, false)
		} else {
			writeString("\n]", true)
			done <- true
			break
		}
	}
}

func createStringWriter(filePath string) func(string, bool) {
	file, err := os.Create(filePath)
	if err != nil {
		print(err)
	}

	return func(data string, close bool) {
		_, err := file.WriteString(data)
		if err != nil {
			print(err)
		}

		if close {
			file.Close()
		}
	}
}

// need to add nesting and data type conversions

func processCsvLine(headers []string, lineData []string, headerIndexes []int) (map[string]string, error) {
	if len(lineData) != len(headers) {
		return nil, errors.New("line doesn't match headers format. skipping")
	}

	recordMap := make(map[string]string)

	for _, index := range headerIndexes {
		recordMap[headers[index]] = lineData[index]
	}

	return recordMap, nil
}
