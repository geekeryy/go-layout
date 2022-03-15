//go:build wireinject
// +build wireinject

package cmd

import (
	"context"

	"github.com/comeonjy/go-kit/pkg/xlog"
	"github.com/comeonjy/go-layout/configs"
	"github.com/comeonjy/go-layout/internal/infrastructure/data"
	"github.com/google/wire"

	"github.com/comeonjy/go-layout/internal/server"
)

func InitApp(ctx context.Context, logger *xlog.Logger) *App {
	panic(wire.Build(
		server.ProviderSet,
		application.ProviderSet,
		newApp,
		configs.ProviderSet,
		data.ProviderSet,
	))
}
