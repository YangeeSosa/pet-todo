package internal

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

const dbServiceURL = "http://db-service:8081"

func CreateTaskInDB(title string) (Task, error) {
	reqBody, _ := json.Marshal(map[string]string{
		"title": title,
	})
	resp, err := http.Post(dbServiceURL+"/tasks", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return Task{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return Task{}, errors.New("ошибка при создании задачи")
	}

	var task Task
	if err := json.NewDecoder(resp.Body).Decode(&task); err != nil {
		return Task{}, err
	}
	return task, nil
}

func GetTasksFromDB() ([]Task, error) {
	resp, err := http.Get(dbServiceURL + "/tasks")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("ошибка при получении задач")
	}

	var tasks []Task
	if err := json.NewDecoder(resp.Body).Decode(&tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func DeleteTaskInDB(id string) error {
	reqBody, _ := json.Marshal(map[string]string{
		"id": id,
	})
	req, err := http.NewRequest(http.MethodDelete, dbServiceURL+"/tasks", bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("ошибка при удалении задачи из бд")
	}
	return nil
}

func MarkTaskDoneInDB(id string) (Task, error) {
	reqBody, _ := json.Marshal(map[string]string{
		"id": id,
	})
	req, err := http.NewRequest(http.MethodPut, dbServiceURL+"/tasks", bytes.NewBuffer(reqBody))
	if err != nil {
		return Task{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return Task{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Task{}, errors.New("ошибка при обновлении задачи в бд")
	}

	var task Task
	if err := json.NewDecoder(resp.Body).Decode(&task); err != nil {
		return Task{}, err
	}
	return task, nil
}
