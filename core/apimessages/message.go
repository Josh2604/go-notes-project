package apimessages

type SuccessMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
