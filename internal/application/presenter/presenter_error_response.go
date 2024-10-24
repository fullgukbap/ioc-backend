package presenter

type ErrorResponse struct {
	Type    string `json:"type"`
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Detail  string `json:"detail,omitempty"`
}
