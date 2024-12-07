package bannercontroller

import (
	"dashboard-ecommerce-team2/helper"
	"dashboard-ecommerce-team2/models"
	"dashboard-ecommerce-team2/service"
	"fmt"
	"log"
	"net/http"

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

func (ctrl *BannerController) CreateBannerController(c *gin.Context) {
	var banner models.BannerGetValue

	// Bind the form data to the banner struct (without the file)
	if err := c.ShouldBind(&banner); err != nil {
		log.Println("Error binding form data:", err.Error())
		helper.ResponseError(c, "INVALID", "invalid data input", http.StatusBadRequest)
		c.Abort()
		return
	}

	// Process image upload
	imageURL, err := helper.UploadImage(banner.ImagePath)
	if err != nil {
		helper.ResponseError(c, "FAILED", "Failed to upload image", http.StatusInternalServerError)
		return
	}

	// Assign the image URL to the banner
	var bannerSave models.Banner
	bannerSave = bannerSave.CopyBannerGetValueToBanner(imageURL, banner)

	fmt.Printf("Banner: %+v\n", bannerSave)

	// Call service to create the banner
	if err := ctrl.Service.Banner.CreateBanner(&bannerSave); err != nil {
		helper.ResponseError(c, "FAILED", "Failed to create banner", http.StatusInternalServerError)
		c.Abort()
		return
	}

	// Respond with the created banner data
	helper.ResponseOK(c, bannerSave, "Successfully created banner", http.StatusOK)
}

func (ctrl *BannerController) UpdateBannerController(c *gin.Context)  {}
func (ctrl *BannerController) GetBannerByIDController(c *gin.Context) {}
func (ctrl *BannerController) DeleteBannerController(c *gin.Context)  {}
