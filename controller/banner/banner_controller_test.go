package bannercontroller_test

// import (
// 	"bytes"
// 	"mime/multipart"
// 	"net/http"
// 	"net/http/httptest"
// 	"os"
// 	"testing"
// 	"time"

// 	bannercontroller "dashboard-ecommerce-team2/controller/banner"
// 	"dashboard-ecommerce-team2/mocks"
// 	"dashboard-ecommerce-team2/models"
// 	"dashboard-ecommerce-team2/service"

// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// 	"go.uber.org/zap"
// )

// func setupRouter(mockService service.Service) *gin.Engine {
// 	gin.SetMode(gin.TestMode)
// 	router := gin.New()
// 	ctrl := bannercontroller.NewBannerController(mockService, zap.NewNop())
// 	router.POST("/banners", ctrl.CreateBannerController)
// 	router.PUT("/banners", ctrl.UpdateBannerController)
// 	router.GET("/banners/:id", ctrl.GetBannerByIDController)
// 	router.DELETE("/banners/:id", ctrl.DeleteBannerController)
// 	return router
// }

// func createMultipartFormFile(t *testing.T, fieldName, fileName string) (*multipart.FileHeader, *os.File) {
// 	file, err := os.Create(fileName)
// 	if err != nil {
// 		t.Fatalf("failed to create temporary file: %v", err)
// 	}
// 	defer file.Close()

// 	// Menulis data ke file
// 	_, err = file.Write([]byte("test image data"))
// 	if err != nil {
// 		t.Fatalf("failed to write to temporary file: %v", err)
// 	}

// 	// Membuat file multipart
// 	body := &bytes.Buffer{}
// 	writer := multipart.NewWriter(body)

// 	part, err := writer.CreateFormFile(fieldName, fileName)
// 	if err != nil {
// 		t.Fatalf("failed to create form file: %v", err)
// 	}
// 	part.Write([]byte("test image data"))

// 	writer.Close()

// 	fileHeader := &multipart.FileHeader{
// 		Filename: fileName,
// 		Size:     int64(len("test image data")),
// 		Header:   make(map[string][]string),
// 	}
// 	return fileHeader, file // Kembalikan fileHeader dan file
// }

// func TestCreateBannerController_Success(t *testing.T) {
// 	mockService := new(mocks.Service)
// 	Service := new(service.Service)
// 	router := setupRouter(*Service)

// 	// Membuat file multipart
// 	imageFileHeader, file := createMultipartFormFile(t, "image_path", "test_image.jpg")
// 	defer os.Remove(file.Name()) // Menghapus file setelah pengujian
// 	now := time.Now()
// 	banner := models.BannerGetValue{
// 		Title:       "New Banner",
// 		ImagePath:   imageFileHeader,
// 		Type:        []string{"seasonal", "promo"},
// 		PathPage:    "/spring-sale",
// 		ReleaseDate: &now, // Gantilah dengan waktu yang sesuai
// 		EndDate:     &now, // Gantilah dengan waktu yang sesuai
// 		Published:   true,
// 	}

// 	mockService.Repo.On("Create", mock.Anything).Return(banner)

// 	// Membuat request
// 	body := &bytes.Buffer{}
// 	writer := multipart.NewWriter(body)
// 	writer.WriteField("title", banner.Title)
// 	writer.WriteField("type", `["seasonal", "promo"]`)
// 	writer.WriteField("path_page", banner.PathPage)
// 	writer.WriteField("release_date", "2024-03-01")
// 	writer.WriteField("end_date", "2024-03-31")
// 	writer.CreateFormFile("image_path", imageFileHeader.Filename)
// 	writer.Close()

// 	req, _ := http.NewRequest(http.MethodPost, "/banners", body)
// 	req.Header.Set("Content-Type", writer.FormDataContentType())
// 	w := httptest.NewRecorder()

// 	// Mengirimkan request
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)
// 	mockService.Repo.AssertExpectations(t)
// }

// func TestCreateBannerController_InvalidData(t *testing.T) {
// 	mockService := new(mocks.Service)
// 	router := setupRouter(mockService)

// 	body := bytes.NewBufferString(`{"title":""}`) // Invalid data
// 	req, _ := http.NewRequest(http.MethodPost, "/banners", body)
// 	req.Header.Set("Content-Type", "application/json")
// 	w := httptest.NewRecorder()

// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusBadRequest, w.Code)
// }

// func TestUpdateBannerController_Success(t *testing.T) {
// mockService := new(mocks.Service)
// router := setupRouter(mockService)

// banner := &models.Banner{
// ID:        1,
// Title:     "Updated Banner",
// Published: true,
// }

// 	mockService.Banner.On("GetBannerByID", 1).Return(banner, nil)
// 	mockService.Banner.On("UpdateBanner", *banner).Return(nil)

// 	req, _ := http.NewRequest(http.MethodPut, "/banners?id=1", nil)
// 	w := httptest.NewRecorder()

// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)
// 	mockService.Banner.AssertExpectations(t)
// }

// func TestGetBannerByIDController_Success(t *testing.T) {
// 	mockService := new(mocks.Service)
// 	Service := new(service.Service)
// 	router := setupRouter(*Service) // Gunakan mockService di sini

// 	banner := &models.Banner{
// 		ID:    1,
// 		Title: "Banner",
// 	}

// 	// Set up the mock expectation for getting banner by ID
// 	mockService.Repo.On("GetByID", 1).Return(banner, nil) // Ganti sesuai dengan metode yang benar
// 	req, _ := http.NewRequest(http.MethodGet, "/banners?id=1", nil)
// 	w := httptest.NewRecorder()

// 	// Send the request to the router
// 	router.ServeHTTP(w, req)

// 	// Assertions
// 	assert.Equal(t, http.StatusOK, w.Code)
// 	mockService.Repo.AssertExpectations(t) // Pastikan untuk memanggil AssertExpectations pada mockService.Banner
// }

// func TestDeleteBannerController_Success(t *testing.T) {
// 	mockService := new(mocks.Service)
// 	router := setupRouter(mockService)

// 	mockService.Banner.On("DeleteBanner", 1).Return(nil)

// 	req, _ := http.NewRequest(http.MethodDelete, "/banners?id=1", nil)
// 	w := httptest.NewRecorder()

// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)
// 	mockService.Banner.AssertExpectations(t)
// }

// func TestDeleteBannerController_InvalidID(t *testing.T) {
// 	mockService := new(mocks.Service)
// 	router := setupRouter(mockService)

// 	req, _ := http.NewRequest(http.MethodDelete, "/banners?id=invalid", nil)
// 	w := httptest.NewRecorder()

// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusBadRequest, w.Code)
// }
