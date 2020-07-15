package http

//Response ...
type Response struct {
	Data    interface{} `json:"data,omitempty"`
	Message interface{} `json:"message,omitempty"`
}
