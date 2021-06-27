package router

// Response is wrap response api
type Response struct {
	Status int         `json:"status,omitempty"`
	Data   interface{} `json:"data,omitempty"`
	Mesage string      `json:"mesage,omitempty"`
	Error  interface{} `json:"error,omitempty"`
}
