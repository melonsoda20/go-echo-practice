package requests

type CreateToDo struct {
	Name string `json:"name" validate:"required"`
}
