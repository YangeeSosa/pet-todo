package internal

import "github.com/google/uuid"

func CreateTask(title string) (Task, error) {
	id := uuid.New()
	task := Task{
		ID:    id.String(),
		Title: title,
		Done:  false,
	}
	_, err := DB.Exec("INSERT INTO tasks (id, title, done) VALUES ($1, $2, $3)", task.ID, task.Title, task.Done)
	return task, err
}

func GetTasks() ([]Task, error) {
	var tasks []Task
	err := DB.Select(&tasks, "SELECT * FROM tasks")
	return tasks, err
}

func DeleteTask(id string) error {
	_, err := DB.Exec("DELETE FROM tasks WHERE id = $1", id)
	return err
}

func MarkTaskDone(id string) (Task, error) {
	_, err := DB.Exec("UPDATE tasks SET done = true WHERE id = $1", id)
	if err != nil {
		return Task{}, err
	}
	var task Task
	err = DB.Get(&task, "SELECT id, title, done FROM tasks WHERE id = $1", id)
	return task, err
}
