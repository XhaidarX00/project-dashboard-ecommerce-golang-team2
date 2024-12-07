package models

import (
	"dashboard-ecommerce-team2/helper"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"mime/multipart"
	"time"

	"gorm.io/gorm"
)

type BannerGetValue struct {
	ID          int                   `json:"id" form:"id" gorm:"primaryKey" swaggerignore:"true"`
	ImagePath   *multipart.FileHeader `json:"image_path" form:"image_path" binding:"required" swaggerignore:"true"`
	Title       string                `json:"title" form:"title" gorm:"size:255;not null" binding:"required" example:"Spring Sale 2024"`
	Type        []string              `json:"type" form:"type" gorm:"type:jsonb;default:'[]'" binding:"omitempty" example:"[\"seasonal\", \"promo\"]"`
	PathPage    string                `json:"path_page" form:"path_page" gorm:"size:255;not null" binding:"required" example:"/spring-sale"`
	ReleaseDate *time.Time            `json:"release_date" form:"release_date" gorm:"type:date" binding:"omitempty" example:"2024-03-01"`
	EndDate     *time.Time            `json:"end_date" form:"end_date" gorm:"type:date" binding:"omitempty" example:"2024-03-31"`
	Published   bool                  `json:"published" form:"published" gorm:"default:false" example:"true"`
	CreatedAt   time.Time             `json:"created_at" form:"created_at" gorm:"autoCreateTime" swaggerignore:"true"`
	UpdatedAt   time.Time             `json:"updated_at" form:"updated_at" gorm:"autoUpdateTime" swaggerignore:"true"`
	DeletedAt   *gorm.DeletedAt       `json:"deleted_at,omitempty" form:"deleted_at" gorm:"index" swaggerignore:"true"`
}

type Banner struct {
	ID          int             `json:"id" form:"id" gorm:"primaryKey;autoIncrement" swaggerignore:"true"`
	Image       string          `json:"image" form:"image" gorm:"size:255;not null" binding:"required" example:"/images/banner1.png"`
	Title       string          `json:"title" form:"title" gorm:"size:255;not null" binding:"required" example:"Spring Sale 2024"`
	Type        JSONB           `json:"type" form:"type" gorm:"type:jsonb;default:'[]'" binding:"omitempty" example:"[\"seasonal\", \"promo\"]"`
	PathPage    string          `json:"path_page" form:"path_page" gorm:"size:255;not null" binding:"required" example:"/spring-sale"`
	ReleaseDate *time.Time      `json:"release_date" form:"release_date" gorm:"type:date" binding:"omitempty" example:"2024-03-01"`
	EndDate     *time.Time      `json:"end_date" form:"end_date" gorm:"type:date" binding:"omitempty" example:"2024-03-31"`
	Published   bool            `json:"published" form:"published" gorm:"default:false" example:"true"`
	CreatedAt   time.Time       `json:"created_at" form:"created_at" gorm:"autoCreateTime" swaggerignore:"true"`
	UpdatedAt   time.Time       `json:"updated_at" form:"updated_at" gorm:"autoUpdateTime" swaggerignore:"true"`
	DeletedAt   *gorm.DeletedAt `json:"deleted_at,omitempty" form:"deleted_at" gorm:"index" swaggerignore:"true"`
}

func (b *Banner) CopyBannerGetValueToBanner(urlImagae string, getValue BannerGetValue) Banner {
	now := time.Now()
	return Banner{
		ID:          getValue.ID,
		Title:       getValue.Title,
		Image:       urlImagae,
		Type:        JSONB{getValue.Type[0]},
		PathPage:    getValue.PathPage,
		ReleaseDate: getValue.ReleaseDate,
		EndDate:     getValue.EndDate,
		Published:   getValue.Published,
		CreatedAt:   now,
		UpdatedAt:   now,
		DeletedAt:   nil,
	}
}

type JSONB []interface{}

// Value implements the driver.Valuer interface for JSONB to store as JSON.
func (a JSONB) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan implements the sql.Scanner interface for JSONB to retrieve JSON data.
func (a *JSONB) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("failed to cast value to []byte")
	}
	return json.Unmarshal(b, &a)
}

// BannerSeed provides a seed data function for the Banner model.
func BannerSeed() []Banner {
	now := time.Now()

	return []Banner{
		{
			ID:          1,
			Image:       "/images/banner1.png",
			Title:       "Winter Sale",
			Type:        JSONB{"seasonal", "promo"},
			PathPage:    "/sale",
			ReleaseDate: helper.PointerToTime(time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)),
			EndDate:     helper.PointerToTime(time.Date(2024, 1, 31, 0, 0, 0, 0, time.UTC)),
			Published:   true,
			CreatedAt:   now,
			UpdatedAt:   now,
		},
		{
			ID:          2,
			Image:       "/images/banner2.png",
			Title:       "Spring Promo 2000",
			Type:        JSONB{"promo"},
			PathPage:    "/promo",
			ReleaseDate: helper.PointerToTime(time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC)),
			EndDate:     helper.PointerToTime(time.Date(2024, 3, 15, 0, 0, 0, 0, time.UTC)),
			Published:   false,
			CreatedAt:   now,
			UpdatedAt:   now,
		},
		{
			ID:          3,
			Image:       "/images/banner3.png",
			Title:       "Black Friday Deals",
			Type:        JSONB{"discount", "exclusive"},
			PathPage:    "/black-friday",
			ReleaseDate: helper.PointerToTime(time.Date(2024, 11, 25, 0, 0, 0, 0, time.UTC)),
			EndDate:     helper.PointerToTime(time.Date(2024, 11, 29, 0, 0, 0, 0, time.UTC)),
			Published:   true,
			CreatedAt:   now,
			UpdatedAt:   now,
		},
	}
}

// seedBanners is an example for seeding Banner data.
func SeedBanners(tx *gorm.DB) error {
	// Check if banners already exist
	var count int64
	if err := tx.Model(&Banner{}).Count(&count).Error; err != nil {
		return fmt.Errorf("failed to check banners count: %v", err)
	}

	// Skip seeding if data already exists
	if count > 0 {
		fmt.Println("Skipping BannerSeed, data already exists.")
		return nil
	}

	// Insert seed data
	banners := BannerSeed()
	if err := tx.Create(&banners).Error; err != nil {
		return fmt.Errorf("failed to seed banners: %v", err)
	}

	fmt.Println("BannerSeed executed successfully.")
	return nil
}

// Custom unmarshaller
func (b *Banner) UnmarshalJSON(data []byte) error {
	var aux struct {
		ReleaseDate string `json:"release_date"`
		EndDate     string `json:"end_date"`
		// Field lain
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Parsing ReleaseDate
	if aux.ReleaseDate != "" {
		parsedDate, err := time.Parse("2006-01-02", aux.ReleaseDate)
		if err != nil {
			return err
		}
		b.ReleaseDate = &parsedDate
	}

	// Parsing EndDate
	if aux.EndDate != "" {
		parsedDate, err := time.Parse("2006-01-02", aux.EndDate)
		if err != nil {
			return err
		}
		b.EndDate = &parsedDate
	}

	return nil
}
