package entity

type DataType struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Field struct {
	Name     string   `json:"name"`
	Type     string   `json:"type"`
	DataType DataType `json:"dataType"`
}

type Schema struct {
	Name   string  `json:"name"`
	Type   string  `json:"type"`
	Fields []Field `json:"fields"`
}

type Model struct {
	Name          string   `json:"name"`
	Type          string   `json:"type"`
	Schemas       []Schema `json:"schemas"`
	BaseSchemaIdx int      `json:"baseSchemaIdx"`
}

type CsvData struct {
	FileName string   `json:"fileName"`
	Location string   `json:"location"`
	Headers  []string `json:"headers"`
}

type HeaderDescriptor struct {
	Header    string `json:"header"`
	SchemaIdx int    `json:"schemaIdx"`
	FieldIdx  int    `json:"fieldIdx"`
}
