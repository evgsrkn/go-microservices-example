package server

import "github.com/evgsrkn/go-microservices-example/task/internal/task"

type (
	TaskHandler task.IHandler

	API struct {
		TaskHandler
	}
)

func NewAPI(task task.IHandler) *API {
	return &API{task}
}
