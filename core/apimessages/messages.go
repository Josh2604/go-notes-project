package apimessages

type Message string

var (
	ErrorBindingNote  Message = "Error binding note creation"
	ErrorCreatingNote Message = "Error creaeting note"
	ErrorDeletingNote Message = "Error trying to delete the message"
)

func (m Message) GetMessage() string {
	return string(m)
}
