package app

import (
	"fmt"
	"net"
	"sync"

	"github.com/evgsrkn/go-microservices-example/auth/internal/auth"
	"github.com/evgsrkn/go-microservices-example/auth/internal/config"
	"github.com/evgsrkn/go-microservices-example/auth/internal/database"
	"github.com/evgsrkn/go-microservices-example/auth/internal/server"
	"github.com/evgsrkn/go-microservices-example/auth/internal/user"
	"github.com/evgsrkn/go-microservices-example/auth/pkg/pb"
	"github.com/evgsrkn/go-microservices-example/gateway/pkg/logger"
	"github.com/evgsrkn/go-microservices-example/gateway/pkg/rpc"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Run() {
	fx.New(CreateApp()).Run()
}

func CreateApp() fx.Option {
	return fx.Options(
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
		fx.Provide(
			logger.New,
			config.New,
			database.New,
			func() *sync.Mutex {
				var mu sync.Mutex
				return &mu
			},

			fx.Annotate(auth.NewStorage, fx.As(new(auth.Repository))),
			fx.Annotate(auth.NewService, fx.As(new(auth.IService))),
			fx.Annotate(auth.NewHandler, fx.As(new(auth.IHandler))),
			fx.Annotate(user.NewClient, fx.As(new(user.Client))),

			server.NewAPI,

			func(logger *zap.Logger) *grpc.Server {
				return rpc.New(
					rpc.WithZapLogger(logger),
				)
			},
		),
		fx.Invoke(serve),
	)
}

func serve(logger *zap.Logger, srv *grpc.Server, cfg *config.Cfg, api *server.API) {
	pb.RegisterAuthServer(srv, api)
	reflection.Register(srv)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Port))
	if err != nil {
		logger.Error("cannot bind server", zap.Error(err))
		return
	}

	if errServe := srv.Serve(lis); errServe != nil {
		logger.Error("cannot listen server", zap.Error(err))
		return
	}
}
