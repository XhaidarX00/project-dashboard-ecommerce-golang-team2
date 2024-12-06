package stockcontroller

import (
	"dashboard-ecommerce-team2/service"

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

func (ctrl *StockController) UpdateProductStockController(c *gin.Context)    {}
func (ctrl *StockController) GetProductStockDetailController(c *gin.Context) {}
func (ctrl *StockController) DeleteProductStockController(c *gin.Context)    {}
