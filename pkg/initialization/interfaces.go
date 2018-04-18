package initialization

import "github.com/grafana/grafana/pkg/services/dashboards"

type SetDashboardProvisioningService interface {
	SetDashboardProvisioningService(dashboards.DashboardProvisioningService)
}

type Service interface{}
