package server

import "github.com/evgsrkn/go-microservices-example/user/internal/user"

type (
	TaskHandler user.IHandler

	API struct {
		TaskHandler
	}
)

func NewAPI(user user.IHandler) *API {
	return &API{user}
}
