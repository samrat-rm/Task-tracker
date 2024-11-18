package task

import (
	"encoding/json"
	"os"
	"testing"
)

func TestTaskManager(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "tasks_*.json")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	tasks := map[int64]Task{
		1: {Id: 1, Description: "Test Task", Status: 1},
		2: {Id: 2, Description: "Test Task 2", Status: 2},
	}

	ts := &TaskStorageStruct{filePath: tmpFile.Name()}

	err = ts.SaveTasksToJson(tasks)
	if err != nil {
		t.Errorf("SaveTasksToJson returned an error: %v", err)
	}

	tm := &TaskManager{
		taskStorage: ts,
		tasks:       tasks,
	}

	createTaskId := testCreateTask(t, tm, tmpFile.Name())
	isTaskCreated := false

	for _, task := range tasks {
		if task.Id == createTaskId {
			isTaskCreated = true
		}
	}

	if !isTaskCreated {
		t.Errorf("created new task is not saved in Json file")
	}

	const deleteTaskId int64 = 1
	testDeleteTask(t, tm, tmpFile.Name(), deleteTaskId)

	for _, task := range tasks {
		if task.Id == deleteTaskId {
			t.Errorf("Task is not deleted from Json file, Id: %d", deleteTaskId)
		}
	}

	const updateTaskId int64 = 2
	const newDescription string = "updated task"

	testUpdateTaskDescription(t, tm, tmpFile.Name(), updateTaskId, newDescription)
	for _, task := range tasks {
		if task.Id == updateTaskId {
			if task.Description != newDescription {
				t.Errorf("Task description is not updated for Task Id : %d", updateTaskId)
			}
		}
	}

	const updateStatus Status = 1
	testUpdateTaskStatus(t, tm, tmpFile.Name(), updateTaskId, updateStatus)
	for _, task := range tasks {
		if task.Id == updateTaskId {
			if task.Status != updateStatus {
				t.Errorf("Task status is not updated for Task Id : %d", updateTaskId)
			}
		}
	}

	allTasks := testGetAllTasks(t, tm)
	if len(allTasks) != len(tasks) {
		t.Errorf("expected %d tasks, got %d", len(tasks), len(allTasks))
	}

	for _, task := range allTasks {
		if _, exists := tasks[task.Id]; !exists {
			t.Errorf("unexpected task found: %+v", task)
		}
	}

	//
	statusToFetch := Status(1)
	filteredTasks := testGetTasksByStatus(t, tm, statusToFetch)
	expectedCount := 0
	for _, task := range tasks {
		if task.Status == statusToFetch {
			expectedCount++
		}
	}

	if len(filteredTasks) != expectedCount {
		t.Errorf("expected %d tasks with status %d, got %d", expectedCount, statusToFetch, len(filteredTasks))
	}

	for _, task := range filteredTasks {
		if task.Status != statusToFetch {
			t.Errorf("unexpected task with status %d found: %+v", task.Status, task)
		}
	}
}

func testCreateTask(t *testing.T, tm *TaskManager, tempFileName string) int64 {

	id, err := tm.CreateTask("new task")
	if err != nil {
		t.Errorf("failed to create new task")
	}

	tasks := getTasksFromJsonFile(t, tempFileName)
	isSavedToJson := false
	for _, task := range tasks {
		if task.Id == id {
			isSavedToJson = true
		}
	}
	if !isSavedToJson {
		t.Errorf("task %d not found in json file", id)
	}

	return id
}

func testDeleteTask(t *testing.T, tm *TaskManager, tempFileName string, deleteTaskId int64) {

	err := tm.DeleteTask(deleteTaskId)
	if err != nil {
		t.Errorf("failed to delete task with Id %d", deleteTaskId)
	}

	tasks := getTasksFromJsonFile(t, tempFileName)

	for _, task := range tasks {
		if task.Id == deleteTaskId {
			t.Errorf("failed to delete, task %d found in json file", deleteTaskId)
		}
	}
}

func testUpdateTaskDescription(t *testing.T, tm *TaskManager, tempFileName string, id int64, description string) {
	err := tm.UpdateTaskDescription(id, description)
	if err != nil {
		t.Errorf("failed to update task with Id %d", id)
	}
	tasks := getTasksFromJsonFile(t, tempFileName)

	for _, task := range tasks {
		if task.Id == id {
			if task.Description != description {
				t.Errorf("failed to update description for task %d ", id)
			}
		}
	}
}

func testUpdateTaskStatus(t *testing.T, tm *TaskManager, tempFileName string, id int64, status Status) {
	err := tm.UpdateTaskStatus(id, status)
	if err != nil {
		t.Errorf("failed to update task status with Id %d", id)
	}
	tasks := getTasksFromJsonFile(t, tempFileName)
	for _, task := range tasks {
		if task.Id == id {
			if task.Status != status {
				t.Errorf("failed to update status for task %d ", id)
			}
		}
	}
}

func testGetAllTasks(t *testing.T, tm *TaskManager) []Task {
	allTasks, err := tm.GetAllTasks()
	if err != nil {
		t.Errorf("GetAllTasks returned an error: %v", err)
	}
	return allTasks
}

func testGetTasksByStatus(t *testing.T, tm *TaskManager, statusToFetch Status) []Task {

	filteredTasks, err := tm.GetTasksByStatus(statusToFetch)
	if err != nil {
		t.Errorf("GetTasksByStatus returned an error: %v", err)
	}

	return filteredTasks
}

func getTasksFromJsonFile(t *testing.T, tempFileName string) []Task {
	content, err := os.ReadFile(tempFileName)
	if err != nil {
		t.Fatalf("failed to read temp file: %v", err)
	}

	var actual []Task
	err = json.Unmarshal(content, &actual)
	if err != nil {
		t.Fatalf("failed to unmarshal JSON content: %v", err)
	}
	return actual
}
