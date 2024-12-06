package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"
)

type Banner struct {
	ID          int             `json:"id" gorm:"primaryKey" swaggerignore:"true"`
	Image       string          `json:"image" gorm:"size:255;not null" binding:"required" example:"/images/banner1.png"`
	Title       string          `json:"title" gorm:"size:255;not null" binding:"required" example:"Winter Sale"`
	Type        JSONB           `json:"type" gorm:"type:jsonb;default:'[]'" binding:"omitempty" example:"[\"seasonal\", \"promo\"]"`
	PathPage    string          `json:"path_page" gorm:"size:255;not null" binding:"required" example:"/sale"`
	ReleaseDate time.Time       `json:"release_date" binding:"omitempty" example:"2024-01-15"`
	EndDate     time.Time       `json:"end_date" binding:"omitempty" example:"2024-01-31"`
	Published   bool            `json:"published" gorm:"default:false" example:"true"`
	CreatedAt   time.Time       `json:"created_at" gorm:"autoCreateTime" swaggerignore:"true"`
	UpdatedAt   time.Time       `json:"updated_at" gorm:"autoUpdateTime" swaggerignore:"true"`
	DeletedAt   *gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty" swaggerignore:"true"`
}

type JSONB []interface{}

// Value Marshal
func (a JSONB) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan Unmarshal
func (a *JSONB) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &a)
}

func BannerSeed() []Banner {
	return []Banner{}
}
