package response

type TemplateResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
