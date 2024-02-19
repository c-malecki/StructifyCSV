package core

import (
	"StructifyCSV/backend/entity"
	"fmt"
)

func CreateHeaderMismatchErr(rowNum int, rowLen int, headerLen int) entity.RowError {
	msg := fmt.Sprintf("row %v: Skipped. Row column count (%v) doesn't match header count (%v)", rowNum, rowLen, headerLen)
	rowErr := entity.RowError{Row: rowNum, ErrorMessage: msg}
	return rowErr
}

func CreateConversionError(val string, propKey string, rowNum int, colNum int, pType string) entity.RowError {
	msg := fmt.Sprintf("could not be converted to type %v", pType)
	col, _ := IndexToColumn(colNum + 1)
	// if err != nil {
	// 	// do what if there are too many cols?
	// }
	rowErr := entity.RowError{Row: rowNum, Column: col, PropertyKey: propKey, PropertyType: pType, Value: val, ErrorType: "data type conversion", ErrorMessage: msg}
	return rowErr
}

func CreateValidationError(val any, propKey string, rowNum int, colNum int, pType string, msg string) entity.RowError {
	col, _ := IndexToColumn(colNum + 1)
	// if err != nil {
	// 	// do what if there are too many cols?
	// }
	rowErr := entity.RowError{Row: rowNum, Column: col, PropertyKey: propKey, PropertyType: pType, Value: val, ErrorType: "property validation failure", ErrorMessage: msg}
	return rowErr
}
