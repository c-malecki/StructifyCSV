package core

import "os"

func CreateStringWriter(filePath string) func(string, bool) {
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
