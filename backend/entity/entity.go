package entity

type Properties map[string]*PropertySchema

type JsonSchema struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	// will always default to an object as top level schema for now
	Type string `json:"type"`
	// MinProperties *int       `json:"minProperties"`
	// MaxProperties *int       `json:"maxProperties"`
	Required   []string   `json:"required"`
	Properties Properties `json:"properties"`
}

type PropertySchema struct {
	Type  string            `json:"type"`
	OneOf []*PropertySchema `json:"oneOf"`
	// object property validators
	MinProperties *int       `json:"minProperties"`
	MaxProperties *int       `json:"maxProperties"`
	Required      []string   `json:"required"`
	Properties    Properties `json:"properties"`
	// array property validators
	MinItems *int        `json:"minItems"`
	MaxItems *int        `json:"maxItems"`
	Items    interface{} `json:"items"`
	// string property validators
	MinLength *int `json:"minLength"`
	MaxLength *int `json:"maxLength"`
	// number property validators
	NumMinimum *float64 `json:"numMinimum"`
	NumMaximum *float64 `json:"numMaximum"`
	// integer property validators
	IntMinimum *int `json:"intMinimum"`
	IntMaximum *int `json:"intMaximum"`
	// csv header index to map to
	CsvHeaderIndex interface{} `json:"csvHeaderIndex"`
}

const (
	Indent string = "    "
)

type CsvHeader struct {
	Column string `json:"column"`
	Header string `json:"header"`
}

type CsvFileData struct {
	FileName      string      `json:"fileName"`
	Location      string      `json:"location"`
	Headers       []CsvHeader `json:"headers"`
	ReferenceRows [][]string  `json:"referenceRows"`
}

type RowError struct {
	Row          int    `json:"row"`
	Column       string `json:"column"`
	PropertyKey  string `json:"propertyKey"`
	PropertyType string `json:"propertyType"`
	Value        any    `json:"value"`
	ErrorType    string `json:"errorType"`
	ErrorMessage string `json:"errorMessage"`
}

type CsvProcessingReport struct {
	Successes []int `json:"successes"`
	// Failures []int        `json:"failures"`
	RowErrors []RowError `json:"rowErrors"`
}
