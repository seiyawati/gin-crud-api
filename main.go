package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"app/models"
	"app/controllers"
)

func main() {
	dsn := "user=postgres password=postgres dbname=postgres host=db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

	db.AutoMigrate(&models.Todo{})

	todoModel := models.NewTodoModel(db)
	todoController := controllers.NewTodoController(todoModel)

	router := gin.Default()
	router.GET("/todos", todoController.GetTodos)
	router.GET("/todos/:id", todoController.GetTodo)
	router.POST("/todos", todoController.CreateTodo)
	router.PATCH("/todos/:id", todoController.UpdateTodo)
	router.DELETE("/todos/:id", todoController.DeleteTodo)

	router.Run(":8080")
}
