package ordercontroller

import (
	"dashboard-ecommerce-team2/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type OrderController struct {
	Service service.Service
	Log     *zap.Logger
}

func NewOrderController(service service.Service, log *zap.Logger) *OrderController {
	return &OrderController{
		Service: service,
		Log:     log,
	}
}

func (ctrl *OrderController) UpdateOrderStatusController(c *gin.Context) {}
func (ctrl *OrderController) GetAllOrdersController(c *gin.Context)      {}
func (ctrl *OrderController) GetOrderByIDController(c *gin.Context)      {}
