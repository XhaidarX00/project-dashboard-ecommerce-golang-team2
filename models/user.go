package models

import "time"

type User struct {
	ID        int       `json:"id" gorm:"primaryKey" swaggerignore:"true"`
	Name      string    `json:"name" gorm:"size:100;not null" binding:"required" example:"John Doe"`
	Email     string    `json:"email" gorm:"size:100;unique" binding:"required,email" example:"johndoe@example.com"`
	Password  string    `json:"-" gorm:"size:255;not null" binding:"required,min=8" swaggerignore:"true"`
	Role      string    `json:"role" gorm:"size:20;check:role IN ('admin','customer', 'staff')" binding:"required,oneof=admin customer staff" example:"admin"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime" swaggerignore:"true"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" example:"johndoe@example.com"`
	Password string `json:"password" binding:"required" example:"password123"`
}

func UserSeed() []User {
	return []User{}
}
