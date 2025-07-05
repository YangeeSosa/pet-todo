package internal

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	
	r.POST("/tasks", CreateTaskHandler)
	r.GET("/tasks", GetTasksHandler)
	r.DELETE("/tasks", DeleteTaskHandler)
	r.PUT("/tasks", MarkTaskDoneHandler)
	
	return r
}

func TestCreateTaskHandler(t *testing.T) {
	r := setupTestRouter()
	
	reqBody := map[string]string{"title": "Тестовая задача"}
	jsonBody, _ := json.Marshal(reqBody)
	
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusCreated, w.Code)
	
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response, "id")
	assert.Equal(t, "Тестовая задача", response["title"])
	assert.Equal(t, false, response["done"])
}

func TestGetTasksHandler(t *testing.T) {
	r := setupTestRouter()
	
	req, _ := http.NewRequest("GET", "/tasks", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusOK, w.Code)
	
	var response []map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.IsType(t, []map[string]interface{}{}, response)
}

func TestDeleteTaskHandler(t *testing.T) {
	r := setupTestRouter()
	
	createReqBody := map[string]string{"title": "Задача для удаления"}
	createJsonBody, _ := json.Marshal(createReqBody)
	
	createReq, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(createJsonBody))
	createReq.Header.Set("Content-Type", "application/json")
	
	createW := httptest.NewRecorder()
	r.ServeHTTP(createW, createReq)
	
	var createResponse map[string]interface{}
	json.Unmarshal(createW.Body.Bytes(), &createResponse)
	taskID := createResponse["id"].(string)
	
	deleteReqBody := map[string]string{"id": taskID}
	deleteJsonBody, _ := json.Marshal(deleteReqBody)
	
	deleteReq, _ := http.NewRequest("DELETE", "/tasks", bytes.NewBuffer(deleteJsonBody))
	deleteReq.Header.Set("Content-Type", "application/json")
	
	deleteW := httptest.NewRecorder()
	r.ServeHTTP(deleteW, deleteReq)
	
	assert.Equal(t, http.StatusOK, deleteW.Code)
}

func TestMarkTaskDoneHandler(t *testing.T) {
	r := setupTestRouter()
	
	createReqBody := map[string]string{"title": "Задача для выполнения"}
	createJsonBody, _ := json.Marshal(createReqBody)
	
	createReq, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(createJsonBody))
	createReq.Header.Set("Content-Type", "application/json")
	
	createW := httptest.NewRecorder()
	r.ServeHTTP(createW, createReq)
	
	var createResponse map[string]interface{}
	json.Unmarshal(createW.Body.Bytes(), &createResponse)
	taskID := createResponse["id"].(string)
	
	updateReqBody := map[string]string{"id": taskID}
	updateJsonBody, _ := json.Marshal(updateReqBody)
	
	updateReq, _ := http.NewRequest("PUT", "/tasks", bytes.NewBuffer(updateJsonBody))
	updateReq.Header.Set("Content-Type", "application/json")
	
	updateW := httptest.NewRecorder()
	r.ServeHTTP(updateW, updateReq)
	
	assert.Equal(t, http.StatusOK, updateW.Code)
	
	var updateResponse map[string]interface{}
	err := json.Unmarshal(updateW.Body.Bytes(), &updateResponse)
	assert.NoError(t, err)
	assert.Equal(t, true, updateResponse["done"])
} 