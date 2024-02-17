package core

import (
	"csvtoschema/backend/entity"
	"fmt"
)

func CreateHeaderMismatchErr(rowNum int, rowLen int, headerLen int) entity.CsvProcessingError {
	msg := fmt.Sprintf("row %v: Skipped. Row column count (%v) doesn't match header count (%v).", rowNum, rowLen, headerLen)
	pErr := entity.CsvProcessingError{RowNum: rowNum, Error: msg}
	return pErr
}

func CreateColValErr(rowNum int, colNum int, val string, cType string) entity.CsvProcessingError {
	msg := fmt.Sprintf("row %v column %v: \"%v\" could not be converted to type %v.", rowNum, colNum, val, cType)
	pErr := entity.CsvProcessingError{RowNum: rowNum, ColNum: colNum, Error: msg}
	return pErr
}

func CreateArrItemErr(rowNum int, colNum int, val string, cType string) entity.CsvProcessingError {
	msg := fmt.Sprintf("row %v column %v: \"%v\" could not be converted to type %v.", rowNum, colNum, val, cType)
	pErr := entity.CsvProcessingError{RowNum: rowNum, ColNum: colNum, Error: msg}
	return pErr
}
