package healthCheck

// Response struct for healthCheck implementation
type Response struct {
	Timestamp string `json:"timestamp"`
	Message   string `json:"message"`
}
