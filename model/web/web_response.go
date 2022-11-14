package web

type WebResponse struct {
	Error   bool        `json:"error,omitempty"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
