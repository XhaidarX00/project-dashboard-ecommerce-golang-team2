package bannerservice

import (
	"dashboard-ecommerce-team2/mocks"
	"dashboard-ecommerce-team2/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func NewBannerServiceMock(repo *mocks.BannerRepositoryMock, log *zap.Logger) *mocks.Service {
	return &mocks.Service{
		Repo:   repo,
		Logger: log,
	}
}

// Create a new BannerService with the mock repository
func createMockService() (*mocks.Service, *mocks.BannerRepositoryMock) {
	mockRepo := new(mocks.BannerRepositoryMock)
	mockLogger := zap.NewNop()
	service := NewBannerServiceMock(mockRepo, mockLogger)
	return service, mockRepo
}

// Unit tests for BannerService

func TestCreateBanner(t *testing.T) {
	service, mockRepo := createMockService()
	banner := &models.Banner{
		ID:    1,
		Title: "New Banner",
	}

	// Setting up the mock expectation
	mockRepo.On("Create", banner).Return(nil)

	// Call the service method
	err := service.Repo.Create(banner)

	// Assertions
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateBanner(t *testing.T) {
	service, mockRepo := createMockService()
	banner := models.Banner{
		ID:    1,
		Title: "Updated Banner",
	}

	// Setting up the mock expectation
	mockRepo.On("Update", banner).Return(nil)

	// Call the service method
	err := service.Repo.Update(banner)

	// Assertions
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetBannerByID(t *testing.T) {
	service, mockRepo := createMockService()
	banner := &models.Banner{
		ID:    1,
		Title: "Banner",
	}

	// Setting up the mock expectation
	mockRepo.On("GetByID", 1).Return(banner, nil)

	// Call the service method
	result, err := service.Repo.GetByID(1)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, banner, result)
	mockRepo.AssertExpectations(t)
}

func TestDeleteBanner(t *testing.T) {
	service, mockRepo := createMockService()

	// Setting up the mock expectation
	mockRepo.On("Delete", 1).Return(nil)

	// Call the service method
	err := service.Repo.Delete(1)

	// Assertions
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
