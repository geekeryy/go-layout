// Package aggregate @Description  TODO
// @Author  	 jiangyang
// @Created  	 2022/3/16 11:21 PM
package aggregate

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewWorkUseCase)
