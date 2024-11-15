package task

type TaskOperations interface {
	CreateTask(description string) int64
	DeleteTask(id int64)
	UpdateTaskDescription(id int64, description string)
	UpdateTaskStatus(id int64, status Status)
}
