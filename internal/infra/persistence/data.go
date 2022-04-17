package persistence

import (
	"github.com/comeonjy/go-layout/config"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewData, NewWorkRepo)

type Data struct {
}

func NewData(cfg config.Interface) *Data {
	return &Data{}
}
