package service

import (
	"context"
	"fmt"
	"io"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"google.golang.org/grpc/metadata"

	"github.com/comeonjy/go-kit/pkg/xerror"
	"github.com/comeonjy/go-kit/pkg/xlog"
	v1 "github.com/comeonjy/go-layout/api/v1"
	"github.com/comeonjy/go-layout/configs"
	"github.com/comeonjy/go-layout/internal/data"
	"github.com/comeonjy/go-layout/pkg/errcode"
)

var ProviderSet = wire.NewSet(NewSchedulerService)

type SchedulerService struct {
	v1.UnimplementedSchedulerServer
	conf     configs.Interface
	logger   *xlog.Logger
	workRepo data.WorkRepo
}

func NewSchedulerService(conf configs.Interface, logger *xlog.Logger, workRepo data.WorkRepo) *SchedulerService {
	return &SchedulerService{
		conf:     conf,
		workRepo: workRepo,
		logger:   logger,
	}
}

func (svc *SchedulerService) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	if mdIn, ok := metadata.FromIncomingContext(ctx); ok {
		mdIn.Get("")
	}
	return ctx, nil
}

func (svc *SchedulerService) Ping(ctx context.Context, in *v1.Empty) (*v1.Result, error) {
	return &v1.Result{
		Code:    200,
		Message: "pong" + svc.conf.Get().Mode,
	}, nil
}

func (svc *SchedulerService) QuerySource(ctx context.Context, in *v1.QuerySourceParam) (*v1.SourceInfo, error) {
	var err error
	if in.SourceId%2 == 0 {
		err = xerror.New(errcode.SourceNotFind)
		err = xerror.New(errcode.SourceNotFind, "覆盖错误码信息")
		err = xerror.NewError(errcode.SourceNotFind, "覆盖错误码信息", errors.New("系统记录错误详情1"), errors.WithMessage(io.EOF, "系统记录错误详情2"))
	}
	get, errR := svc.workRepo.Get(int(in.SourceId))
	if errR != nil {
		return nil, errR
	}
	return &v1.SourceInfo{
		SourceId: in.SourceId,
		Text:     fmt.Sprintf("source%v", get),
	}, err
}
