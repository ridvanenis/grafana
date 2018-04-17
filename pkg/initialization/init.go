package initialization

import (
	"context"

	"github.com/go-xorm/xorm"
	"github.com/grafana/grafana/pkg/services/dashboards"
	ini "gopkg.in/ini.v1"
)

type ServiceLocator interface {
	GetDashboardService() dashboards.DashboardService
	GetDashboardProvisioningService() dashboards.DashboardProvisioningService
}

type InitFunc func(ctx context.Context, engine *xorm.Engine, settings *ini.File, sl ServiceLocator) error

var initFuncs = []InitFunc{}

func RegisterInitFunc(fn InitFunc) {
	initFuncs = append(initFuncs, fn)
}

func GetAllInitFuncs() []InitFunc {
	return initFuncs
}
