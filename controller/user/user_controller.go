package usercontroller

import (
	"dashboard-ecommerce-team2/database"
	"dashboard-ecommerce-team2/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserController struct {
	Service service.Service
	Log     *zap.Logger
	Cacher  database.Cacher
}

func NewUserController(service service.Service, log *zap.Logger, cacher database.Cacher) *UserController {
	return &UserController{
		Service: service,
		Log:     log,
		Cacher:  cacher,
	}
}

func (ctrl *UserController) CreateUserController(c *gin.Context)        {}
func (ctrl *UserController) LoginController(c *gin.Context)             {}
func (ctrl *UserController) CheckEmailUserController(c *gin.Context)    {}
func (ctrl *UserController) ResetUserPasswordController(c *gin.Context) {}
