package requests

type RegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email,unique_email"`
	Password string `json:"password" validate:"required,min=6"`
}
