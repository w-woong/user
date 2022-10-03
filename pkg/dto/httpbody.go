package dto

type HttpBody struct {
	Status    int         `json:"status,omitempty"`
	Message   string      `json:"message,omitempty"`
	Count     int         `json:"count,omitempty"`
	Document  interface{} `json:"document,omitempty"`
	Documents interface{} `json:"documents,omitempty"`
}
