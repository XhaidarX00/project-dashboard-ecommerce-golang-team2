package routes

import (
	"dashboard-ecommerce-team2/infra"

	"github.com/gin-gonic/gin"
	// swaggerFiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRoutes(ctx infra.ServiceContext) *gin.Engine {
	r := gin.Default()

	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	banner := r.Group("/api")
	{
		banner.GET("/banner", ctx.Ctl.Banner.GetBannerByIDController)
		banner.DELETE("/banner", ctx.Ctl.Banner.DeleteBannerController)
		banner.PUT("/banner", ctx.Ctl.Banner.UpdateBannerController)
		banner.POST("/create-banner", ctx.Ctl.Banner.CreateBannerController)
	}
	return r
}
