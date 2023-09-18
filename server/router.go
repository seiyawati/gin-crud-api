package server

import (
	"github.com/gin-gonic/gin"

	"app/middleware"
	"app/controllers"
)

func InitServer() *gin.Engine {
	router := gin.Default()
	corsMiddleware := middleware.CORSMiddleware()
	todoController := controllers.NewTodoController()

	router.Use(corsMiddleware)
	router.GET("/todos", todoController.GetTodos)
	router.GET("/todos/:id", todoController.GetTodo)
	router.POST("/todos", todoController.CreateTodo)
	router.PATCH("/todos/:id", todoController.UpdateTodo)
	router.DELETE("/todos/:id", todoController.DeleteTodo)

	router.Run(":8080")

	return router
}
