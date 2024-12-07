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

	router := r.Group("/products")
	{
		router.POST("/", ctx.Ctl.Product.CreateProductController)
	}

	router := r.Group("/orders")
	{
		router.GET("/", ctx.Ctl.Order.GetAllOrdersController)
		router.GET("/:id", ctx.Ctl.Order.GetOrderByIDController)
		router.PUT("/update/:id", ctx.Ctl.Order.UpdateOrderStatusController)
		router.DELETE("/:id", ctx.Ctl.Order.DeleteOrderController)
		router.GET("/detail/:id", ctx.Ctl.Order.GetOrderDetailController)
	}

	return r
}

func authRoutes(r *gin.Engine, ctx infra.ServiceContext) {
	authGroup := r.Group("/auth")
	authGroup.GET("/login", ctx.Ctl.User.LoginController)
	authGroup.GET("/check-email", ctx.Ctl.User.CheckEmailUserController)
	authGroup.POST("/register", ctx.Ctl.User.CreateUserController)
	authGroup.PATCH("/reset-password", ctx.Ctl.User.ResetUserPasswordController)
}
