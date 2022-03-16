package service

import (
	"context"
	"log"

	"github.com/comeonjy/go-layout/internal/domain/aggregate"
	"github.com/google/wire"
	"google.golang.org/grpc/metadata"

	"github.com/comeonjy/go-kit/pkg/xlog"
	v1 "github.com/comeonjy/go-layout/api/v1"
	"github.com/comeonjy/go-layout/configs"
)

var ProviderSet = wire.NewSet(NewSchedulerService)

type SchedulerService struct {
	v1.UnimplementedSchedulerServer
	conf        configs.Interface
	logger      *xlog.Logger
	workUseCase *aggregate.WorkUseCase
}

func NewSchedulerService(conf configs.Interface, logger *xlog.Logger, workUseCase *aggregate.WorkUseCase) *SchedulerService {
	return &SchedulerService{
		conf:        conf,
		workUseCase: workUseCase,
		logger:      logger,
	}
}

func (svc *SchedulerService) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	if mdIn, ok := metadata.FromIncomingContext(ctx); ok {
		mdIn.Get("")
	}
	return ctx, nil
}

func (svc *SchedulerService) Ping(ctx context.Context, in *v1.Empty) (*v1.Result, error) {
	info, err := svc.workUseCase.GetInfo(1)
	if err != nil {
		return nil, err
	}
	log.Println(info)
	return &v1.Result{
		Code:    200,
		Message: svc.conf.Get().Mode,
		Data:    nil,
	}, nil
}
