package server

import "task/internal/task"

type (
	TaskHandler task.IHandler

	API struct {
		TaskHandler
	}
)

func NewAPI(task task.IHandler) *API {
	return &API{task}
}
