package task

import (
	"time"
)

type Status int

const (
	NotStarted Status = iota
	InProgress
	Completed
)

type Task struct {
	Id          int64     `json:"id"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
