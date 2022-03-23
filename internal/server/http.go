package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/comeonjy/go-kit/pkg/xenv"
	"github.com/comeonjy/go-kit/pkg/xlog"
	"github.com/comeonjy/go-kit/pkg/xmiddleware"
	"github.com/comeonjy/go-layout/api/v1"
	"github.com/comeonjy/go-layout/internal/service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func NewHttpServer(ctx context.Context, logger *xlog.Logger, workingService *service.SchedulerService) *http.Server {
	mux := runtime.NewServeMux(runtime.WithErrorHandler(xmiddleware.HttpErrorHandler(logger)))
	server := http.Server{
		Handler:           xmiddleware.HttpUse(mux, xmiddleware.HttpLogger(xenv.GetEnv(xenv.TraceName), logger)),
		ReadHeaderTimeout: 2 * time.Second,
		WriteTimeout:      5 * time.Second,
	}
	Router(mux, workingService)
	if err := v1.RegisterSchedulerHandlerFromEndpoint(ctx, mux, "localhost:"+xenv.GetEnv(xenv.GrpcPort), []grpc.DialOption{grpc.WithInsecure()}); err != nil {
		panic("RegisterSchedulerHandlerFromEndpoint" + err.Error())
	}
	return &server
}

func Router(mux *runtime.ServeMux, svc *service.SchedulerService) {
	AddRouter(mux, http.MethodGet, "/v2/ping", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		log.Println(pathParams)
		time.Sleep(time.Second * 3)
		log.Println("out...")
		w.Write([]byte(`{"code":100}`))
	})
}

func AddRouter(mux *runtime.ServeMux, meth string, pathPattern string, h runtime.HandlerFunc) {
	if err := mux.HandlePath(meth, pathPattern, h); err != nil {
		panic(err)
	}
}
