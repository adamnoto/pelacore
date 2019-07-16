package renderings

// Execution is a summary, or status, of an API execution
type Execution struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
