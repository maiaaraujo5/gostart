package response

type Response struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Status      string `json:"status,omitempty"`
	Error       error  `json:"error,omitempty"`
}
