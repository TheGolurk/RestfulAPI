package response

// Model es una estructura de respuesta
type Model struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
