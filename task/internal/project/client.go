package project

import (
	"github.com/evgsrkn/go-microservices-example/project/pkg/pb"
	"github.com/evgsrkn/go-microservices-example/task/internal/config"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client pb.ProjectServiceClient

func NewClient(cfg *config.Cfg) (Client, error) {
	conn, err := grpc.Dial(
		cfg.Services.Project.Host,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, errors.Wrap(err, "can't establish gppc connection")
	}

	client := pb.NewProjectServiceClient(conn)

	return client, nil
}
