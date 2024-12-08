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

	// authMiddleware := ctx.Middleware.Authentication()
	adminMiddleware := ctx.Middleware.RoleAuthorization("admin")

	productRoutes := r.Group("/products")
	{
		productRoutes.POST("/", ctx.Ctl.Product.CreateProductController)
		productRoutes.GET("/", ctx.Ctl.Product.GetAllProductsController)
		productRoutes.GET("/:id", ctx.Ctl.Product.GetProductByIDController)
		productRoutes.DELETE("/:id", adminMiddleware, ctx.Ctl.Product.DeleteProductController)
		productRoutes.PUT("/:id", ctx.Ctl.Product.UpdateProductController)
	}

	return r
}
