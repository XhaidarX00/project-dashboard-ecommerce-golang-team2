package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        int             `json:"id" gorm:"primaryKey" swaggerignore:"true"`
	Name      string          `json:"name" gorm:"size:50;unique;not null" binding:"required" example:"Electronics"`
	Variant   JSONB           `json:"variant" gorm:"type:jsonb;default:'{}'" example:"{\"color\":\"red\", \"size\":\"medium\"}"`
	CreatedAt time.Time       `json:"created_at" gorm:"autoCreateTime" swaggerignore:"true"`
	UpdatedAt time.Time       `json:"updated_at" gorm:"autoUpdateTime" swaggerignore:"true"`
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty" swaggerignore:"true"`
}

func CategorySeed() []Category {
	return []Category{}
}
