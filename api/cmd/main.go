package main

import (
	"log"
	"os"

	"github.com/YangeeSosa/pet-todo-api/internal"
	"github.com/gin-gonic/gin"
)

func main() {
	logFile, err := os.OpenFile("api.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Не удалось открыть файл для логов: %v", err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)

	internal.InitKafkaProducer()

	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.String(200, "OK")
	})
	r.POST("/tasks", internal.CreateTask)
	r.GET("/tasks", internal.GetTasks)
	r.DELETE("/tasks", internal.DeleteTask)
	r.PUT("/tasks", internal.MarkTaskDone)

	log.Println("API-сервис (Gin) запущен на :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
