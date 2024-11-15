package task

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type TaskStorage interface {
	FetchTaskFromJson(filePath string) (map[int64]Task, error)
	SaveTasksToJson(tasks map[int64]Task) error
}

type TaskStorageStruct struct {
	filePath string
}

// FetchTasksFromJson reads tasks from a JSON file and returns them as a map.
func (t *TaskStorageStruct) FetchTasksFromJson() (map[int64]Task, error) {
	file, err := os.Open(t.filePath)
	if err != nil {
		return nil, fmt.Errorf("could not open file %s: %v", t.filePath, err)
	}
	defer file.Close()

	fileContent, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("could not read file %s: %v", t.filePath, err)
	}

	var tasks []Task
	err = json.Unmarshal(fileContent, &tasks)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal JSON: %v", err)
	}

	taskMap := make(map[int64]Task)
	for _, task := range tasks {
		taskMap[task.Id] = task
	}

	return taskMap, nil
}

// SaveTasksToJson saves a map of tasks into a JSON file.
func (t *TaskStorageStruct) SaveTasksToJson(tasks map[int64]Task) error {
	// Convert the map of tasks into a slice
	var taskList []Task
	for _, task := range tasks {
		taskList = append(taskList, task)
	}

	// Marshal the tasks to JSON
	fileContent, err := json.MarshalIndent(taskList, "", "  ")
	if err != nil {
		return fmt.Errorf("could not marshal tasks: %v", err)
	}

	// Write the JSON content to the file
	err = os.WriteFile(t.filePath, fileContent, 0644)
	if err != nil {
		return fmt.Errorf("could not write to file %s: %v", t.filePath, err)
	}

	return nil
}
