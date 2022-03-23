//go:build wireinject

package main

import (
	"context"

	"github.com/comeonjy/go-kit/app"
	"github.com/comeonjy/go-kit/pkg/xlog"
	"github.com/comeonjy/go-layout/internal/config"

	"github.com/comeonjy/go-layout/internal/domain/aggregate"
	"github.com/comeonjy/go-layout/internal/infra/persistence"
	"github.com/comeonjy/go-layout/internal/server"
	"github.com/comeonjy/go-layout/internal/service"
	"github.com/google/wire"
)

func InitApp(ctx context.Context, logger *xlog.Logger) *app.App {
	panic(wire.Build(
		server.ProviderSet,
		service.ProviderSet,
		aggregate.ProviderSet,
		persistence.ProviderSet,
		config.ProviderSet,
		app.NewApp,
	))
}
