package requests

type UpdateUserRequest struct {
	ID    uint   `json:"-"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email,unique_email=ID"`
}
