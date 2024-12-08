package stockrepository_test

// import (
// 	"dashboard-ecommerce-team2/models"
// 	stockrepository "dashboard-ecommerce-team2/repository/stock"
// 	"fmt"
// 	"testing"
// 	"time"

// 	"github.com/DATA-DOG/go-sqlmock"
// 	"github.com/stretchr/testify/assert"
// 	"go.uber.org/zap"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/logger"
// )

// func setupTestDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock, func()) {
// 	mockDB, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("Failed to create mock database: %v", err)
// 	}

// 	dialector := postgres.New(postgres.Config{
// 		Conn: mockDB,
// 	})
// 	db, err := gorm.Open(dialector, &gorm.Config{
// 		Logger: logger.Default.LogMode(logger.Silent), // Menonaktifkan logging di pengujian
// 	})
// 	if err != nil {
// 		t.Fatalf("Failed to open gorm DB: %v", err)
// 	}

// 	// Mengambil objek sql.DB dan membuat fungsi untuk menutup koneksi.
// 	sqlDB, err := db.DB()
// 	if err != nil {
// 		t.Fatalf("Failed to get underlying sql.DB: %v", err)
// 	}

// 	// Fungsi untuk menutup koneksi DB
// 	cleanup := func() {
// 		sqlDB.Close()
// 	}

// 	return db, mock, cleanup
// }

// func TestDelete(t *testing.T) {
// 	// Mock DB and Logger
// 	db, mock, err := sqlmock.New()
// 	assert.NoError(t, err)
// 	defer db.Close()

// 	gormDB, err := gorm.Open(postgres.New(postgres.Config{
// 		Conn: db,
// 	}), &gorm.Config{})
// 	assert.NoError(t, err)

// 	logger, _ := zap.NewProduction()

// 	repo := stockrepository.NewStockRepository(gormDB, logger)

// 	t.Run("successfully delete stock history", func(t *testing.T) {
// 		mock.ExpectQuery("SELECT * FROM \"stocks\" WHERE \"stocks\".\"id\" = \\$1 ORDER BY \"stocks\".\"id\" LIMIT \\$2").
// 			WithArgs(1, 1).
// 			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

// 		mock.ExpectExec("DELETE FROM \"stocks\" WHERE \"stocks\".\"id\" = \\$1").
// 			WithArgs(1).
// 			WillReturnResult(sqlmock.NewResult(1, 1))

// 		err := repo.Delete(1)
// 		assert.NoError(t, err)
// 	})

// 	t.Run("error when stock history not found", func(t *testing.T) {
// 		mock.ExpectQuery("SELECT * FROM \"stocks\" WHERE \"stocks\".\"id\" = \\$1 ORDER BY \"stocks\".\"id\" LIMIT \\$2").
// 			WithArgs(1, 1).
// 			WillReturnError(fmt.Errorf("record not found"))

// 		err := repo.Delete(1)
// 		assert.EqualError(t, err, "history not found")
// 	})

// 	t.Run("error on delete failure", func(t *testing.T) {
// 		mock.ExpectQuery("SELECT * FROM \"stocks\" WHERE \"stocks\".\"id\" = \\$1 ORDER BY \"stocks\".\"id\" LIMIT \\$2").
// 			WithArgs(1, 1).
// 			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

// 		mock.ExpectExec("DELETE FROM \"stocks\" WHERE \"stocks\".\"id\" = \\$1").
// 			WithArgs(1).
// 			WillReturnError(fmt.Errorf("delete error"))

// 		err := repo.Delete(1)
// 		assert.EqualError(t, err, "delete error")
// 	})
// }

// func TestGetByID(t *testing.T) {
// 	// Mock DB and Logger
// 	db, mock, err := sqlmock.New()
// 	assert.NoError(t, err)
// 	defer db.Close()

// 	gormDB, err := gorm.Open(postgres.New(postgres.Config{
// 		Conn: db,
// 	}), &gorm.Config{})
// 	assert.NoError(t, err)

// 	logger, _ := zap.NewProduction()

// 	repo := stockrepository.NewStockRepository(gormDB, logger)

// 	t.Run("successfully fetch stock by ID", func(t *testing.T) {
// 		mock.ExpectQuery("SELECT stocks.id, stocks.product_id, stocks.type, stocks.quantity, stocks.created_at, stocks.updated_at, products.name as product_name, CASE WHEN categories.variant = '{}' THEN NULL ELSE categories.variant END AS variant FROM \"stocks\" JOIN products ON products.id = stocks.product_id JOIN categories ON categories.id = products.category_id WHERE stocks.id = \\$1").
// 			WithArgs(1).
// 			WillReturnRows(sqlmock.NewRows([]string{"id", "product_id", "type", "quantity", "created_at", "updated_at", "product_name", "variant"}).
// 				AddRow(1, 1, "in", 10, time.Now(), time.Now(), "Product Name", `{"size": "L"}`))

// 		result, err := repo.GetByID(1)
// 		assert.NoError(t, err)
// 		assert.NotNil(t, result) // Pastikan result tidak nil

// 		expectedVariant := map[string]string{"size": "L"}
// 		assert.Equal(t, expectedVariant, result.Variant)
// 	})

// 	t.Run("error when stock history not found", func(t *testing.T) {
// 		mock.ExpectQuery("SELECT stocks.id, stocks.product_id, stocks.type, stocks.quantity, stocks.created_at, stocks.updated_at, products.name as product_name, CASE WHEN categories.variant = '{}' THEN NULL ELSE categories.variant END AS variant FROM \"stocks\" JOIN products ON products.id = stocks.product_id JOIN categories ON categories.id = products.category_id WHERE stocks.id = \\$1").
// 			WithArgs(1).
// 			WillReturnError(fmt.Errorf("record not found"))

// 		result, err := repo.GetByID(1)
// 		assert.Error(t, err)
// 		assert.Nil(t, result)
// 	})

// 	t.Run("error on fetch failure", func(t *testing.T) {
// 		mock.ExpectQuery("SELECT stocks.id, stocks.product_id, stocks.type, stocks.quantity, stocks.created_at, stocks.updated_at, products.name as product_name, CASE WHEN categories.variant = '{}' THEN NULL ELSE categories.variant END AS variant FROM \"stocks\" JOIN products ON products.id = stocks.product_id JOIN categories ON categories.id = products.category_id WHERE stocks.id = \\$1").
// 			WithArgs(1).
// 			WillReturnError(fmt.Errorf("fetch error"))

// 		result, err := repo.GetByID(1)
// 		assert.Error(t, err)
// 		assert.Nil(t, result)
// 	})
// }

// func TestUpdate(t *testing.T) {
// 	db, mock, cleanup := setupTestDB(t)
// 	defer cleanup()

// 	logger := zap.NewNop()
// 	repo := stockrepository.NewStockRepository(db, logger)

// 	t.Run("successfully update stock", func(t *testing.T) {
// 		mock.ExpectQuery("SELECT stocks.id, stocks.product_id, stocks.type, stocks.quantity, stocks.created_at, stocks.updated_at, products.name as product_name, CASE WHEN categories.variant = '{}' THEN NULL ELSE categories.variant END AS variant FROM \"stocks\" JOIN products ON products.id = stocks.product_id JOIN categories ON categories.id = products.category_id WHERE stocks.id = \\$1").
// 			WithArgs(1).
// 			WillReturnRows(sqlmock.NewRows([]string{"id", "product_id", "type", "quantity", "created_at", "updated_at", "product_name", "variant"}).
// 				AddRow(1, 1, "in", 10, time.Now(), time.Now(), "Product Name", `{"size": "L"}`))

// 		mock.ExpectExec("UPDATE \"products\" SET .* WHERE id = \\$1").
// 			WithArgs(1).
// 			WillReturnResult(sqlmock.NewResult(1, 1))

// 		mock.ExpectExec("INSERT INTO \"stocks\" .*").
// 			WillReturnResult(sqlmock.NewResult(1, 1))

// 		err := repo.Update(&models.StockRequest{
// 			ProductID: 1,
// 			Type:      "in",
// 			Quantity:  10,
// 		})
// 		assert.NoError(t, err)
// 	})

// 	t.Run("error when invalid stock type", func(t *testing.T) {
// 		err := repo.Update(&models.StockRequest{
// 			ProductID: 1,
// 			Type:      "invalid",
// 			Quantity:  10,
// 		})
// 		assert.Error(t, err)
// 		assert.Equal(t, "invalid stock type", err.Error())
// 	})
// }
