package requests

type UpdateNoteRequest struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Deleted     *bool   `json:"deleted,omitempty" binding:"omitempty"`
}
