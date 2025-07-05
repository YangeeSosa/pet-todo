package main

import (
	"log"

	"github.com/YangeeSosa/pet-todo-db/internal"
	"github.com/gin-gonic/gin"
)

func main() {
	internal.InitDB()
	r := gin.Default()
	r.POST("/tasks", internal.CreateTaskHandler)
	r.GET("/tasks", internal.GetTasksHandler)
	r.DELETE("/tasks", internal.DeleteTaskHandler)
	r.PUT("/tasks", internal.MarkTaskDoneHandler)

	log.Println("DB-сервис (Gin) запущен на :8081")
	if err := r.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
