package dto

type GetTodosReq struct {
	Title       string `form:"title"`
	Description string `form:"description"`
	PageNo      int    `form:"pageNo" binding:"min=0"`
	PageSize    int    `form:"pageSize" binding:"min=0"`
}

type GetTodosRes struct {
	Todos []Todo `json:"todos"`
}

type GetTodoRes struct {
	Todo Todo `json:"todo"`
}

type Todo struct {
	ID          uint   `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}
