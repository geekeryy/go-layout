package config

import (
	"context"

	"github.com/comeonjy/go-kit/pkg/xconfig"
	"github.com/comeonjy/go-kit/pkg/xconfig/apollo"
	"github.com/comeonjy/go-kit/pkg/xenv"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewConfig)

type Config struct {
	Mode        string `json:"mode"`
	MysqlConf   string `json:"mysql_conf"`
	ApmUrl      string `json:"apm_url"`
	AccountGrpc string `json:"account_grpc"`
}

// Interface 对外暴露接口（用于功能扩展）
type Interface interface {
	Get() Config
	xconfig.ReloadConfigInterface
}

// NewConfig 获取配置
func NewConfig(ctx context.Context) Interface {
	cfg := &config{
		xconfig.New(ctx,
			apollo.NewSource(xenv.GetEnv(xenv.ApolloUrl), xenv.GetEnv(xenv.ApolloAppID), xenv.GetApolloCluster("default"), xenv.GetApolloSecret(), xenv.GetApolloNamespace("grpc"), xenv.GetApolloNamespace("common"))),
	}
	return cfg
}

// 内部配置类载体
type config struct {
	xconfig.IConfig
}

// Get 获取配置
func (c *config) Get() Config {
	conf := Config{}
	c.Scan(&conf)
	return conf
}
