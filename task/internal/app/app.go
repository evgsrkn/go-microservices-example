package app

import (
	"fmt"
	"net"
	"sync"

	"gateway/pkg/logger"
	"gateway/pkg/rpc"
	"task/internal/config"
	"task/internal/database"
	"task/internal/server"
	"task/internal/task"
	"task/pkg/pb"

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

			fx.Annotate(task.NewStorage, fx.As(new(task.IRepository))),
			fx.Annotate(task.NewService, fx.As(new(task.IService))),
			fx.Annotate(task.NewHandler, fx.As(new(task.IHandler))),

			server.NewAPI,

			func(logger *zap.Logger) *grpc.Server {
				return rpc.New(
					rpc.WithLoggerLogrus(logger),
				)
			},
		),
		fx.Invoke(serve),
	)
}

func serve(logger *zap.Logger, srv *grpc.Server, cfg *config.Cfg, api *server.API) {
	pb.RegisterTaskServiceServer(srv, api)
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
