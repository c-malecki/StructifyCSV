package entity

type Properties map[string]*Schema

type JsonSchema struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Properties  Properties `json:"properties"`
}

type Schema struct {
	Type  string    `json:"type"`
	OneOf []*Schema `json:"oneOf"`
	// object
	MinProperties *int       `json:"minProperties"`
	MaxProperties *int       `json:"maxProperties"`
	Required      []string   `json:"required"`
	Properties    Properties `json:"properties"`
	// array
	MinItems *int         `json:"minItems"`
	MaxItems *int         `json:"maxItems"`
	Items    *interface{} `json:"items"`
	// string
	MinLength *int `json:"minLength"`
	MaxLength *int `json:"maxLength"`
	// number
	NumMinimum *float64 `json:"numMinimum"`
	NumMaximum *float64 `json:"numMaximum"`
	// integer
	IntMinimum *int `json:"intMinimum"`
	IntMaximum *int `json:"intMaximum"`
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

type CsvFileData struct {
	FileName string   `json:"fileName"`
	Location string   `json:"location"`
	Headers  []string `json:"headers"`
}
