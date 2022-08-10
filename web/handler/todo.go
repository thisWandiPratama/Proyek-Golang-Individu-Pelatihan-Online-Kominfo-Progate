package handler

import (
	"net/http"
	"strconv"
	"todos/todos"

	"github.com/gin-gonic/gin"
)

type todoHandler struct {
	todoService todos.Service
}

func NewTodoHandler(todoService todos.Service) *todoHandler {
	return &todoHandler{todoService}
}

func (h *todoHandler) Index(c *gin.Context) {
	todos, err := h.todoService.FindAll()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.HTML(http.StatusOK, "todo_index.html", gin.H{"todos": todos})
}

func (h *todoHandler) New(c *gin.Context) {
	c.HTML(http.StatusOK, "todo_new.html", nil)
}

func (h *todoHandler) Create(c *gin.Context) {

	createTodo := todos.CreateTodosInput{}
	createTodo.Task = c.PostForm("task")
	createTodo.Asignee = c.PostForm("asignee")
	createTodo.Deadline = c.PostForm("deadline")
	_, err := h.todoService.Save(createTodo)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/")
}

func (h *todoHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	updateTodo := todos.UpdateTodosInput{}
	updateTodo.ID = id
	updateTodo.Task = c.PostForm("task")
	updateTodo.Asignee = c.PostForm("asignee")
	updateTodo.Deadline = c.PostForm("deadline")

	_, err := h.todoService.Update(updateTodo)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/")
}

func (h *todoHandler) Edit(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	todo, err := h.todoService.FindByID(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	input := todos.UpdateTodosInput{}
	input.ID = todo.ID
	input.Task = todo.Task
	input.Asignee = todo.Asignee
	input.Deadline = todo.Deadline

	c.HTML(http.StatusOK, "todo_edit.html", input)
}

func (h *todoHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	_, err := h.todoService.Delete(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	c.Redirect(http.StatusFound, "/")
}

func (h *todoHandler) Done(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	todo, err := h.todoService.FindByID(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	input := todos.UpdateTodosInput{}
	input.ID = todo.ID
	input.Task = todo.Task
	input.Asignee = todo.Asignee
	input.Deadline = todo.Deadline
	if todo.Done == "0" {
		input.Done = "1"
	} else {
		input.Done = "0"
	}

	_, err = h.todoService.Done(input)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/")
}
