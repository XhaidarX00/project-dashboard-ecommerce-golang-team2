package userservice

import (
	"dashboard-ecommerce-team2/models"
	"dashboard-ecommerce-team2/repository"

	"go.uber.org/zap"
)

type UserService interface {
	CreateUser(userInput models.User) error
	Login(input models.LoginRequest) error
	CheckUserEmail(email string) error
	ResetUserPassword(password string) error
}

type userService struct {
	Repo repository.Repository
	Log  *zap.Logger
}

// CheckUserEmail implements UserService.
func (u *userService) CheckUserEmail(email string) error {
	panic("unimplemented")
}

// CreateUser implements UserService.
func (u *userService) CreateUser(userInput models.User) error {
	panic("unimplemented")
}

// Login implements UserService.
func (u *userService) Login(input models.LoginRequest) error {
	panic("unimplemented")
}

// ResetUserPassword implements UserService.
func (u *userService) ResetUserPassword(password string) error {
	panic("unimplemented")
}

func NewUserService(repo repository.Repository, log *zap.Logger) UserService {
	return &userService{Repo: repo, Log: log}
}
