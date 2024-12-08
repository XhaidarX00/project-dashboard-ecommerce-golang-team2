package bannercontroller

import (
	"dashboard-ecommerce-team2/helper"
	"dashboard-ecommerce-team2/models"
	"dashboard-ecommerce-team2/service"
	"log"
	"net/http"
	"strconv"

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
		helper.ResponseError(c, "INVALID", "invalid data input", http.StatusBadRequest)
		c.Abort()
		return
	}

	// Process image upload
	imageURL, err := helper.UploadImage(banner.ImagePath)
	if err != nil {
		log.Println(err.Error())
		helper.ResponseError(c, "FAILED", "Failed to upload image", http.StatusInternalServerError)
		return
	}

	// Assign the image URL to the banner
	var bannerSave models.Banner
	bannerSave, err = bannerSave.CopyBannerGetValueToBanner(imageURL, banner)
	if err != nil {
		log.Println("Error binding form data:", err.Error())
		helper.ResponseError(c, "INVALID", "invalid data input", http.StatusBadRequest)
		c.Abort()
		return
	}

	// Call service to create the banner
	if err := ctrl.Service.Banner.CreateBanner(&bannerSave); err != nil {
		helper.ResponseError(c, "FAILED", "Failed to create banner", http.StatusInternalServerError)
		c.Abort()
		return
	}

	helper.ResponseOK(c, bannerSave, "Successfully created banner", http.StatusOK)
}

func (ctrl *BannerController) UpdateBannerController(c *gin.Context) {
	id := c.Request.FormValue("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.ResponseError(c, "FAILED", "id banner invalid", http.StatusBadRequest)
		c.Abort()
		return
	}

	banner, err := ctrl.Service.Banner.GetBannerByID(idInt)
	if err != nil {
		helper.ResponseError(c, "FAILED", err.Error(), http.StatusInternalServerError)
		c.Abort()
		return
	}

	// change status publis
	banner.Published = !banner.Published
	err = ctrl.Service.Banner.UpdateBanner(*banner)
	if err != nil {
		helper.ResponseError(c, "FAILED", err.Error(), http.StatusInternalServerError)
		c.Abort()
		return
	}

	helper.ResponseOK(c, banner, "Successfully update published banner", http.StatusOK)
}

func (ctrl *BannerController) GetBannerByIDController(c *gin.Context) {
	id := c.Request.FormValue("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.ResponseError(c, "FAILED", "id banner invalid", http.StatusBadRequest)
		c.Abort()
		return
	}

	banner, err := ctrl.Service.Banner.GetBannerByID(idInt)
	if err != nil {
		helper.ResponseError(c, "FAILED", err.Error(), http.StatusInternalServerError)
		c.Abort()
		return
	}

	helper.ResponseOK(c, banner, "Successfully get banner", http.StatusOK)
}

func (ctrl *BannerController) DeleteBannerController(c *gin.Context) {
	id := c.Request.FormValue("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.ResponseError(c, "FAILED", "id banner invalid", http.StatusBadRequest)
		c.Abort()
		return
	}

	err = ctrl.Service.Banner.DeleteBanner(idInt)
	if err != nil {
		helper.ResponseError(c, "FAILED", err.Error(), http.StatusInternalServerError)
		c.Abort()
		return
	}

	helper.ResponseOK(c, nil, "Successfully delete banner", http.StatusOK)
}
