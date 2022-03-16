//go:build wireinject
// +build wireinject

package cmd

import (
	"context"

	"github.com/comeonjy/go-kit/pkg/xlog"
	"github.com/comeonjy/go-layout/configs"
	"github.com/comeonjy/go-layout/internal/domain/aggregate"
	"github.com/comeonjy/go-layout/internal/infra/persistence"
	"github.com/comeonjy/go-layout/internal/service"
	"github.com/google/wire"

	"github.com/comeonjy/go-layout/internal/server"
)

func InitApp(ctx context.Context, logger *xlog.Logger) *App {
	panic(wire.Build(
		server.ProviderSet,
		service.ProviderSet,
		aggregate.ProviderSet,
		persistence.ProviderSet,
		newApp,
		configs.ProviderSet,
	))
}
