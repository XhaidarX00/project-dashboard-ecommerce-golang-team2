package dashboardcontroller

import (
	"dashboard-ecommerce-team2/helper"
	"dashboard-ecommerce-team2/service"
	"net/http"

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

func (ctrl *DashboardController) GetSummaryController(c *gin.Context) {
	summary, err := ctrl.Service.Dashboard.GetDashboardSummary()
	if err != nil {
		ctrl.Log.Error("Error getting summary", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Error getting summary", http.StatusInternalServerError)
		return
	}

	helper.ResponseOK(c, summary, "dashboard summary successfully retrieved", http.StatusOK)
}
func (ctrl *DashboardController) CurrentMonthEarningController(c *gin.Context) {
	earnings, err := ctrl.Service.Dashboard.CurrentMonthEarning()
	if err != nil {
		ctrl.Log.Error("Error getting earnings", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Error getting earnings", http.StatusInternalServerError)
		return
	}

	helper.ResponseOK(c, earnings, "current month earnings successfully retrieved", http.StatusOK)

}
func (ctrl *DashboardController) RenevueChartController(c *gin.Context) {
	revenueChart, err := ctrl.Service.Dashboard.GenerateRenevueChart()
	if err != nil {
		ctrl.Log.Error("Error generating revenue chart", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Error generating revenue chart", http.StatusInternalServerError)
		return
	}
	c.Header("Content-Type", "text/html")
	c.Writer.Write(revenueChart.Bytes())
}
func (ctrl *DashboardController) GetBestProductListController(c *gin.Context) {
	bestProduct, err := ctrl.Service.Dashboard.GetBestItemList()
	if err != nil {
		ctrl.Log.Error("Error getting best product list", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Error getting best product list", http.StatusInternalServerError)
		return
	}
	helper.ResponseOK(c, bestProduct, "best product list successfully retrieved", http.StatusOK)
}
