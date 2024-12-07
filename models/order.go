package models

import "time"

type Order struct {
	ID              int       `json:"id" gorm:"primaryKey" swaggerignore:"true"`
	UserID          int       `json:"user_id" gorm:"not null" binding:"required" example:"1"`
	TotalAmount     float64   `json:"total_amount" gorm:"type:decimal(10,2);not null" binding:"required" example:"150.75"`
	PaymentMethod   string    `json:"payment_method" gorm:"size:20;not null" binding:"required" example:"credit_card"`
	ShippingAddress string    `json:"shipping_address" gorm:"size:255" binding:"omitempty" example:"123 Main St"`
	Status          string    `json:"status" gorm:"size:20;check:status IN ('pending','shipped','completed','canceled');default:'created'" binding:"required" example:"pending"`
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime" swaggerignore:"true"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"autoUpdateTime" swaggerignore:"true"`
}

type OrderItem struct {
	ID        int     `json:"id" gorm:"primaryKey" swaggerignore:"true"`
	OrderID   int     `json:"order_id" gorm:"not null" binding:"required" example:"1"`
	ProductID int     `json:"product_id" gorm:"not null" binding:"required" example:"101"`
	Quantity  int     `json:"quantity" gorm:"not null" binding:"required" example:"2"`
	Price     float64 `json:"price" gorm:"type:decimal(10,2);not null" binding:"required" example:"75.00"`
	Total     float64 `json:"total" gorm:"type:decimal(10,2);not null" binding:"required"`
}

func OrderSeed() []Order {
	return []Order{
		{
			ID:              1,
			UserID:          1,
			TotalAmount:     150.75,
			PaymentMethod:   "credit_card",
			ShippingAddress: "123 Main St",
			Status:          "pending",
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		},
		{
			ID:              2,
			UserID:          2,
			TotalAmount:     300.00,
			PaymentMethod:   "paypal",
			ShippingAddress: "456 Elm St",
			Status:          "shipped",
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		},
		{
			ID:              3,
			UserID:          3,
			TotalAmount:     500.50,
			PaymentMethod:   "debit_card",
			ShippingAddress: "789 Oak St",
			Status:          "completed",
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		},
	}
}

func OrderItemSeed() []OrderItem {
	return []OrderItem{
		{
			ID:        1,
			OrderID:   1,
			ProductID: 101,
			Quantity:  2,
			Price:     75.00,
			Total:     150.00,
		},
		{
			ID:        2,
			OrderID:   2,
			ProductID: 102,
			Quantity:  3,
			Price:     100.00,
			Total:     300.00,
		},
		{
			ID:        3,
			OrderID:   3,
			ProductID: 103,
			Quantity:  5,
			Price:     100.10,
			Total:     500.50,
		},
	}
}
