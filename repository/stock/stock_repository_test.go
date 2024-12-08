package stockrepository_test

import (
	"dashboard-ecommerce-team2/repository/stock"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupTestDB() (*gorm.DB, sqlmock.Sqlmock) {

	db, mock, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       db,
		DriverName: "postgres",
	})
	gormDB, _ := gorm.Open(dialector, &gorm.Config{})
	return gormDB, mock
}

func TestDelete(t *testing.T) {

	logger, _ := zap.NewDevelopment()
	db, mock := SetupTestDB()

	repo := stockrepository.NewStockRepository(db, logger)

	mock.ExpectQuery(`SELECT \* FROM "stocks" WHERE "stocks"."id" = \$1 ORDER BY "stocks"."id" LIMIT \$2`).
		WithArgs(1, 1).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	mock.ExpectQuery(`SELECT \* FROM "stocks" WHERE "stocks"."id" = \$1 ORDER BY "stocks"."id" LIMIT \$2`).
		WithArgs(999, 1).
		WillReturnError(gorm.ErrRecordNotFound)

	mock.ExpectQuery(`SELECT \* FROM "stocks" WHERE "stocks"."id" = \$1 ORDER BY "stocks"."id" LIMIT 1`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectExec(`DELETE FROM "stocks" WHERE "stocks"."id" = \$1`).
		WithArgs(1).
		WillReturnError(errors.New("some error"))
	err := repo.Delete(1)
	assert.Error(t, err)
}
func TestGetByID(t *testing.T) {

	logger, _ := zap.NewDevelopment()
	db, mock := SetupTestDB()

	repo := stockrepository.NewStockRepository(db, logger)

	mock.ExpectQuery(`SELECT stocks.id, stocks.product_id, stocks.type, stocks.quantity, stocks.created_at, stocks.updated_at, products.name as product_name, CASE WHEN categories.variant = '{}' THEN NULL ELSE categories.variant END AS variant FROM "stocks" JOIN products ON products.id = stocks.product_id JOIN categories ON categories.id = products.category_id WHERE stocks.id = \$1 ORDER BY "stocks"."id" LIMIT \$2`).
		WithArgs(1, 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "product_id", "type", "quantity", "product_name", "variant"}).
			AddRow(1, 1, "in", 100, "Item A", nil))

	stock, err := repo.GetByID(1)
	assert.NoError(t, err)
	assert.Equal(t, 1, int(stock.ID))
	assert.Equal(t, "Item A", stock.ProductName)
	assert.Equal(t, 100, stock.Quantity)

	mock.ExpectQuery(`SELECT stocks.id, stocks.product_id, stocks.type, stocks.quantity, stocks.created_at, stocks.updated_at, products.name as product_name, CASE WHEN categories.variant = '{}' THEN NULL ELSE categories.variant END AS variant FROM "stocks" JOIN products ON products.id = stocks.product_id JOIN categories ON categories.id = products.category_id WHERE stocks.id = \$1 ORDER BY "stocks"."id" LIMIT \$2`).
		WithArgs(999, 1).
		WillReturnError(gorm.ErrRecordNotFound)

	stock, err = repo.GetByID(999)
	assert.Error(t, err)
	assert.Nil(t, stock)
}

// func TestStockRepository_Update_Postgres(t *testing.T) {

// 	logger, _ := zap.NewDevelopment()
// 	db, mock := SetupTestDB()

// 	repo := stockrepository.NewStockRepository(db, logger)

// 	tests := []struct {
// 		name        string
// 		input       *models.StockRequest
// 		setupMocks  func()
// 		expectedErr error
// 	}{
// 		{
// 			name: "Successful stock update",
// 			input: &models.StockRequest{
// 				ProductID: 1,
// 				Type:      "in",
// 				Quantity:  10,
// 			},
// 			setupMocks: func() {

// 				mock.ExpectQuery(`SELECT stock FROM products WHERE id = \?`).
// 					WithArgs(1).
// 					WillReturnRows(sqlmock.NewRows([]string{"stock"}).AddRow(5))

// 				mock.ExpectBegin()

// 				mock.ExpectExec(`UPDATE products SET stock = \?, updated_at = \? WHERE id = \?`).
// 					WithArgs(15, sqlmock.AnyArg(), 1).
// 					WillReturnResult(sqlmock.NewResult(1, 1))

// 				mock.ExpectExec(`INSERT INTO stocks \(product_id, type, quantity, created_at\)`).
// 					WithArgs(1, "in", 10, sqlmock.AnyArg()).
// 					WillReturnResult(sqlmock.NewResult(1, 1))

// 				mock.ExpectCommit()
// 			},
// 			expectedErr: nil,
// 		},
// 		{
// 			name: "Insufficient stock on out",
// 			input: &models.StockRequest{
// 				ProductID: 1,
// 				Type:      "out",
// 				Quantity:  10,
// 			},
// 			setupMocks: func() {

// 				mock.ExpectQuery(`SELECT stock FROM products WHERE id = \?`).
// 					WithArgs(1).
// 					WillReturnRows(sqlmock.NewRows([]string{"stock"}).AddRow(5))
// 			},
// 			expectedErr: fmt.Errorf("insufficient stock"),
// 		},
// 		{
// 			name: "Invalid stock type",
// 			input: &models.StockRequest{
// 				ProductID: 1,
// 				Type:      "invalid",
// 				Quantity:  10,
// 			},
// 			setupMocks:  func() {},
// 			expectedErr: fmt.Errorf("invalid stock type"),
// 		},
// 		{
// 			name: "Failed to update product stock",
// 			input: &models.StockRequest{
// 				ProductID: 1,
// 				Type:      "in",
// 				Quantity:  10,
// 			},
// 			setupMocks: func() {

// 				mock.ExpectQuery(`SELECT stock FROM products WHERE id = \?`).
// 					WithArgs(1).
// 					WillReturnRows(sqlmock.NewRows([]string{"stock"}).AddRow(5))

// 				mock.ExpectBegin()

// 				mock.ExpectExec(`UPDATE products SET stock = \?, updated_at = \? WHERE id = \?`).
// 					WithArgs(15, sqlmock.AnyArg(), 1).
// 					WillReturnError(fmt.Errorf("failed to update product stock"))

// 				mock.ExpectRollback()
// 			},
// 			expectedErr: fmt.Errorf("failed to update product stock: failed to update product stock"),
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.setupMocks()

// 			err := repo.Update(tt.input)

// 			if tt.expectedErr != nil {
// 				assert.Error(t, err)
// 				assert.Equal(t, tt.expectedErr.Error(), err.Error())
// 			} else {
// 				assert.NoError(t, err)
// 			}

// 			if err := mock.ExpectationsWereMet(); err != nil {
// 				t.Errorf("There were unfulfilled expectations: %v", err)
// 			}
// 		})
// 	}
// }
