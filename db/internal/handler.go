package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTaskHandler(c *gin.Context) {
	var input struct {
		Title string `json:"title"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат запроса"})
		return
	}

	task, err := CreateTask(input.Title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании задачи"})
		return
	}
	c.JSON(http.StatusCreated, task)
}

func GetTasksHandler(c *gin.Context) {
	tasks, err := GetTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении задач"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func DeleteTaskHandler(c *gin.Context) {
	var input struct {
		ID string `json:"id"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат запроса"})
		return
	}

	if err := DeleteTask(input.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Задача не найдена"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Задача удалена"})
}

func MarkTaskDoneHandler(c *gin.Context) {
	var input struct {
		ID string `json:"id"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат запроса"})
		return
	}
	task, err := MarkTaskDone(input.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Задача не найдена"})
		return
	}
	c.JSON(http.StatusOK, task)
}
