package response

type CustomMessage struct {
	Code          int    `json:"code"`
	Message       string `json:"message"`
	CustomMessage string `json:"custom_message"`
}
