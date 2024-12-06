package dashboardcontroller

import (
	"dashboard-ecommerce-team2/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DashboardController struct {
	Service service.Service
	Log     *zap.Logger
}

func NewDashboardController(service service.Service, log *zap.Logger) *DashboardController {
	return &DashboardController{
		Service: service,
		Log:     log,
	}
}

func (ctrl *DashboardController) GetSummaryController(c *gin.Context)          {}
func (ctrl *DashboardController) CurrentMonthEarningController(c *gin.Context) {}
func (ctrl *DashboardController) RenevueChartController(c *gin.Context)        {}
func (ctrl *DashboardController) GetBestItemListController(c *gin.Context)     {}
