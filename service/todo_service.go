package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todo-app/dto"
	"todo-app/model"
	repo "todo-app/repository"
)

type TodoService interface {
	GetTodos(c *gin.Context)
	GetTodo(c *gin.Context)
	CreateTodo(c *gin.Context)
	UpdateTodo(c *gin.Context)
	DeleteTodo(c *gin.Context)
}

type TodoServiceHandler struct {
	TodoRepo repo.TodoRepository
}

func NewTodoService(todoRepo repo.TodoRepository) *TodoServiceHandler {
	return &TodoServiceHandler{TodoRepo: todoRepo}
}

func (h *TodoServiceHandler) GetTodos(c *gin.Context) {
	var getTodoReq dto.GetTodosReq
	var getTodoRes dto.GetTodosRes
	if err := c.Bind(&getTodoReq); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	todos, err := h.TodoRepo.FindAllTodos(getTodoReq.PageNo*getTodoReq.PageSize, getTodoReq.PageSize)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	todosJson, err := json.Marshal(todos)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := json.Unmarshal(todosJson, &getTodoRes.Todos); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	res := dto.BaseRes{Data: getTodoRes}

	c.JSON(http.StatusOK, res)
}

func (h *TodoServiceHandler) GetTodo(c *gin.Context) {
	var todoRes dto.Todo
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	todo, err := h.TodoRepo.FindTodoById(id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	todosJson, err := json.Marshal(todo)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := json.Unmarshal(todosJson, &todoRes); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	if todoRes.ID == 0 {
		return
	}

	var getTodoRes dto.GetTodoRes
	getTodoRes.Todo = todoRes
	res := dto.BaseRes{Data: getTodoRes}
	c.JSON(http.StatusOK, res)
}

func (h *TodoServiceHandler) CreateTodo(c *gin.Context) {
	var todo dto.Todo

	if err := c.BindJSON(&todo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.TodoRepo.CreateTodo(todo)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
		return
	}

	newTodo, err := h.TodoRepo.FindTodoById(id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
		return
	}

	todosJson, err := json.Marshal(newTodo)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := json.Unmarshal(todosJson, &todo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	var getTodoRes dto.GetTodoRes
	getTodoRes.Todo = todo
	res := dto.BaseRes{Data: getTodoRes}

	c.JSON(http.StatusCreated, res)
}

func (h TodoServiceHandler) UpdateTodo(c *gin.Context) {
	tx, err := h.TodoRepo.GetConnection().BeginTx(c, nil)

	defer tx.Rollback()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	var todoDto dto.Todo
	var todoModel model.Todo
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		return
	}

	if err := c.BindJSON(&todoDto); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	todosJson, err := json.Marshal(todoDto)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := json.Unmarshal(todosJson, &todoModel); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	todoModel.ID = id
	todo, err := h.TodoRepo.FindTodoById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, todo)
		return
	}

	rows, err := h.TodoRepo.UpdateTodo(tx, todoModel)

	if err != nil || rows < 1 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	todosJson, err = json.Marshal(todo)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := json.Unmarshal(todosJson, &todoDto); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	var getTodoRes dto.GetTodoRes
	getTodoRes.Todo = todoDto
	res := dto.BaseRes{Data: getTodoRes}

	if err := tx.Commit(); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *TodoServiceHandler) DeleteTodo(c *gin.Context) {
	res := dto.BaseRes{}
	tx, err := h.TodoRepo.GetConnection().BeginTx(c, nil)
	defer tx.Rollback()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	idStr := c.Params.ByName("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	rows, err := h.TodoRepo.DeleteTodo(tx, id)

	if err != nil || rows < 1 {
		c.AbortWithStatusJSON(http.StatusNotFound, res)
		return
	}

	if err := tx.Commit(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	c.JSON(http.StatusOK, res)
}
