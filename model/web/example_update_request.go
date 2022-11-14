package web

type ExampleUpdateRequest struct {
	Id    string `validate:"required" json:"id"`
	Name  string `validate:"required,min=1,max=50" json:"name"`
	Email string `validate:"required,email" json:"email"`
}
