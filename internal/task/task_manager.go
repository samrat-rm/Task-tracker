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
	Tasks       map[int64]Task
	TaskStorage *TaskStorageStruct
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
	tm.Tasks[id] = task
	err := tm.TaskStorage.SaveTasksToJson(tm.Tasks)
	if err != nil {
		return 0, fmt.Errorf("failed to save new Task %w", err)
	}
	return id, nil
}

func (tm *TaskManager) UpdateTaskStatus(id int64, status Status) error {
	val, ok := tm.Tasks[id]
	if !ok {
		return fmt.Errorf("update status failed for task with ID: %d", id)
	}
	val.Status = status
	tm.Tasks[id] = val
	err := tm.TaskStorage.SaveTasksToJson(tm.Tasks)
	if err != nil {
		return fmt.Errorf("failed to save new Task %w", err)
	}
	return nil
}

func (tm *TaskManager) DeleteTask(id int64) error {
	_, ok := tm.Tasks[id]
	if !ok {
		return fmt.Errorf("task with ID %d not found", id)
	}

	delete(tm.Tasks, id)

	err := tm.TaskStorage.SaveTasksToJson(tm.Tasks)
	if err != nil {
		return fmt.Errorf("failed to save after task deletion: %w", err)
	}
	return nil
}

func (tm *TaskManager) UpdateTaskDescription(id int64, description string) error {
	task, ok := tm.Tasks[id]
	if !ok {
		return fmt.Errorf("task with ID %d not found", id)
	}

	task.Description = description
	task.UpdatedAt = time.Now()
	tm.Tasks[id] = task

	err := tm.TaskStorage.SaveTasksToJson(tm.Tasks)
	if err != nil {
		return fmt.Errorf("failed to save task after description update: %w", err)
	}
	return nil
}

func (tm *TaskManager) GetAllTasks() ([]Task, error) {
	tasks := []Task{}
	for _, task := range tm.Tasks {
		tasks = append(tasks, task)
	}
	if len(tasks) == 0 {
		return []Task{}, fmt.Errorf("no task found in the JSON file path %s ", tm.TaskStorage.FilePath)
	}
	return tasks, nil
}

func (tm *TaskManager) GetTasksByStatus(status Status) ([]Task, error) {
	tasks := []Task{}
	for _, task := range tm.Tasks {
		if task.Status == status {
			tasks = append(tasks, task)
		}
	}
	if len(tasks) == 0 {
		return []Task{}, fmt.Errorf("no task with the %s status found in the JSON file path %s ", statusToString[status], tm.TaskStorage.FilePath)
	}
	return tasks, nil
}

func InitTaskManager(filePath string) *TaskManager {
	ts := &TaskStorageStruct{
		FilePath: filePath,
	}
	tasks, err := ts.FetchTasksFromJson()
	if err != nil {
		fmt.Printf("Failed to get Json data from %s file path, %s", filePath, err)
	}

	return &TaskManager{
		Tasks:       tasks,
		TaskStorage: ts,
	}
}
