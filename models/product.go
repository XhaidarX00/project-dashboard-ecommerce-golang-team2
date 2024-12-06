package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          int             `json:"id" gorm:"primaryKey" swaggerignore:"true"`
	CategoryID  int             `json:"category_id" gorm:"not null" binding:"required" example:"1"`
	Name        string          `json:"name" gorm:"size:100;not null" binding:"required" example:"Smartphone"`
	CodeProduct string          `json:"code_product" gorm:"size:100;unique;not null" binding:"required" example:"SPH-001"`
	Images      JSONB           `json:"images" gorm:"type:jsonb;default:'[]'" binding:"omitempty" example:"[\"/images/smartphone1.png\"]"`
	Description string          `json:"description" binding:"omitempty" example:"Latest smartphone with advanced features"`
	Stock       int             `json:"stock" gorm:"not null" binding:"required" example:"50"`
	Price       float64         `json:"price" gorm:"type:decimal(10,2);not null" binding:"required" example:"699.99"`
	Published   bool            `json:"published" gorm:"default:false" example:"true"`
	CreatedAt   time.Time       `json:"created_at" gorm:"autoCreateTime" swaggerignore:"true"`
	UpdatedAt   time.Time       `json:"updated_at" gorm:"autoUpdateTime" swaggerignore:"true"`
	DeletedAt   *gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty" swaggerignore:"true"`
}

type StockHistory struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ProductID uint      `gorm:"not null" json:"product_id"`
	Type      string    `gorm:"size:10;check:type IN ('in','out')" json:"type"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type BestProduct struct {
	Name  string `json:"name"`
	Total string `json:"total"`
}

func ProductSeed() []Product {
	return []Product{}
}

func StockHistorySeed() []StockHistory {
	return []StockHistory{}
}
