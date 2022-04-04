package helpers

type Response struct {
	StatusCode int         `json:"statusCode,omitempty"`
	Message    interface{} `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Halaman    int         `json:"halaman,omitempty"`
}
