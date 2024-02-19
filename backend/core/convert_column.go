package core

import (
	"StructifyCSV/backend/entity"
	"strconv"
)

func ConvertToFloat(val string, propKey string, propSchema entity.PropertySchema, rowNum int, colNum int, rowErrCh chan<- entity.RowError) (float64, bool) {
	ok := true
	v, err := strconv.ParseFloat(val, 64)
	if err != nil {
		ok = false
		rowErr := CreateConversionError(val, propKey, rowNum, colNum, propSchema.Type)
		rowErrCh <- rowErr
	}
	return v, ok
}

func ConvertToInt(val string, propKey string, propSchema entity.PropertySchema, rowNum int, colNum int, rowErrCh chan<- entity.RowError) (int, bool) {
	ok := true
	v, err := strconv.Atoi(val)
	if err != nil {
		ok = false
		rowErr := CreateConversionError(val, propKey, rowNum, colNum, propSchema.Type)
		rowErrCh <- rowErr
	}
	return v, ok
}

func ConvertToBool(val string, propKey string, propSchema entity.PropertySchema, rowNum int, colNum int, rowErrCh chan<- entity.RowError) (bool, bool) {
	ok := true
	v, err := strconv.ParseBool(val)
	if err != nil {
		ok = false
		rowErr := CreateConversionError(val, propKey, rowNum, colNum, propSchema.Type)
		rowErrCh <- rowErr
	}
	return v, ok
}

func ConvertIndexArrToInts(csvHeaderIndex interface{}) []int {
	indexes := csvHeaderIndex.([]interface{})
	colNums := make([]int, 0, len(indexes))
	for _, v := range indexes {
		i64 := v.(float64)
		index := int(i64)
		colNums = append(colNums, index)
	}
	return colNums
}

func ConvertArrayItemType(val string, propKey string, propSchema entity.PropertySchema, rowNum int, colNum int, rowErrCh chan<- entity.RowError) (any, bool) {
	ok := true
	if propSchema.Items != nil {
		// Wonky because Wails won't generate a correct TypeScript class when using a struct
		// for Items in PropertySchema and the convertValues method is broken in it
		arrayItemAttr := propSchema.Items.(map[string]interface{})
		itemTypeAttr := arrayItemAttr["type"]
		itemType := itemTypeAttr.(string)
		switch itemType {
		case "string":
			return val, true
		case "number":
			v, err := strconv.ParseFloat(val, 64)
			if err != nil {
				ok = false
				rowErr := CreateConversionError(val, propKey, rowNum, colNum, itemType)
				rowErrCh <- rowErr
			}
			return v, ok
		case "integer":
			v, err := strconv.Atoi(val)
			if err != nil {
				ok = false
				rowErr := CreateConversionError(val, propKey, rowNum, colNum, itemType)
				rowErrCh <- rowErr
			}
			return v, ok
		default:
			// if itemType is provided but doesn't match any case, type is invalid
			return val, false
		}
	}
	// if itemType is not provided, any type is valid
	return val, false
}
