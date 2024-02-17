package core

import (
	"csvtoschema/backend/entity"
	"fmt"
)

// todo add keys (property names) to errors

func ValidateStringProperty(propSchema entity.PropertySchema, rowNum int, colNum int, str string, pErrCh chan<- entity.CsvProcessingError) bool {
	isValid := true
	if propSchema.MinLength != nil {
		min := int(*propSchema.MinLength)
		if len(str) < min {
			isValid = false
			msg := fmt.Sprintf("row %v column %v: String value \"%v\"'s length is less than %v.", rowNum, colNum, str, min)
			pErr := CreatePErr(msg, rowNum)
			pErrCh <- pErr
		}
	}
	if propSchema.MaxLength != nil {
		max := int(*propSchema.MaxLength)
		if len(str) > max {
			isValid = false
			msg := fmt.Sprintf("row %v column %v: String value \"%v\"'s length is greater than %v.", rowNum, colNum, str, max)
			pErr := CreatePErr(msg, rowNum)
			pErrCh <- pErr
		}
	}
	return isValid
}

func ValidateNumberProperty(propSchema entity.PropertySchema, rowNum int, colNum int, num float64, pErrCh chan<- entity.CsvProcessingError) bool {
	isValid := true
	if propSchema.NumMinimum != nil {
		min := float64(*propSchema.NumMinimum)
		if num < min {
			isValid = false
			msg := fmt.Sprintf("row %v column %v: Number value \"%v\" is less than %v.", rowNum, colNum, num, min)
			pErr := CreatePErr(msg, rowNum)
			pErrCh <- pErr
		}
	}
	if propSchema.NumMaximum != nil {
		max := float64(*propSchema.NumMaximum)
		if num > max {
			isValid = false
			msg := fmt.Sprintf("row %v column %v: Number value \"%v\" is greater than %v.", rowNum, colNum, num, max)
			pErr := CreatePErr(msg, rowNum)
			pErrCh <- pErr
		}
	}
	return isValid
}

func ValidateIntegerProperty(propSchema entity.PropertySchema, rowNum int, colNum int, vInt int, pErrCh chan<- entity.CsvProcessingError) bool {
	isValid := true
	if propSchema.IntMinimum != nil {
		min := int(*propSchema.IntMinimum)
		if vInt < min {
			isValid = false
			msg := fmt.Sprintf("row %v column %v: Integer value \"%v\" is less than %v.", rowNum, colNum, vInt, min)
			pErr := CreatePErr(msg, rowNum)
			pErrCh <- pErr
		}
	}
	if propSchema.IntMaximum != nil {
		max := int(*propSchema.IntMaximum)
		if vInt > max {
			isValid = false
			msg := fmt.Sprintf("row %v column %v: Integer value \"%v\" is greater than %v.", rowNum, colNum, vInt, max)
			pErr := CreatePErr(msg, rowNum)
			pErrCh <- pErr
		}
	}
	return isValid
}

// needs more complex error for why an array length may not match due to value of converting
// items types
func ValidateArrayProperty(propSchema entity.PropertySchema, rowNum int, arr []any, pErrCh chan<- entity.CsvProcessingError) bool {
	isValid := true
	if propSchema.MinItems != nil {
		min := int(*propSchema.MinItems)
		if len(arr) < min {
			isValid = false
			msg := fmt.Sprintf("row %v: Array \"%v\"'s length is less than %v.", rowNum, arr, min)
			pErr := CreatePErr(msg, rowNum)
			pErrCh <- pErr
		}
	}
	if propSchema.MaxItems != nil {
		max := int(*propSchema.MaxItems)
		if len(arr) > max {
			isValid = false
			msg := fmt.Sprintf("row %v: Array \"%v\"'s length is greater than %v.", rowNum, arr, max)
			pErr := CreatePErr(msg, rowNum)
			pErrCh <- pErr
		}
	}
	return isValid
}

// needs more complex validation for required
func ValidateObjectProperty(propSchema entity.PropertySchema, key string, rowNum int, pErrCh chan<- entity.CsvProcessingError) bool {
	isValid := true
	if propSchema.MinProperties != nil {
		min := int(*propSchema.MinProperties)
		if len(propSchema.Properties) < min {
			isValid = false
			msg := fmt.Sprintf("row %v: Object \"%v\"'s total properties is less than %v.", rowNum, key, min)
			pErr := CreatePErr(msg, rowNum)
			pErrCh <- pErr
		}
	}
	if propSchema.MaxProperties != nil {
		max := int(*propSchema.MaxProperties)
		if len(propSchema.Properties) > max {
			isValid = false
			msg := fmt.Sprintf("row %v: Object \"%v\"'s total properties is greater than %v.", rowNum, key, max)
			pErr := CreatePErr(msg, rowNum)
			pErrCh <- pErr
		}
	}
	return isValid
}
