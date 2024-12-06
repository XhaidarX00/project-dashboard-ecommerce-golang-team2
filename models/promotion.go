package models

import (
	"time"

	"gorm.io/gorm"
)

type Promotion struct {
	ID          int             `json:"id" gorm:"primaryKey" swaggerignore:"true"`
	Name        string          `json:"name" gorm:"size:255;not null" binding:"required" example:"Holiday Sale"`
	ProductName JSONB           `json:"product_name" gorm:"type:jsonb;default:'[]'" binding:"omitempty" example:"[\"Smartphone\", \"Laptop\"]"`
	Type        JSONB           `json:"type" gorm:"type:jsonb;default:'[]'" binding:"omitempty" example:"[\"discount\", \"bundle\"]"`
	Description string          `json:"description" binding:"omitempty" example:"Special holiday discounts"`
	Discount    JSONB           `json:"discount" gorm:"type:jsonb;default:'[]'" binding:"omitempty" example:"[{\"value\":10,\"type\":\"percentage\"}]"`
	StartDate   time.Time       `json:"start_date" binding:"required" example:"2024-12-01"`
	EndDate     time.Time       `json:"end_date" binding:"required,gtfield=StartDate" example:"2024-12-31"`
	Quota       int             `json:"quota" gorm:"default:0" example:"100"`
	Status      bool            `json:"status" gorm:"default:false" example:"true"`
	Published   bool            `json:"published" gorm:"default:false" example:"true"`
	CreatedAt   time.Time       `json:"created_at" gorm:"autoCreateTime" swaggerignore:"true"`
	UpdatedAt   time.Time       `json:"updated_at" gorm:"autoUpdateTime" swaggerignore:"true"`
	DeletedAt   *gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty" swaggerignore:"true"`
}

func PromotionSeed() []Promotion {
	return []Promotion{}
}
