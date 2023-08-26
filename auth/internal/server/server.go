package server

import "github.com/evgsrkn/go-microservices-example/auth/internal/auth"

type (
	AuthHandler auth.IHandler

	API struct {
		AuthHandler
	}
)

func NewAPI(auth auth.IHandler) *API {
	return &API{auth}
}
