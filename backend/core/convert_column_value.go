package core

import (
	"csvtoschema/backend/entity"
	"strconv"
)

func ConvertColumnValueType(propSchema entity.PropertySchema, rowSchema entity.CsvRowSchema, colNum int, pErrCh chan<- entity.CsvProcessingError) (any, bool) {
	isValid := true
	switch propSchema.Type {
	case "string":
		return rowSchema.RowData[colNum], true
	case "number":
		v, err := strconv.ParseFloat(rowSchema.RowData[colNum], 64)
		if err != nil {
			isValid = false
			pErr := CreateColValErr(rowSchema.RowNum, colNum, rowSchema.RowData[colNum], propSchema.Type)
			pErrCh <- pErr
		}
		return v, isValid
	case "integer":
		v, err := strconv.Atoi(rowSchema.RowData[colNum])
		if err != nil {
			isValid = false
			pErr := CreateColValErr(rowSchema.RowNum, colNum, rowSchema.RowData[colNum], propSchema.Type)
			pErrCh <- pErr
		}
		return v, isValid
	case "boolean":
		v, err := strconv.ParseBool(rowSchema.RowData[colNum])
		if err != nil {
			isValid = false
			pErr := CreateColValErr(rowSchema.RowNum, colNum, rowSchema.RowData[colNum], propSchema.Type)
			pErrCh <- pErr
		}
		return v, isValid
	case "null":
		return nil, true
	default:
		return rowSchema.RowData[colNum], false
	}
}

func ConvertArrayItemType(propSchema entity.PropertySchema, rowSchema entity.CsvRowSchema, colNum int, pErrCh chan<- entity.CsvProcessingError) (any, bool) {
	isValid := true
	if propSchema.Items != nil {
		// wonky since the model.ts for frontend won't generate a struct for Items in JSON schema
		arrayItemAttr := propSchema.Items.(map[string]interface{})
		itemTypeAttr := arrayItemAttr["type"]
		itemType := itemTypeAttr.(string)
		switch itemType {
		case "string":
			return rowSchema.RowData[colNum], true
		case "number":
			v, err := strconv.ParseFloat(rowSchema.RowData[colNum], 64)
			if err != nil {
				isValid = false
				pErr := CreateArrItemErr(rowSchema.RowNum, colNum, rowSchema.RowData[colNum], itemType)
				pErrCh <- pErr
			}
			return v, isValid
		case "integer":
			v, err := strconv.Atoi(rowSchema.RowData[colNum])
			if err != nil {
				isValid = false
				pErr := CreateArrItemErr(rowSchema.RowNum, colNum, rowSchema.RowData[colNum], itemType)
				pErrCh <- pErr
			}
			return v, isValid
		default:
			// if itemType is provided but doesn't match any case, type is invalid
			return rowSchema.RowData[colNum], false
		}
	}
	// if itemType is not provided, any type is valid
	return rowSchema.RowData[colNum], true
}
