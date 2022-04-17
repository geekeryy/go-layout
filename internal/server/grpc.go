package server

import (
	"time"

	"github.com/comeonjy/go-kit/grpc/reloadconfig"
	"github.com/comeonjy/go-kit/pkg/xenv"
	"github.com/comeonjy/go-kit/pkg/xlog"
	"github.com/comeonjy/go-kit/pkg/xmiddleware"
	"github.com/comeonjy/go-layout/config"
	"github.com/google/wire"
	"google.golang.org/grpc"

	"github.com/comeonjy/go-layout/api/v1"
	"github.com/comeonjy/go-layout/internal/service"
)

var ProviderSet = wire.NewSet(NewGrpcServer, NewHttpServer)

func NewGrpcServer(srv *service.SchedulerService, conf config.Interface, logger *xlog.Logger) *grpc.Server {
	server := grpc.NewServer(
		grpc.ConnectionTimeout(2*time.Second),
		grpc.ChainUnaryInterceptor(
			xmiddleware.GrpcLogger(xenv.GetEnv(xenv.TraceName), logger), xmiddleware.GrpcValidate, xmiddleware.GrpcRecover(logger), xmiddleware.GrpcAuth, xmiddleware.GrpcApm(conf.Get().ApmUrl, xenv.GetEnv(xenv.AppName), xenv.GetEnv(xenv.AppVersion), xenv.GetEnv(xenv.AppEnv))),
	)
	v1.RegisterSchedulerServer(server, srv)
	reloadconfig.RegisterReloadConfigServer(server, reloadconfig.NewServer(conf))
	return server
}
