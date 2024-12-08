package mocks

import (
	"dashboard-ecommerce-team2/models"

	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

// BannerRepositoryMock is a mock of the BannerRepository interface
type BannerRepositoryMock struct {
	mock.Mock
}

// Create is a mock method for creating a banner
func (m *BannerRepositoryMock) Create(banner *models.Banner) error {
	args := m.Called(banner)
	return args.Error(0)
}

// Update is a mock method for updating a banner
func (m *BannerRepositoryMock) Update(banner models.Banner) error {
	args := m.Called(banner)
	return args.Error(0)
}

// Delete is a mock method for deleting a banner by ID
func (m *BannerRepositoryMock) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

// GetByID is a mock method for retrieving a banner by ID
func (m *BannerRepositoryMock) GetByID(id int) (*models.Banner, error) {
	args := m.Called(id)
	if banner := args.Get(0); banner != nil {
		return banner.(*models.Banner), args.Error(1)
	}
	return nil, args.Error(1)
}

// Service adalah mock dari interface Service
type Service struct {
	Repo   *BannerRepositoryMock
	Logger *zap.Logger
}

// NewService membuat instance baru dari mock Service
func NewService(bnr *BannerRepositoryMock, log *zap.Logger) *Service {
	return &Service{
		Repo:   bnr,
		Logger: log,
	}
}
