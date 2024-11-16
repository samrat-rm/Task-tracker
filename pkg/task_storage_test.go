package task

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

func TestFetchTasksFromJson(t *testing.T) {

	tmpFile, err := os.CreateTemp("", "tasks_*.json")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	sampleData := `[{"Id":1, "Description":"Test Task", "Status":1}]`
	_, err = tmpFile.Write([]byte(sampleData))
	if err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	ts := &TaskStorageStruct{filePath: tmpFile.Name()}

	tasks, err := ts.FetchTasksFromJson()
	if err != nil {
		t.Errorf("FetchTasksFromJson returned an error: %v", err)
	}

	if len(tasks) != 1 || tasks[1].Description != "Test Task" {
		t.Errorf("FetchTasksFromJson did not return expected tasks, got: %v", tasks)
	}
}
func TestSaveTasksToJson(t *testing.T) {

	tmpFile, err := os.CreateTemp("", "tasks_*.json")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	ts := &TaskStorageStruct{filePath: tmpFile.Name()}

	tasks := map[int64]Task{
		1: {Id: 1, Description: "Test Task", Status: 1},
	}

	err = ts.SaveTasksToJson(tasks)
	if err != nil {
		t.Errorf("SaveTasksToJson returned an error: %v", err)
	}

	content, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("failed to read temp file: %v", err)
	}

	expected := []Task{
		{Id: 1, Description: "Test Task", Status: 1},
	}

	var actual []Task
	err = json.Unmarshal(content, &actual)
	if err != nil {
		t.Fatalf("failed to unmarshal JSON content: %v", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("SaveTasksToJson did not write expected %+v, got: %+v", expected, actual)
	}
}
