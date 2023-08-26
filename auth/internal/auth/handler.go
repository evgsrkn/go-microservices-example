package auth

import (
	"context"
	"errors"

	"github.com/evgsrkn/go-microservices-example/auth/pkg/pb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type (
	authHandler struct {
		service IService
		log     *zap.Logger
		server  *grpc.Server
	}

	IHandler interface {
		Login(c context.Context, creds *pb.Credentials) (*pb.Token, error)
		Register(c context.Context, creds *pb.BasicCredentials) (*pb.Token, error)
	}
)

func NewHandler(service IService, log *zap.Logger, server *grpc.Server) IHandler {
	return &authHandler{service, log, server}
}

func (h *authHandler) Login(ctx context.Context, creds *pb.Credentials) (*pb.Token, error) {
	if !h.service.IsUserExist(creds.Username) {
		return nil, errors.New("user doesn't exists")
	}

	token, err := h.service.ValidateUser(creds)
	if err != nil {
		return nil, err
	}

	return &pb.Token{
		Token: token,
	}, nil
}

func (h *authHandler) Register(c context.Context, creds *pb.BasicCredentials) (*pb.Token, error) {
	if h.service.IsUserExist(creds.Username) {
		return nil, errors.New("user already exist")
	}

	token, err := h.service.RegisterUser(creds)
	if err != nil {
		return nil, err
	}

	return &pb.Token{
		Token: token,
	}, nil
}
