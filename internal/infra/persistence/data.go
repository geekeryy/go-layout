package persistence

import (
	"github.com/comeonjy/go-layout/configs"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewData, NewWorkRepo)

type Data struct {
}

func NewData(cfg configs.Interface) *Data {
	return &Data{}
}
