package renderings

// Row is a key-value data structure for representing a row in a table
type Row map[string]interface{}

// DatabaseColumns represents a column of a given query result
type DatabaseColumns struct {
	Name string `json:"name"`
	// For VARCHAR and TEXT type, it will be the length of the text.
	// For NUMERIC, it will be in the form of `precision, scale`
	Length   *string `json:"length"`
	Nullable *bool   `json:"nullable"`
	Type     string  `json:"type"`
}

// DatabaseQuery is a response for database query API
type DatabaseQuery struct {
	Execution Execution         `json:"execution"`
	Rows      []Row             `json:"rows"`
	Columns   []DatabaseColumns `json:"columns"`
}
