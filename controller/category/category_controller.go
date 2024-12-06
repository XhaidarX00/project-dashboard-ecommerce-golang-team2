package categorycontroller

import (
	"dashboard-ecommerce-team2/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CategoryController struct {
	Service service.Service
	Log     *zap.Logger
}

func NewCategoryController(service service.Service, log *zap.Logger) *CategoryController {
	return &CategoryController{
		Service: service,
		Log:     log,
	}
}

func (ctrl *CategoryController) CreateCatergoryController(c *gin.Context)  {}
func (ctrl *CategoryController) GetAllCategoriesController(c *gin.Context) {}
func (ctrl *CategoryController) GetCategoryByIDController(c *gin.Context)  {}
func (ctrl *CategoryController) UpdateCategoryController(c *gin.Context)   {}
func (ctrl *CategoryController) DeleteCategoryController(c *gin.Context)   {}
