package task

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

func TestFetchTasksFromJson(t *testing.T) {
	// Setup: Create a temporary file

	// notStarted := 1
	tmpFile, err := os.CreateTemp("", "tasks_*.json")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name()) // Clean up after test

	// Write sample valid JSON data to the file
	sampleData := `[{"Id":1, "Description":"Test Task", "Status":1}]`
	_, err = tmpFile.Write([]byte(sampleData))
	if err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	// Initialize TaskStorageStruct
	ts := &TaskStorageStruct{filePath: tmpFile.Name()}

	// Run FetchTasksFromJson
	tasks, err := ts.FetchTasksFromJson()
	if err != nil {
		t.Errorf("FetchTasksFromJson returned an error: %v", err)
	}

	// Validate the results
	if len(tasks) != 1 || tasks[1].Description != "Test Task" {
		t.Errorf("FetchTasksFromJson did not return expected tasks, got: %v", tasks)
	}
}
func TestSaveTasksToJson(t *testing.T) {
	// Setup: Create a temporary file
	tmpFile, err := os.CreateTemp("", "tasks_*.json")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name()) // Clean up after test

	// Initialize TaskStorageStruct
	ts := &TaskStorageStruct{filePath: tmpFile.Name()}

	// Create sample tasks to save
	tasks := map[int64]Task{
		1: {Id: 1, Description: "Test Task", Status: 1},
	}

	// Run SaveTasksToJson
	err = ts.SaveTasksToJson(tasks)
	if err != nil {
		t.Errorf("SaveTasksToJson returned an error: %v", err)
	}

	// Read the file content
	content, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("failed to read temp file: %v", err)
	}

	// Define the expected structure
	expected := []Task{
		{Id: 1, Description: "Test Task", Status: 1},
	}

	var actual []Task
	err = json.Unmarshal(content, &actual)
	if err != nil {
		t.Fatalf("failed to unmarshal JSON content: %v", err)
	}

	// Compare the actual structure with the expected one
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("SaveTasksToJson did not write expected %+v, got: %+v", expected, actual)
	}
}
