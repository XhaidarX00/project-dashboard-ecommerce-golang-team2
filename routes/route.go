package routes

import (
	"dashboard-ecommerce-team2/infra"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRoutes(ctx infra.ServiceContext) *gin.Engine {

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	authRoutes(r, ctx)
	dashboardRoutes(r, ctx)

	productRoutes := r.Group("/products")
	{
		productRoutes.POST("/", ctx.Ctl.Product.CreateProductController)
		productRoutes.GET("/", ctx.Ctl.Product.GetAllProductsController)
	}

	orderRoutes := r.Group("/orders")
	{
		orderRoutes.GET("/", ctx.Ctl.Order.GetAllOrdersController)
		orderRoutes.GET("/:id", ctx.Ctl.Order.GetOrderByIDController)
		orderRoutes.PUT("/update/:id", ctx.Ctl.Order.UpdateOrderStatusController)
		orderRoutes.DELETE("/:id", ctx.Ctl.Order.DeleteOrderController)
		orderRoutes.GET("/detail/:id", ctx.Ctl.Order.GetOrderDetailController)
	}

	bannerRoutes(r, ctx)
	promotionRoutes(r, ctx)
	return r
}

func bannerRoutes(r *gin.Engine, ctx infra.ServiceContext) {
	banner := r.Group("/api")
	banner.GET("/banner", ctx.Ctl.Banner.GetBannerByIDController)
	banner.DELETE("/banner", ctx.Ctl.Banner.DeleteBannerController)
	banner.PUT("/banner", ctx.Ctl.Banner.UpdateBannerController)
	banner.POST("/create-banner", ctx.Ctl.Banner.CreateBannerController)
}

func promotionRoutes(r *gin.Engine, ctx infra.ServiceContext) {
	promotion := r.Group("/api")
	promotion.GET("/promotion", ctx.Ctl.Promotion.GetAllPromotionsController)
	promotion.GET("/promotion:id", ctx.Ctl.Promotion.GetByIdPromotionsController)
	promotion.DELETE("/promotion", ctx.Ctl.Promotion.DeletePromotionController)
	promotion.PUT("/promotion", ctx.Ctl.Promotion.UpdatePromotionController)
	promotion.POST("/create-promotion", ctx.Ctl.Promotion.CreatePromotionController)
}

func authRoutes(r *gin.Engine, ctx infra.ServiceContext) {
	authGroup := r.Group("/auth")
	authGroup.GET("/login", ctx.Ctl.User.LoginController)
	authGroup.GET("/check-email", ctx.Ctl.User.CheckEmailUserController)
	authGroup.POST("/register", ctx.Ctl.User.CreateUserController)
	authGroup.PATCH("/reset-password", ctx.Ctl.User.ResetUserPasswordController)
}

func dashboardRoutes(r *gin.Engine, ctx infra.ServiceContext) {
	dashboardGroup := r.Group("/dashboard")
	dashboardGroup.Use(ctx.Middleware.Authentication())
	{
		dashboardGroup.GET("/summary", ctx.Ctl.Dashboard.GetSummaryController)
		dashboardGroup.GET("/current-month-earning", ctx.Ctl.Dashboard.CurrentMonthEarningController)
		dashboardGroup.GET("/revenue-chart", ctx.Ctl.Dashboard.RenevueChartController)
		dashboardGroup.GET("/best-item-list", ctx.Ctl.Dashboard.GetBestProductListController)
	}
}
