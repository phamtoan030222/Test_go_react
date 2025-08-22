package storage

import (
	"errors"
	"sync"

	"github.com/phamtoan030222/test/backend/models"
)

var (
	tasks = make(map[int]models.Task)
	nextID = 1
	taskMutex sync.Mutex
)

func CreateTask(title, description string) models.Task {
	taskMutex.Lock()
	defer taskMutex.Unlock()

	task := models.Task{
		ID:          nextID,
		Title:       title,
		Description: description,
		Completed:   false,
	}
	tasks[nextID] = task
	nextID++
	return task
}

func GetAllTasks() []models.Task {
	taskMutex.Lock()
	defer taskMutex.Unlock()

	list := make([]models.Task, 0, len(tasks))
	for _, t := range tasks {
		list = append(list, t)
	}
	return list
}

func UpdateTaskStatus(id int, completed bool) (models.Task, error) {
	taskMutex.Lock()
	defer taskMutex.Unlock()

	task, exists := tasks[id]
	if !exists {
		return models.Task{}, errors.New("not found")
	}

	task.Completed = completed
	tasks[id] = task
	return task, nil
}

func DeleteTask(id int) error {
	taskMutex.Lock()
	defer taskMutex.Unlock()

	_, exists := tasks[id]
	if !exists {
		return errors.New("not found")
	}

	delete(tasks, id)
	return nil
}