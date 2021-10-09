//go:build wireinject
// +build wireinject

package cmd

import (
	"context"

	"github.com/comeonjy/go-kit/pkg/xlog"
	"github.com/comeonjy/go-layout/configs"
	"github.com/google/wire"

	"github.com/comeonjy/go-layout/internal/data"
	"github.com/comeonjy/go-layout/internal/server"
	"github.com/comeonjy/go-layout/internal/service"
)

func InitApp(ctx context.Context,logger *xlog.Logger) *App {
	panic(wire.Build(
		server.ProviderSet,
		service.ProviderSet,
		newApp,
		configs.ProviderSet,
		data.ProviderSet,
	))
}
