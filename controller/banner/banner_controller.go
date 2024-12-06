package bannercontroller

import (
	"dashboard-ecommerce-team2/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BannerController struct {
	Service service.Service
	Log     *zap.Logger
}

func NewBannerController(service service.Service, log *zap.Logger) *BannerController {
	return &BannerController{
		Service: service,
		Log:     log,
	}
}

func (ctrl *BannerController) CreateBannerController(c *gin.Context)  {}
func (ctrl *BannerController) UpdateBannerController(c *gin.Context)  {}
func (ctrl *BannerController) GetBannerByIDController(c *gin.Context) {}
func (ctrl *BannerController) DeleteBannerController(c *gin.Context)  {}
