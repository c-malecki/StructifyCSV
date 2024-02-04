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

type SchemaProperty struct {
	Key   string `json:"key"`
	Path  string `json:"path"`
	Value string `json:"value"`
}

type HeaderDescriptor struct {
	IsSelected     bool            `json:"isSelected"`
	HeaderText     string          `json:"headerText"`
	HeaderIndex    int             `json:"headerIndex"`
	SchemaProperty *SchemaProperty `json:"schemaProperty"`
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
