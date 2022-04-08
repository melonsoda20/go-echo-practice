package requests

type CreateToDo struct {
	Name       string `json:"name" validate:"required"`
	Age        int    `json:"age"`
	Created_by string `json:"createdBy" validate:"required"`
}
