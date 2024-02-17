package core

import (
	"csvtoschema/backend/entity"
	"strconv"
)

func ConvertColumnValueType(propSchema entity.PropertySchema, rowSchema entity.CsvRowSchema, colNum int, pErrCh chan<- entity.CsvProcessingError) (any, bool) {
	ok := true
	switch propSchema.Type {
	case "string":
		// valid by default since CSV values are all strings
		return rowSchema.RowData[colNum], true
	case "number":
		v, err := strconv.ParseFloat(rowSchema.RowData[colNum], 64)
		if err != nil {
			ok = false
			pErr := CreateColValErr(rowSchema.RowNum, colNum, rowSchema.RowData[colNum], propSchema.Type)
			pErrCh <- pErr
		}
		return v, ok
	case "integer":
		v, err := strconv.Atoi(rowSchema.RowData[colNum])
		if err != nil {
			ok = false
			pErr := CreateColValErr(rowSchema.RowNum, colNum, rowSchema.RowData[colNum], propSchema.Type)
			pErrCh <- pErr
		}
		return v, ok
	case "boolean":
		v, err := strconv.ParseBool(rowSchema.RowData[colNum])
		if err != nil {
			ok = false
			pErr := CreateColValErr(rowSchema.RowNum, colNum, rowSchema.RowData[colNum], propSchema.Type)
			pErrCh <- pErr
		}
		return v, ok
	case "null":
		return nil, ok
	default:
		return rowSchema.RowData[colNum], false
	}
}

func ConvertToFloat(propSchema entity.PropertySchema, rowSchema entity.CsvRowSchema, colNum int, pErrCh chan<- entity.CsvProcessingError) (float64, bool) {
	ok := true
	v, err := strconv.ParseFloat(rowSchema.RowData[colNum], 64)
	if err != nil {
		ok = false
		pErr := CreateColValErr(rowSchema.RowNum, colNum, rowSchema.RowData[colNum], propSchema.Type)
		pErrCh <- pErr
	}
	return v, ok
}

func ConvertToInt(propSchema entity.PropertySchema, rowSchema entity.CsvRowSchema, colNum int, pErrCh chan<- entity.CsvProcessingError) (int, bool) {
	ok := true
	v, err := strconv.Atoi(rowSchema.RowData[colNum])
	if err != nil {
		ok = false
		pErr := CreateColValErr(rowSchema.RowNum, colNum, rowSchema.RowData[colNum], propSchema.Type)
		pErrCh <- pErr
	}
	return v, ok
}

func ConvertToBool(propSchema entity.PropertySchema, rowSchema entity.CsvRowSchema, colNum int, pErrCh chan<- entity.CsvProcessingError) (bool, bool) {
	ok := true
	v, err := strconv.ParseBool(rowSchema.RowData[colNum])
	if err != nil {
		ok = false
		pErr := CreateColValErr(rowSchema.RowNum, colNum, rowSchema.RowData[colNum], propSchema.Type)
		pErrCh <- pErr
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

func ConvertArrayItemType(propSchema entity.PropertySchema, rowSchema entity.CsvRowSchema, colNum int, pErrCh chan<- entity.CsvProcessingError) (any, bool) {
	ok := true
	if propSchema.Items != nil {
		// Wonky because Wails won't generate a correct TypeScript class when using a struct
		// for Items in PropertySchema and the convertValues method is broken in it
		arrayItemAttr := propSchema.Items.(map[string]interface{})
		itemTypeAttr := arrayItemAttr["type"]
		itemType := itemTypeAttr.(string)
		switch itemType {
		case "string":
			return rowSchema.RowData[colNum], true
		case "number":
			v, err := strconv.ParseFloat(rowSchema.RowData[colNum], 64)
			if err != nil {
				ok = false
				pErr := CreateArrItemErr(rowSchema.RowNum, colNum, rowSchema.RowData[colNum], itemType)
				pErrCh <- pErr
			}
			return v, ok
		case "integer":
			v, err := strconv.Atoi(rowSchema.RowData[colNum])
			if err != nil {
				ok = false
				pErr := CreateArrItemErr(rowSchema.RowNum, colNum, rowSchema.RowData[colNum], itemType)
				pErrCh <- pErr
			}
			return v, ok
		default:
			// if itemType is provided but doesn't match any case, type is invalid
			return rowSchema.RowData[colNum], false
		}
	}
	// if itemType is not provided, any type is valid
	return rowSchema.RowData[colNum], false
}
