package orderrepository

import (
	"dashboard-ecommerce-team2/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type OrderRepository interface {
	UpdateStatus(id int, status string) error
	GetByID(id int) (*models.Order, error)
	GetAll(page, limit int) ([]models.Order, int64, error)
	CountOrder() (int, error)
	CountTotalPriceOrder() (int, error)
	GetEarningEachMonth() (int, error)
	GetDetail(id int) (*models.Order, []models.OrderItem, error)
	DeleteOrder(id int) error
}

type orderRepository struct {
	DB  *gorm.DB
	Log *zap.Logger
}

// GetEarningEachMonth implements OrderRepository.
func (o *orderRepository) GetEarningEachMonth() (int, error) {
	panic("unimplemented")
}

// CountTotalPriceOrder implements OrderRepository.
func (o *orderRepository) CountTotalPriceOrder() (int, error) {
	panic("unimplemented")
}

// CountOrder implements OrderRepository.
func (o *orderRepository) CountOrder() (int, error) {
	panic("unimplemented")
}

// GetAll implements OrderRepository.
func (o *orderRepository) GetAll(page, limit int) ([]models.Order, int64, error) {
	var orders []models.Order
	var totalItems int64

	offset := (page - 1) * limit

	if err := o.DB.Model(&models.Order{}).Count(&totalItems).Error; err != nil {
		o.Log.Error("Failed to count orders", zap.Error(err))
		return nil, 0, err
	}

	if err := o.DB.Limit(limit).Offset(offset).Find(&orders).Error; err != nil {
		o.Log.Error("Failed to retrieve paginated orders", zap.Error(err))
		return nil, 0, err
	}

	return orders, totalItems, nil
}

// GetByID implements OrderRepository.
func (o *orderRepository) GetByID(id int) (*models.Order, error) {
	order := models.Order{}
	if err := o.DB.First(&order, id).Error; err != nil {
		o.Log.Error("Failed to find order", zap.Error(err))
		return nil, err
	}
	return &order, nil
}

// UpdateStatus implements OrderRepository.
func (o *orderRepository) UpdateStatus(id int, status string) error {
	order := models.Order{}
	if err := o.DB.First(&order, id).Error; err != nil {
		o.Log.Error("Failed to find order", zap.Error(err))
		return err
	}
	order.Status = status
	return o.DB.Save(&order).Error
}

func (o *orderRepository) DeleteOrder(id int) error {
	var orderItems []models.OrderItem
	if err := o.DB.Where("order_id = ?", id).Delete(&orderItems).Error; err != nil {
		o.Log.Error("Failed to delete order items", zap.Error(err))
		return err
	}

	if err := o.DB.Delete(&models.Order{}, id).Error; err != nil {
		o.Log.Error("Failed to delete order", zap.Error(err))
		return err
	}
	return nil
}

func (o *orderRepository) GetDetail(id int) (*models.Order, []models.OrderItem, error) {
	var order models.Order
	if err := o.DB.First(&order, id).Error; err != nil {
		o.Log.Error("Failed to find order", zap.Error(err))
		return nil, nil, err
	}

	var orderItems []models.OrderItem
	if err := o.DB.Where("order_id = ?", id).Find(&orderItems).Error; err != nil {
		o.Log.Error("Failed to retrieve order items", zap.Error(err))
		return nil, nil, err
	}

	return &order, orderItems, nil
}

func NewOrderRepository(db *gorm.DB, log *zap.Logger) OrderRepository {
	return &orderRepository{DB: db, Log: log}
}
