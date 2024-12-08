package stockcontroller

import (
	"dashboard-ecommerce-team2/helper"
	"dashboard-ecommerce-team2/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type StockController struct {
	Service service.Service
	Log     *zap.Logger
}

func NewStockController(service service.Service, log *zap.Logger) *StockController {
	return &StockController{
		Service: service,
		Log:     log,
	}
}

func (ctrl *StockController) UpdateProductStockController(c *gin.Context) {}
func (ctrl *StockController) GetProductStockDetailController(c *gin.Context) {
	id := helper.StringToInt(c.Param("id"))
	stockHistory, err := ctrl.Service.Stock.GetProductStockDetail(id)
	if err != nil {
		ctrl.Log.Warn("handler: No history found", zap.Int("id", id))
		helper.ResponseError(c, "No history found", err.Error(), http.StatusNotFound)
		return
	}

	helper.ResponseOK(c, stockHistory, "", http.StatusOK)
}
func (ctrl *StockController) DeleteProductStockController(c *gin.Context) {}
