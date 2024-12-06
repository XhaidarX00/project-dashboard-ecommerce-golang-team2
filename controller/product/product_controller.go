package productcontroller

import (
	"dashboard-ecommerce-team2/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ProductController struct {
	Service service.Service
	Log     *zap.Logger
}

func NewProductController(service service.Service, log *zap.Logger) *ProductController {
	return &ProductController{
		Service: service,
		Log:     log,
	}
}

func (ctrl *ProductController) CreateProductController(c *gin.Context)  {}
func (ctrl *ProductController) GetAllProductsController(c *gin.Context) {}
func (ctrl *ProductController) GetProductByIDController(c *gin.Context) {}
func (ctrl *ProductController) UpdateProductController(c *gin.Context)  {}
func (ctrl *ProductController) DeleteProductController(c *gin.Context)  {}
