package interfaces

import "alfath_lms/api/definitions"

type DashboardServiceInterface interface {
	GetDashboardData() (definitions.SimpleDashboardData, error)
}
