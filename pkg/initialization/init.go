package initialization

import (
	"context"

	"github.com/go-xorm/xorm"
)

type InitFunc func(ctx context.Context, engine *xorm.Engine) error

var initFuncs = []InitFunc{}

func RegisterInitFunc(fn InitFunc) {
	initFuncs = append(initFuncs, fn)
}

func GetAllInitFuncs() []InitFunc {
	return initFuncs
}
