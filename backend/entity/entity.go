package entity

type DataType struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type JsonSchema struct {
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Properties  map[string]interface{} `json:"properties"`
}

// type JsonSchema struct {
// 	Title       string `json:"title"`
// 	Description string `json:"description"`
// 	Properties  string `json:"properties"`
// }

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
