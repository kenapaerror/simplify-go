package web

type ExampleCreateRequest struct {
	Name  string `validate:"required,min=1,max=50" json:"name"`
	Email string `validate:"required,email" json:"email"`
}
