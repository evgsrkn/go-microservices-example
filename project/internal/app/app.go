package app

import (
	"fmt"
	"net"
	"sync"

	"github.com/evgsrkn/go-microservices-example/project/internal/config"
	"github.com/evgsrkn/go-microservices-example/project/internal/database"
	"github.com/evgsrkn/go-microservices-example/project/internal/project"
	"github.com/evgsrkn/go-microservices-example/project/internal/server"
	"github.com/evgsrkn/go-microservices-example/project/pkg/pb"

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

			fx.Annotate(project.NewStorage, fx.As(new(project.IRepository))),
			fx.Annotate(project.NewService, fx.As(new(project.IService))),
			fx.Annotate(project.NewHandler, fx.As(new(project.IHandler))),

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
	pb.RegisterProjectServiceServer(srv, api)
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
