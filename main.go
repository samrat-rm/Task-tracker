package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	task "github.com/samrat-rm/task_tracker/internal/task"
)

const taskCLI = "./task-cli"
const filePath = "data/data.json"

func main() {

	program := os.Args[0]
	if program != taskCLI {
		log.Printf("Invalid program name: %s. Expected: %s. Exiting.", program, taskCLI)
		return
	}
	command := os.Args[1]
	args := os.Args[2:]
	tm := task.InitTaskManager(filePath)
	var id int64
	var err error

	if command != "add" && command != "list" {
		id, err = strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Printf("Task ID format is invalid :%d\n", id)
			return
		}
	}

	switch command {
	case "add":
		id, err = tm.CreateTask(args[0])
		if err != nil {
			log.Printf("Failed to create task with argument '%s': %v", args[0], err)
		}
		log.Printf("Task created with Id :%d\n", id)

	case "update":

		err = tm.UpdateTaskDescription(id, args[1])
		if err != nil {
			log.Printf("Failed to update the task %d \n", id)
			return
		}
		log.Printf("Task updated with Id :%d\n", id)

	case "delete":
		err = tm.DeleteTask(id)
		if err != nil {
			log.Printf("Failed to delete the task with Id :%d\n", id)
			return
		}
		log.Printf("Task deleted with Id :%d\n", id)

	case "mark-in-progress":
		err = tm.UpdateTaskStatus(id, 1)
		if err != nil {
			log.Printf("Failed to update the task %d status \n", id)
			return
		}
		log.Printf("Task with Id :%d marked as in progress\n", id)

	case "mark-done":
		err = tm.UpdateTaskStatus(id, 2)
		if err != nil {
			log.Printf("Failed to update the task %d status \n", id)
			return
		}
		log.Printf("Task with Id :%d marked as done \n", id)

	case "list":
		var tasks []task.Task
		if len(args) == 0 {
			tasks, _ = tm.GetAllTasks()

		} else if args[0] == "todo" {
			tasks, _ = tm.GetTasksByStatus(0)
		} else if args[0] == "in-progress" {
			tasks, _ = tm.GetTasksByStatus(1)
		} else if args[0] == "done" {
			tasks, _ = tm.GetTasksByStatus(2)
		}
		if len(tasks) > 0 {
			for _, t := range tasks {
				status := mapStatus(t.Status)
				fmt.Printf("%-5d %-50s %-15s %-20s %-20s\n",
					t.Id, t.Description, status, t.CreatedAt.Format("2006-01-02 15:04:05"), t.UpdatedAt.Format("2006-01-02 15:04:05"))
			}
		} else {
			log.Printf("No tasks found")
		}
	default:
		fmt.Println("Unknown command:", command)
	}
}

func mapStatus(status task.Status) string {
	switch status {
	case task.NotStarted:
		return "Not Started"
	case task.InProgress:
		return "In Progress"
	case task.Completed:
		return "Completed"
	default:
		return "Unknown"
	}
}
