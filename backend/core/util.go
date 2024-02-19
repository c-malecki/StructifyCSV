package core

import (
	"errors"
	"fmt"
	"math"
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

// indexToColumn takes in an index value & converts it to A1 Notation
// Index 1 is Column A
// E.g. 3 == C, 29 == AC, 731 == ABC
func IndexToColumn(index int) (string, error) {

	// Validate index size
	maxIndex := 18278
	if index > maxIndex {
		err := errors.New("index cannot be greater than 18278 (column ZZZ)")
		return "", err
	}

	// Get column from index
	l := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	if index > 26 {
		letterA, _ := IndexToColumn(int(math.Floor(float64(index-1) / 26)))
		letterB, _ := IndexToColumn(index % 26)
		return letterA + letterB, nil
	} else {
		if index == 0 {
			index = 26
		}
		return string(l[index-1]), nil
	}

}
