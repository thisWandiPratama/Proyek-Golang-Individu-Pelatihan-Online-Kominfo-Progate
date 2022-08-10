package todos

type CreateTodosInput struct {
	Task     string `json:"task" binding:"required"`
	Asignee  string `json:"asignee" binding:"required"`
	Deadline string `json:"deadline" binding:"required"`
	Done     string `json:"done" binding:"required"`
	Error    error
}

type UpdateTodosInput struct {
	ID       int    `json:"id" binding:"required"`
	Task     string `json:"task" binding:"required"`
	Asignee  string `json:"asignee" binding:"required"`
	Deadline string `json:"deadline" binding:"required"`
	Done     string `json:"done" binding:"required"`
	Error    error
}

type DeleteTodosInput struct {
	ID int `json:"id_todo" binding:"required"`
}
