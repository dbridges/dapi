package models

// Request represents a single named API request
type Request struct {
	Name    string
	URL     string
	Params  string
	Body    string
	Method  string
	Headers map[string]string
	Saved   bool `json:"-"`
}
