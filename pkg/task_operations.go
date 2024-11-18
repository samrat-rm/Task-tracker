package task

type TaskOperations interface {
	CreateTask(description string) (int64, error)
	DeleteTask(id int64) error
	UpdateTaskDescription(id int64, description string) error
	UpdateTaskStatus(id int64, status Status) error
	GetAllTasks() ([]Task, error)
	GetTasksByStatus(status Status) ([]Task, error)
}
