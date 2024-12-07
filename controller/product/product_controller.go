package productcontroller

import (
	"dashboard-ecommerce-team2/helper"
	"dashboard-ecommerce-team2/models"
	"dashboard-ecommerce-team2/service"
	"net/http"

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

func (ctrl *ProductController) CreateProductController(c *gin.Context) {
	var input models.Product

	if err := c.ShouldBind(&input); err != nil {
		ctrl.Log.Error("handler: Failed to bind input", zap.Error(err))
		helper.ResponseError(c, "invalid payload", helper.FormatValidationError(err), http.StatusBadRequest)
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		ctrl.Log.Error("handler: Failed to get uploaded file", zap.Error(err))
		helper.ResponseError(c, "image is required", err.Error(), http.StatusBadRequest)
		return
	}

	filePath := "/tmp/" + file.Filename
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		ctrl.Log.Error("handler: Failed to save uploaded file", zap.Error(err))
		helper.ResponseError(c, "Failed to save file", err.Error(), http.StatusInternalServerError)
		return
	}
	ctrl.Log.Info("File saved successfully", zap.String("file_path", filePath))

	product, err := ctrl.Service.Product.CreateProduct(&input, filePath)
	if err != nil {
		ctrl.Log.Error("handler: Failed to create product", zap.Error(err))
		helper.ResponseError(c, "Failed to create product", err.Error(), http.StatusBadRequest)
		return
	}

	ctrl.Log.Info("handler: Product created successfully", zap.Int("product_id", product.ID))
	helper.ResponseOK(c, product, "create successfully", http.StatusCreated)

}
func (ctrl *ProductController) GetAllProductsController(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "10")

	pageInt := helper.StringToInt(page)
	pageSizeInt := helper.StringToInt(pageSize)

	products, totalItems, err := ctrl.Service.Product.GetAllProducts(pageInt, pageSizeInt)
	if err != nil {
		ctrl.Log.Error("Failed to fetch products", zap.Error(err))
		helper.ResponseError(c, "Failed to fetch products", err.Error(), http.StatusInternalServerError)
		return
	}
	totalPages := (totalItems + pageSizeInt - 1) / pageSizeInt

	helper.ResponseOKPagination(c, products, "", pageInt, pageSizeInt, totalItems, totalPages, http.StatusOK)
}
func (ctrl *ProductController) GetProductByIDController(c *gin.Context) {}
func (ctrl *ProductController) UpdateProductController(c *gin.Context)  {}
func (ctrl *ProductController) DeleteProductController(c *gin.Context)  {}
