package controllers

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"

	"app/models"
	"app/requests"
)

type TodoController struct {
  Model *models.TodoModel
}

func NewTodoController(m *models.TodoModel) *TodoController {
  return &TodoController{Model: m}
}

func (mc *TodoController) GetTodos(c *gin.Context) {
  todos, err := mc.Model.GetAll()
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }

  c.JSON(http.StatusOK, gin.H{"data": todos})
}

func (mc *TodoController) GetTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

  todo, err := mc.Model.GetByID(uint(id))
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }

  c.JSON(http.StatusOK, gin.H{"data": todo})
}

func (mc *TodoController) CreateTodo(c *gin.Context) {
  var input requests.CreateTodoInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  todo, err := mc.Model.Create(input.Title, input.Author)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }

  c.JSON(http.StatusOK, gin.H{"data": todo})
}

func (mc *TodoController) UpdateTodo(c *gin.Context) {
  id, err := strconv.Atoi(c.Param("id"))
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
    return
  }

  var input requests.UpdateTodoInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  todo, err := mc.Model.Update(uint(id), input.Title, input.Author)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }

  c.JSON(http.StatusOK, gin.H{"data": todo})
}

func (mc *TodoController) DeleteTodo(c *gin.Context) {
  id, err := strconv.Atoi(c.Param("id"))
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
    return
  }

  if err := mc.Model.Delete(uint(id)); err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }

  c.JSON(http.StatusOK, gin.H{"data": true})
}
