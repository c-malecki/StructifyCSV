package core

import (
	"csvtoschema/backend/entity"
	"fmt"
)

func ValidateStringSchema(propSchema entity.PropertySchema, rowNum int, colNum int, str string, pErrCh chan<- entity.CsvProcessingError) bool {
	isValid := true
	if propSchema.MinLength != nil {
		min := int(*propSchema.MinLength)
		if len(str) != min {
			msg := fmt.Sprintf("row %v column %v: string value \"%v\"'s length is less than %v", rowNum, colNum, str, min)
			pErr := CreatePErr(msg, rowNum)
			pErrCh <- pErr
			isValid = false
		}
	}
	if propSchema.MaxLength != nil {
		max := int(*propSchema.MaxLength)
		if len(str) != max {
			msg := fmt.Sprintf("row %v column %v: string value \"%v\"'s length is greater than %v", rowNum, colNum, str, max)
			pErr := CreatePErr(msg, rowNum)
			pErrCh <- pErr
			isValid = false
		}
	}
	return isValid
}

func ValidateColumnValue(propSchema entity.PropertySchema, rowNum int, colNum int, value any, pErrCh chan<- entity.CsvProcessingError) bool {
	switch propSchema.Type {
	case "string":
		isValid := ValidateStringSchema(propSchema, rowNum, colNum, value.(string), pErrCh)
		return isValid
	// case "number":
	// handle error for improper value conversion result
	// 	v, err := strconv.ParseFloat(lineSchema.rowData[colNum], 64)
	// 	return v, err
	// case "integer":
	// handle error for improper value conversion result
	// v, err := strconv.Atoi(lineSchema.rowData[colNum])
	// return v, err
	// case "object":
	//  map ??
	// case "array":
	//  slice ??
	// case "boolean":
	// handle error for improper value conversion result
	// 	v, err := strconv.ParseBool(lineSchema.rowData[colNum])
	// 	return v, err
	// case "null":
	// 	return nil, nil
	default:
		msg := fmt.Sprintf("row %v column %v: column value \"%v\" doesn't match accepted data types", rowNum, colNum, value)
		pErr := CreatePErr(msg, rowNum)
		pErr.ColNum = colNum
		pErrCh <- pErr
		return false
	}
}
