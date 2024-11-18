package task

import (
	"fmt"
	"time"

	"github.com/samrat-rm/task_tracker/utils"
)

var statusToString = map[Status]string{
	NotStarted: "NotStarted",
	InProgress: "InProgress",
	Completed:  "Completed",
}

type TaskManager struct {
	tasks       map[int64]Task
	taskStorage *TaskStorageStruct
}

func (tm *TaskManager) CreateTask(description string) (int64, error) {
	id := int64(utils.GenerateRandomID())
	task := Task{
		Id:          id,
		Description: description,
		Status:      NotStarted,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	tm.tasks[id] = task
	err := tm.taskStorage.SaveTasksToJson(tm.tasks)
	if err != nil {
		return 0, fmt.Errorf("failed to save new Task %w", err)
	}
	return id, nil
}

func (tm *TaskManager) UpdateTaskStatus(id int64, status Status) error {
	val, ok := tm.tasks[id]
	if !ok {
		return fmt.Errorf("update status failed for task with ID: %d", id)
	}
	val.Status = status
	tm.tasks[id] = val
	err := tm.taskStorage.SaveTasksToJson(tm.tasks)
	if err != nil {
		return fmt.Errorf("failed to save new Task %w", err)
	}
	return nil
}

func (tm *TaskManager) DeleteTask(id int64) error {
	_, ok := tm.tasks[id]
	if !ok {
		return fmt.Errorf("task with ID %d not found", id)
	}

	delete(tm.tasks, id)

	err := tm.taskStorage.SaveTasksToJson(tm.tasks)
	if err != nil {
		return fmt.Errorf("failed to save after task deletion: %w", err)
	}
	return nil
}

func (tm *TaskManager) UpdateTaskDescription(id int64, description string) error {
	task, ok := tm.tasks[id]
	if !ok {
		return fmt.Errorf("task with ID %d not found", id)
	}

	task.Description = description
	task.UpdatedAt = time.Now()
	tm.tasks[id] = task

	err := tm.taskStorage.SaveTasksToJson(tm.tasks)
	if err != nil {
		return fmt.Errorf("failed to save task after description update: %w", err)
	}
	return nil
}

func (tm *TaskManager) GetAllTasks() ([]Task, error) {
	tasks := []Task{}
	for _, task := range tm.tasks {
		tasks = append(tasks, task)
	}
	if len(tasks) == 0 {
		return []Task{}, fmt.Errorf("no task found in the JSON file path %s ", tm.taskStorage.filePath)
	}
	return tasks, nil
}

func (tm *TaskManager) GetTasksByStatus(status Status) ([]Task, error) {
	tasks := []Task{}
	for _, task := range tm.tasks {
		if task.Status == status {
			tasks = append(tasks, task)
		}
	}
	if len(tasks) == 0 {
		return []Task{}, fmt.Errorf("no task with the %s status found in the JSON file path %s ", statusToString[status], tm.taskStorage.filePath)
	}
	return tasks, nil
}
