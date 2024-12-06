package dashboardservice

import (
	"dashboard-ecommerce-team2/repository"

	"go.uber.org/zap"
)

type DashboardService interface {
	CountCustomer() error
	CountOrder() error
	CountTotalPriceOrder() error
	CountProduct() error
	CurrentMonthEarning() error
	RenevueChart() error
	GetBestItemList() error
}

type dashboardService struct {
	Repo repository.Repository
	Log  *zap.Logger
}

// CountCustomer implements DashboardService.
func (d *dashboardService) CountCustomer() error {
	panic("unimplemented")
}

// CountOrder implements DashboardService.
func (d *dashboardService) CountOrder() error {
	panic("unimplemented")
}

// CountProduct implements DashboardService.
func (d *dashboardService) CountProduct() error {
	panic("unimplemented")
}

// CountTotalPriceOrder implements DashboardService.
func (d *dashboardService) CountTotalPriceOrder() error {
	panic("unimplemented")
}

// CurrentMonthEarning implements DashboardService.
func (d *dashboardService) CurrentMonthEarning() error {
	panic("unimplemented")
}

// GetBestItemList implements DashboardService.
func (d *dashboardService) GetBestItemList() error {
	panic("unimplemented")
}

// RenevueChart implements DashboardService.
func (d *dashboardService) RenevueChart() error {
	panic("unimplemented")
}

func NewDashboardService(repo repository.Repository, log *zap.Logger) DashboardService {
	return &dashboardService{Repo: repo, Log: log}
}
