package server

import (
	"time"

	"github.com/comeonjy/go-kit/pkg/xenv"
	"github.com/comeonjy/go-kit/pkg/xlog"
	"github.com/comeonjy/go-kit/pkg/xmiddleware"
	"github.com/google/wire"
	"google.golang.org/grpc"

	"github.com/comeonjy/go-layout/api/v1"
	"github.com/comeonjy/go-layout/configs"
	"github.com/comeonjy/go-layout/internal/service"
	"github.com/comeonjy/go-layout/pkg/consts"
)

var ProviderSet = wire.NewSet(NewGrpcServer, NewHttpServer)

func NewGrpcServer(srv *service.SchedulerService, conf configs.Interface,logger *xlog.Logger) *grpc.Server {
	server := grpc.NewServer(
		grpc.ConnectionTimeout(2*time.Second),
		grpc.ChainUnaryInterceptor(
			xmiddleware.GrpcLogger(consts.TraceName,logger), xmiddleware.GrpcValidate, xmiddleware.GrpcRecover(logger), xmiddleware.GrpcAuth, xmiddleware.GrpcApm(conf.Get().ApmUrl, consts.AppName, consts.AppVersion, xenv.GetEnv(consts.AppEnv))),
	)
	v1.RegisterSchedulerServer(server, srv)
	return server
}
