package initialization

import (
	"context"

	"github.com/go-xorm/xorm"
	ini "gopkg.in/ini.v1"
)

type InjectServicesFunc func(s Service)
type InitFunc func(ctx context.Context, engine *xorm.Engine, settings *ini.File, injectDeps InjectServicesFunc) error

var initFuncs = []InitFunc{}

func RegisterInitFunc(fn InitFunc) {
	initFuncs = append(initFuncs, fn)
}

func GetAllInitFuncs() []InitFunc {
	return initFuncs
}
