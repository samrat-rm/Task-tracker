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
	FilePath string
}

func (t *TaskStorageStruct) FetchTasksFromJson() (map[int64]Task, error) {

	file, err := os.OpenFile(t.FilePath, os.O_RDWR|os.O_CREATE, 0644)
	taskMap := make(map[int64]Task)

	if err != nil {
		return taskMap, fmt.Errorf("could not open file %s: %v", t.FilePath, err)
	}
	defer file.Close()

	// Check if the file is empty
	stat, err := file.Stat()
	if err != nil {
		return taskMap, fmt.Errorf("could not get file stats %s: %v", t.FilePath, err)
	}
	if stat.Size() == 0 {
		return taskMap, nil
	}

	fileContent, err := io.ReadAll(file)
	if err != nil {
		return taskMap, fmt.Errorf("could not read file %s: %v]", t.FilePath, err)
	}

	var tasks []Task
	err = json.Unmarshal(fileContent, &tasks)
	if err != nil {
		return taskMap, fmt.Errorf("could not unmarshal JSON: %v", err)
	}

	for _, task := range tasks {
		taskMap[task.Id] = task
	}

	return taskMap, nil
}

func (t *TaskStorageStruct) SaveTasksToJson(tasks map[int64]Task) error {
	var taskList []Task
	for _, task := range tasks {
		taskList = append(taskList, task)
	}

	fileContent, err := json.MarshalIndent(taskList, "", "  ")
	if err != nil {
		return fmt.Errorf("could not marshal tasks: %v", err)
	}

	err = os.WriteFile(t.FilePath, fileContent, 0644)
	if err != nil {
		return fmt.Errorf("could not write to file %s: %v", t.FilePath, err)
	}

	return nil
}
