package entity

type JsonSchema struct {
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Properties  map[string]interface{} `json:"properties"`
}

const (
	Indent string = "    "
)

type DataType string

const (
	String  DataType = "string"
	Number  DataType = "number"
	Integer DataType = "integer"
	Object  DataType = "object"
	Array   DataType = "array"
	Boolean DataType = "boolean"
	Null    DataType = "null"
)

type CsvSchemaProperty struct {
	HeaderIndexes      []int  `json:"headerIndexes"`
	SchemaPropertyType string `json:"schemaPropertyType"`
}

type CsvModelMap map[string]interface{}

type CsvData struct {
	FileName string   `json:"fileName"`
	Location string   `json:"location"`
	Headers  []string `json:"headers"`
}
