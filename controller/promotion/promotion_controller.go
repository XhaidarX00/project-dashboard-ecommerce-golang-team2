package promotioncontroller

import (
	"dashboard-ecommerce-team2/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PromotionController struct {
	Service service.Service
	Log     *zap.Logger
}

func NewPromotionController(service service.Service, log *zap.Logger) *PromotionController {
	return &PromotionController{
		Service: service,
		Log:     log,
	}
}

func (ctrl *PromotionController) GetAllPromotionsController(c *gin.Context) {}
func (ctrl *PromotionController) CreatePromotionController(c *gin.Context)  {}
func (ctrl *PromotionController) UpdatePromotionController(c *gin.Context)  {}
func (ctrl *PromotionController) DeletePromotionController(c *gin.Context)  {}
