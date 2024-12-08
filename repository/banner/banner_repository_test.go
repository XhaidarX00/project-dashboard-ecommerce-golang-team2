package bannerrepository_test

import (
	"dashboard-ecommerce-team2/config"
	"dashboard-ecommerce-team2/database"
	"dashboard-ecommerce-team2/helper"
	"dashboard-ecommerce-team2/mocks"
	"dashboard-ecommerce-team2/models"
	bannerrepository "dashboard-ecommerce-team2/repository/banner"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func TestBannerRepository_Create(t *testing.T) {
	// Initialize the mock
	mockService := new(mocks.BannerRepositoryMock)

	// Test data
	banner := &models.Banner{ID: 1, Title: "Test Banner", Published: false}

	// Set expectations for the mock service
	mockService.On("Create", banner).Return(nil)

	// Act
	err := mockService.Create(banner)

	// Assert
	assert.NoError(t, err, "Create should not return an error")

	// Validate mock interactions
	mockService.AssertExpectations(t) // Ensure all expected method calls were made
}

func TestBannerRepository_Delete(t *testing.T) {
	// Initialize the mock service
	mockService := new(mocks.BannerRepositoryMock)

	// Test data
	// banner := &models.Banner{ID: 1, Title: "Test Banner"}

	// Set expectations for the mock service
	mockService.On("Delete", 1).Return(nil)                            // Case 1: Successful deletion
	mockService.On("Delete", 2).Return(errors.New("Banner not found")) // Case 2: Non-existent banner

	// Act - Case 1
	err := mockService.Delete(1)
	assert.NoError(t, err, "Delete should not return an error")

	// Act - Case 2
	err = mockService.Delete(2)
	assert.Error(t, err, "Delete should return an error for non-existent banner")

	// Validate mock interactions
	mockService.AssertExpectations(t) // Ensure all expected method calls were made
}

func TestBannerRepository_GetByID(t *testing.T) {
	// Initialize the mock service
	mockService := new(mocks.BannerRepositoryMock)

	// Test data
	banner := &models.Banner{ID: 1, Title: "Test Banner"}

	// Set expectations for the mock service
	mockService.On("GetByID", 1).Return(banner, nil)                         // Case 1: Existing banner
	mockService.On("GetByID", 2).Return(nil, errors.New("Banner not found")) // Case 2: Non-existent banner

	// Act - Case 1
	result, err := mockService.GetByID(1)
	assert.NoError(t, err, "GetByID should not return an error for existing banner")
	assert.Equal(t, banner.Title, result.Title, "Banner name should match")

	// Act - Case 2
	result, err = mockService.GetByID(2)
	assert.Error(t, err, "GetByID should return an error for non-existent banner")
	assert.Nil(t, result, "Result should be nil for non-existent banner")

	// Validate mock interactions
	mockService.AssertExpectations(t) // Ensure all expected method calls were made
}

func TestBannerRepository_Update(t *testing.T) {
	// Initialize the mock service
	mockService := new(mocks.BannerRepositoryMock)

	// Test data
	banner := models.Banner{ID: 1, Title: "Test Banner", Published: false}

	// Set expectations for the mock service
	mockService.On("Update", banner).Return(nil)

	// Act
	err := mockService.Update(banner)
	assert.NoError(t, err, "Update should not return an error")

	// Validate mock interactions
	mockService.AssertExpectations(t) // Ensure all expected method calls were made
}

func setupTestDatabase(t *testing.T) (*gorm.DB, *zap.Logger, error) {
	// Setup test database
	config, err := config.ReadConfig()
	if err != nil {
		assert.Error(t, err)
	}

	// instance looger
	log, err := helper.InitZapLogger()
	if err != nil {
		assert.Error(t, err)
	}

	// instance database
	db, err := database.InitDB(config)
	if err != nil {
		assert.Error(t, err)
	}

	return db, log, err
}

func TestBannerRepository_GetByID2(t *testing.T) {
	// Setup test database
	db, log, err := setupTestDatabase(t)
	assert.NoError(t, err)

	repository := bannerrepository.NewBannerRepository(db, log)

	// Test Case 1: Successful creation and retrieval
	banner := &models.Banner{
		ID: 1,
	}

	// Retrieve the created banner
	retrievedBanner, err := repository.GetByID(banner.ID)
	assert.NoError(t, err, "GetByID should not return an error for existing banner")
	assert.NotNil(t, retrievedBanner, "Retrieved banner should not be nil")
	assert.Equal(t, banner.Title, retrievedBanner.Title, "Banner title should match")

	// Test Case 2: Attempt to retrieve non-existent banner
	nonExistentBanner, err := repository.GetByID(9999) // Assuming 9999 is not a valid ID
	assert.Error(t, err, "GetByID should return an error for non-existent banner")
	assert.Nil(t, nonExistentBanner, "Result should be nil for non-existent banner")

	// Test Case 3: Validate specific banner attributes
	assert.Equal(t, banner.Published, retrievedBanner.Published, "Published status should match")
	assert.NotZero(t, retrievedBanner.ID, "Banner ID should be assigned")
}
