package requests

type CreateTodoInput struct {
  Title  string `json:"title" binding:"required"`
  Author string `json:"author" binding:"required"`
}

type UpdateTodoInput struct {
  Title  string `json:"title" binding:"required"`
  Author string `json:"author" binding:"required"`
}
