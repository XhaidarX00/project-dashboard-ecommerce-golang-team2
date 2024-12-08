package productrepository_test

import (
	"dashboard-ecommerce-team2/models"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)
type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) Create(productInput *models.Product) error {
	args := m.Called(productInput)
	return args.Error(0)
}

func (m *MockProductRepository) Update(id int, productInput models.Product) (*models.Product, error) {
	args := m.Called(id, productInput)
	if product := args.Get(0); product != nil {
		return product.(*models.Product), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockProductRepository) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockProductRepository) GetByID(id int) (*models.ProductID, error) {
	args := m.Called(id)
	if product := args.Get(0); product != nil {
		return product.(*models.ProductID), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockProductRepository) GetAll(page, pageSize int) ([]*models.ProductWithCategory, int64, error) {
	args := m.Called(page, pageSize)
	if products := args.Get(0); products != nil {
		return products.([]*models.ProductWithCategory), args.Get(1).(int64), args.Error(2)
	}
	return nil, args.Get(1).(int64), args.Error(2)
}

func (m *MockProductRepository) CountProduct() (int, error) {
	args := m.Called()
	return args.Get(0).(int), args.Error(1)
}

func (m *MockProductRepository) CountEachProduct() ([]models.BestProduct, error) {
	args := m.Called()
	if products := args.Get(0); products != nil {
		return products.([]models.BestProduct), args.Error(1)
	}
	return nil, args.Error(1)
}
func TestGetByID_Success(t *testing.T) {
	// Buat mock repository
	mockRepo := new(MockProductRepository)

	// Data mock untuk produk
	expectedProduct := &models.ProductID{
		ID:          14,
		Name:        "Mock Product",
		CodeProduct: "MP123",
		Stock:       100,
		Price:       200.0,
		CategoryName: "Mock Category",
	}

	// Set expectation pada mock
	mockRepo.On("GetByID", 14).Return(expectedProduct, nil)


	// Panggil metode GetByID
	result, err := mockRepo.GetByID(14)

	// Validasi hasil
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedProduct.ID, result.ID)
	assert.Equal(t, expectedProduct.Name, result.Name)

	// Validasi apakah metode mock dipanggil
	mockRepo.AssertCalled(t, "GetByID", 14)
}

func TestGetByID_Error(t *testing.T) {
	// Buat mock repository
	mockRepo := new(MockProductRepository)
	
	// Set expectation pada mock bahwa metode GetByID akan dipanggil dan mengembalikan error
	mockRepo.On("GetByID", 999).Return(nil, errors.New("product not found"))

	// Panggil metode GetByID
	result, err := mockRepo.GetByID(999)

	// Validasi hasil
	assert.Error(t, err)           // Pastikan ada error
	assert.Nil(t, result)          // Pastikan result nil
	assert.Equal(t, "product not found", err.Error()) // Pastikan error sesuai dengan yang diharapkan

	// Validasi apakah metode mock dipanggil dengan parameter yang benar
	mockRepo.AssertCalled(t, "GetByID", 999)
}


func TestCreate_Success(t *testing.T) {
	// Buat mock repository
	mockRepo := new(MockProductRepository)

	// Data input yang valid
	validProduct := &models.Product{
		CategoryID:  1,
		Name:        "Smartphone",
		CodeProduct: "SPH-001",
		Images:      []string{"/images/smartphone1.png"},
		Description: "Latest smartphone with advanced features",
		Stock:       50,
		Price:       699.99,
		Published:   true,
	}

	// Set expectation pada mock
	mockRepo.On("Create", validProduct).Return(nil)

	// Panggil metode Create
	err := mockRepo.Create(validProduct)

	// Validasi hasil
	assert.NoError(t, err)

	// Validasi apakah metode mock dipanggil
	mockRepo.AssertCalled(t, "Create", validProduct)
}


func TestCreate_Error_DBError(t *testing.T) {
	// Buat mock repository
	mockRepo := new(MockProductRepository)

	// Siapkan data produk yang valid
	validProduct := models.Product{
		CategoryID:  1,
		Name:        "Valid Product",
		CodeProduct: "VP-001",
		Stock:       10,
		Price:       99.99,
		Published:   true,
	}

	// Simulasikan error pada database
	mockRepo.On("Create", &validProduct).Return(fmt.Errorf("database error"))

	// Panggil metode Create
	err := mockRepo.Create(&validProduct)

	// Validasi bahwa error terjadi
	assert.Error(t, err)
	assert.Equal(t, "database error", err.Error())

	// Validasi apakah metode mock dipanggil dengan parameter yang benar
	mockRepo.AssertCalled(t, "Create", &validProduct)
}

func TestDelete_Success(t *testing.T) {
	// Buat mock repository
	mockRepo := new(MockProductRepository)

	// Data produk yang valid
	productID := 1

	// Set expectation pada mock bahwa metode Delete akan dipanggil dan tidak ada error
	mockRepo.On("Delete", productID).Return(nil)

	// Panggil metode Delete
	err := mockRepo.Delete(productID)

	// Validasi hasil
	assert.NoError(t, err)

	// Validasi apakah metode mock dipanggil
	mockRepo.AssertCalled(t, "Delete", productID)
}

func TestDelete_ProductNotFound(t *testing.T) {
	// Buat mock repository
	mockRepo := new(MockProductRepository)

	// Data produk yang tidak ada
	productID := 999 // Produk dengan ID ini tidak ada

	// Set expectation pada mock bahwa metode Delete akan dipanggil dan mengembalikan error
	mockRepo.On("Delete", productID).Return(fmt.Errorf("product not found"))

	// Panggil metode Delete
	err := mockRepo.Delete(productID)

	// Validasi hasil
	assert.Error(t, err)
	assert.Equal(t, "product not found", err.Error())

	// Validasi apakah metode mock dipanggil
	mockRepo.AssertCalled(t, "Delete", productID)
}

func TestDelete_Failure(t *testing.T) {
	// Buat mock repository
	mockRepo := new(MockProductRepository)

	// Data produk yang valid
	productID := 1

	// Set expectation pada mock bahwa metode Delete akan dipanggil dan mengembalikan error lainnya
	mockRepo.On("Delete", productID).Return(fmt.Errorf("failed to delete product"))

	// Panggil metode Delete
	err := mockRepo.Delete(productID)

	// Validasi hasil
	assert.Error(t, err)
	assert.Equal(t, "failed to delete product", err.Error())

	// Validasi apakah metode mock dipanggil
	mockRepo.AssertCalled(t, "Delete", productID)
}


func TestGetAll_Success(t *testing.T) {
	// Buat mock repository
	mockRepo := new(MockProductRepository)

	// Data produk yang valid
	products := []*models.ProductWithCategory{
		{
			ID:           1,
			CategoryName: "Electronics",
			Name:         "Smartphone",
			CodeProduct:  "SPH-001",
			Stock:        50,
			Price:        699.99,
			Published:    true,
		},
		{
			ID:           2,
			CategoryName: "Electronics",
			Name:         "Laptop",
			CodeProduct:  "LPT-002",
			Stock:        30,
			Price:        999.99,
			Published:    true,
		},
	}

	// Total item yang ditemukan
	totalItems := int64(2)

	// Set expectation pada mock bahwa metode GetAll akan dipanggil dan mengembalikan produk dan total item
	mockRepo.On("GetAll", 1, 10).Return(products, totalItems, nil)

	// Panggil metode GetAll
	result, total, err := mockRepo.GetAll(1, 10)

	// Validasi hasil
	assert.NoError(t, err)
	assert.Equal(t, totalItems, total)
	assert.Len(t, result, 2)
	assert.Equal(t, result[0].CategoryName, "Electronics")
	assert.Equal(t, result[1].Name, "Laptop")

	// Validasi apakah metode mock dipanggil
	mockRepo.AssertCalled(t, "GetAll", 1, 10)
}

func TestGetAll_Error(t *testing.T) {
	// Buat mock repository
	mockRepo := new(MockProductRepository)

	// Set expectation pada mock bahwa metode GetAll akan dipanggil dan mengembalikan error
	mockRepo.On("GetAll", 1, 10).Return(nil, int64(0), fmt.Errorf("failed to fetch products"))

	// Panggil metode GetAll
	result, total, err := mockRepo.GetAll(1, 10)

	// Validasi hasil
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, int64(0), total)

	// Validasi apakah metode mock dipanggil
	mockRepo.AssertCalled(t, "GetAll", 1, 10)
}

func TestUpdate_Success(t *testing.T) {
	// Buat mock repository
	mockRepo := new(MockProductRepository)

	// Siapkan data produk yang akan di-update
	productToUpdate := models.Product{
		Name:        "Updated Smartphone",
		CodeProduct: "SPH-002",
		Images:      []string{"/images/updated-smartphone.png"},
		Description: "Updated smartphone with new features",
		Stock:       30,
		Price:       799.99,
		Published:   true,
	}

	// Siapkan produk yang diharapkan setelah update
	updatedProduct := &models.Product{
		ID:          1,
		Name:        "Updated Smartphone",
		CodeProduct: "SPH-002",
		Images:      []string{"/images/updated-smartphone.png"},
		Description: "Updated smartphone with new features",
		Stock:       30,
		Price:       799.99,
		Published:   true,
	}

	// Set expectation pada mock bahwa metode Update akan dipanggil dan mengembalikan produk yang sudah di-update
	mockRepo.On("Update", 1, productToUpdate).Return(updatedProduct, nil)

	// Panggil metode Update
	result, err := mockRepo.Update(1, productToUpdate)

	// Validasi hasil
	assert.NoError(t, err)            // Pastikan tidak ada error
	assert.NotNil(t, result)          // Pastikan result tidak nil
	assert.Equal(t, updatedProduct, result)  // Pastikan hasilnya sama dengan produk yang diharapkan

	// Validasi apakah metode mock dipanggil dengan parameter yang benar
	mockRepo.AssertCalled(t, "Update", 1, productToUpdate)
}

func TestUpdate_Error(t *testing.T) {
	// Buat mock repository
	mockRepo := new(MockProductRepository)

	// Siapkan data produk yang akan di-update
	productToUpdate := models.Product{
		Name:        "Non-existent Product",
		CodeProduct: "NEP-001",
		Images:      []string{"/images/non-existent.png"},
		Description: "This product does not exist",
		Stock:       0,
		Price:       0.00,
		Published:   false,
	}

	// Set expectation pada mock bahwa metode Update akan dipanggil dan mengembalikan error
	mockRepo.On("Update", 999, productToUpdate).Return(nil, fmt.Errorf("product not found"))

	// Panggil metode Update
	result, err := mockRepo.Update(999, productToUpdate)

	// Validasi hasil
	assert.Error(t, err)           // Pastikan ada error
	assert.Nil(t, result)          // Pastikan result nil
	assert.Equal(t, "product not found", err.Error()) // Pastikan error sesuai dengan yang diharapkan

	// Validasi apakah metode mock dipanggil dengan parameter yang benar
	mockRepo.AssertCalled(t, "Update", 999, productToUpdate)
}

func TestCountProduct_Success(t *testing.T) {
	// Setup mock database dan repository
	mockRepo := new(MockProductRepository)
	
	// Siapkan mock data
	mockRepo.On("CountProduct").Return(100, nil)

	// Panggil metode CountProduct
	count, err := mockRepo.CountProduct()

	// Validasi bahwa tidak ada error dan hasilnya sesuai
	assert.NoError(t, err)
	assert.Equal(t, 100, count)

	// Validasi bahwa metode mock dipanggil dengan benar
	mockRepo.AssertExpectations(t)
}

func TestCountProduct_Error(t *testing.T) {
	// Setup mock database dan repository
	mockRepo := new(MockProductRepository)
	
	// Simulasikan error pada query
	mockRepo.On("CountProduct").Return(0, fmt.Errorf("database error"))

	// Panggil metode CountProduct
	count, err := mockRepo.CountProduct()

	// Validasi bahwa ada error dan hasilnya sesuai
	assert.Error(t, err)
	assert.Equal(t, "database error", err.Error())
	assert.Equal(t, 0, count)

	// Validasi bahwa metode mock dipanggil dengan benar
	mockRepo.AssertExpectations(t)
}

func TestCountEachProduct_Success(t *testing.T) {
	// Setup mock database dan repository
	mockRepo := new(MockProductRepository)
	
	// Siapkan mock data yang akan dikembalikan oleh CountEachProduct
	mockProducts := []models.BestProduct{
		{ProductID: 1, Name: "Product 1", Total: 100},
		{ProductID: 2, Name: "Product 2", Total: 50},
	}

	// Simulasikan CountEachProduct yang sukses
	mockRepo.On("CountEachProduct").Return(mockProducts, nil)

	// Panggil metode CountEachProduct
	products, err := mockRepo.CountEachProduct()

	// Validasi bahwa tidak ada error dan hasilnya sesuai
	assert.NoError(t, err)
	assert.Len(t, products, 2)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, 100, products[0].Total)

	// Validasi bahwa metode mock dipanggil dengan benar
	mockRepo.AssertExpectations(t)
}

func TestCountEachProduct_Error(t *testing.T) {
	// Setup mock database dan repository
	mockRepo := new(MockProductRepository)

	// Simulasikan error pada CountEachProduct
	mockRepo.On("CountEachProduct").Return(nil, fmt.Errorf("database error"))

	// Panggil metode CountEachProduct
	products, err := mockRepo.CountEachProduct()

	// Validasi bahwa ada error dan hasilnya sesuai
	assert.Error(t, err)
	assert.Equal(t, "database error", err.Error())
	assert.Nil(t, products)

	// Validasi bahwa metode mock dipanggil dengan benar
	mockRepo.AssertExpectations(t)
}

