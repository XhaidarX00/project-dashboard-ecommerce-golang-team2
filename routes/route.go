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
