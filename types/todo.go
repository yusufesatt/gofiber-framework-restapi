package types

type TodoCreateDTO struct {
	Title string `json:"title" validate:"required"`
}

type TodoUpdateDTO struct {
	Title     *string `json:"title" validate:"required"`
	Completed *bool   `json:"completed" validate:"required"`
}

type TodoResponse struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
