package core

import (
	"StructifyCSV/backend/entity"
	"fmt"
)

// todo add keys (property names) to errors

func ValidateStringProperty(val string, propKey string, propSchema entity.PropertySchema, rowNum int, colNum int, rowErrCh chan<- entity.RowError) bool {
	isValid := true
	if propSchema.MinLength != nil {
		min := int(*propSchema.MinLength)
		if len(val) < min {
			isValid = false
			msg := fmt.Sprintf("length is less than %v", min)
			rowErr := CreateValidationError(val, propKey, rowNum, colNum, propSchema.Type, msg)
			rowErrCh <- rowErr
		}
	}
	if propSchema.MaxLength != nil {
		max := int(*propSchema.MaxLength)
		if len(val) > max {
			isValid = false
			msg := fmt.Sprintf("length is greater than %v", max)
			rowErr := CreateValidationError(val, propKey, rowNum, colNum, propSchema.Type, msg)
			rowErrCh <- rowErr
		}
	}
	return isValid
}

func ValidateNumberProperty(val float64, propKey string, propSchema entity.PropertySchema, rowNum int, colNum int, rowErrCh chan<- entity.RowError) bool {
	isValid := true
	if propSchema.NumMinimum != nil {
		min := float64(*propSchema.NumMinimum)
		if val < min {
			isValid = false
			msg := fmt.Sprintf("is less than %v.", min)
			rowErr := CreateValidationError(val, propKey, rowNum, colNum, propSchema.Type, msg)
			rowErrCh <- rowErr
		}
	}
	if propSchema.NumMaximum != nil {
		max := float64(*propSchema.NumMaximum)
		if val > max {
			isValid = false
			msg := fmt.Sprintf("is greater than %v.", max)
			rowErr := CreateValidationError(val, propKey, rowNum, colNum, propSchema.Type, msg)
			rowErrCh <- rowErr
		}
	}
	return isValid
}

func ValidateIntegerProperty(val int, propKey string, propSchema entity.PropertySchema, rowNum int, colNum int, rowErrCh chan<- entity.RowError) bool {
	isValid := true
	if propSchema.IntMinimum != nil {
		min := int(*propSchema.IntMinimum)
		if val < min {
			isValid = false
			msg := fmt.Sprintf("is less than %v.", min)
			rowErr := CreateValidationError(val, propKey, rowNum, colNum, propSchema.Type, msg)
			rowErrCh <- rowErr
		}
	}
	if propSchema.IntMaximum != nil {
		max := int(*propSchema.IntMaximum)
		if val > max {
			isValid = false
			msg := fmt.Sprintf("is greater than %v.", max)
			rowErr := CreateValidationError(val, propKey, rowNum, colNum, propSchema.Type, msg)
			rowErrCh <- rowErr
		}
	}
	return isValid
}

// needs more complex error for why an array length may not match due to value of converting
// items types
func ValidateArrayProperty(val []any, propKey string, propSchema entity.PropertySchema, rowNum int, rowErrCh chan<- entity.RowError) bool {
	isValid := true
	if propSchema.MinItems != nil {
		min := int(*propSchema.MinItems)
		if len(val) < min {
			isValid = false
			msg := fmt.Sprintf("length is less than %v.", min)
			rowErr := entity.RowError{Row: rowNum, PropertyKey: propKey, PropertyType: propSchema.Type, ErrorType: "property validation failure", ErrorMessage: msg}
			rowErrCh <- rowErr
		}
	}
	if propSchema.MaxItems != nil {
		max := int(*propSchema.MaxItems)
		if len(val) > max {
			isValid = false
			msg := fmt.Sprintf("length is greater than %v.", max)
			rowErr := entity.RowError{Row: rowNum, PropertyKey: propKey, PropertyType: propSchema.Type, ErrorType: "property validation failure", ErrorMessage: msg}
			rowErrCh <- rowErr
		}
	}
	return isValid
}

// needs more complex validation for required
func ValidateObjectProperty(val map[string]interface{}, propKey string, propSchema entity.PropertySchema, rowNum int, rowErrCh chan<- entity.RowError) bool {
	isValid := true
	if propSchema.MinProperties != nil {
		min := int(*propSchema.MinProperties)
		if len(val) < min {
			isValid = false
			msg := fmt.Sprintf("total properties is less than %v.", min)
			rowErr := entity.RowError{Row: rowNum, PropertyKey: propKey, PropertyType: propSchema.Type, ErrorType: "property validation failure", ErrorMessage: msg}
			rowErrCh <- rowErr
		}
	}
	if propSchema.MaxProperties != nil {
		max := int(*propSchema.MaxProperties)
		if len(val) > max {
			isValid = false
			msg := fmt.Sprintf("total properties is greater than %v.", max)
			rowErr := entity.RowError{Row: rowNum, PropertyKey: propKey, PropertyType: propSchema.Type, ErrorType: "property validation failure", ErrorMessage: msg}
			rowErrCh <- rowErr
		}
	}
	return isValid
}
