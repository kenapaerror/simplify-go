package web

type TokenCreateRequest struct {
	UserId    string `validate:"required,min=1" json:"user_id"`
	Email     string `validate:"required,min=1,email" json:"email"`
	FirstName string `validate:"required,min=1" json:"first_name"`
	LastName  string `validate:"required,min=1" json:"last_name"`
	Role      string `validate:"required,min=1" json:"role"`
}
