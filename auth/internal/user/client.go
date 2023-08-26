//go:generate mockgen -source client.go -destination ./../../mocks/user/client.go -package user_mock
package user

import (
	"github.com/evgsrkn/go-microservices-example/auth/internal/config"
	"github.com/evgsrkn/go-microservices-example/user/pkg/pb"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client pb.UserServiceClient

func NewClient(cfg *config.Cfg) (Client, error) {
	conn, err := grpc.Dial(
		cfg.Services.User.Host,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, errors.Wrap(err, "can't establish gppc connection")
	}

	client := pb.NewUserServiceClient(conn)

	return client, nil
}
