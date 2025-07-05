package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var input struct {
		Title string `json:"title"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат запроса"})
		return
	}

	task, err := CreateTaskInDB(input.Title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании задачи"})
		return
	}
	SendTaskEvent("task_created", task.ID, task.Title)
	c.JSON(http.StatusCreated, task)
}

func GetTasks(c *gin.Context) {
	tasks, err := GetTasksFromDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении задач"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func DeleteTask(c *gin.Context) {
	var input struct {
		ID string `json:"id"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат запроса"})
		return
	}

	err := DeleteTaskInDB(input.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Задача не найдена"})
		return
	}
	SendTaskEvent("task_deleted", input.ID, "")
	c.JSON(http.StatusOK, gin.H{"message": "Задача удалена"})
}

func MarkTaskDone(c *gin.Context) {
	var input struct {
		ID string `json:"id"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат запроса"})
		return
	}

	task, err := MarkTaskDoneInDB(input.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Задача не найдена"})
		return
	}
	SendTaskEvent("task_completed", task.ID, task.Title)
	c.JSON(http.StatusOK, task)
}
