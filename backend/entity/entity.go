package entity

type JsonSchema struct {
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Properties  map[string]interface{} `json:"properties"`
}

type DataType string

const (
	Indent string = "    "
)

const (
	String  DataType = "string"
	Number  DataType = "number"
	Integer DataType = "integer"
	Object  DataType = "object"
	Array   DataType = "array"
	Boolean DataType = "boolean"
	Null    DataType = "null"
)

type SchemaProperty struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Path  string `json:"path"`
}

type HeaderDescriptor struct {
	IsSelected     bool            `json:"isSelected"`
	HeaderText     string          `json:"headerText"`
	HeaderIndex    int             `json:"headerIndex"`
	SchemaProperty *SchemaProperty `json:"schemaProperty"`
}

type CsvModelNodeValue struct {
	DataType   string  `json:"dataType"`
	Header     *string `json:"header"`
	HeaderIdx  *int    `json:"headerIdx"`
	SchemaPath string  `json:"schemaPath"`
}

type CsvModel struct {
	HeaderDescriptors []HeaderDescriptor     `json:"headerDescriptors"`
	UsedHeaderIndexes []int                  `json:"usedHeaderIndexes"`
	Map               map[string]interface{} `json:"map"`
}

type CsvData struct {
	FileName string   `json:"fileName"`
	Location string   `json:"location"`
	Headers  []string `json:"headers"`
}
