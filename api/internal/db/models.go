package db

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Firstname string    `gorm:"not null" json:"firstname"`
	Lastname  string    `gorm:"not null" json:"lastname"`
	Username  string    `gorm:"unique;not null" json:"username"`
	Email     string    `gorm:"unique;not null" json:"email"`
	Password  string    `gorm:"not null" json:"password"`

	IsAdmin    bool `gorm:"default:false" json:"isAdmin"`
	IsVerified bool `gorm:"default:false" json:"isVerified"`

  Projects []Project `gorm:"type:uuid" json:"projects"`

  CreatedAt time.Time `gorm:"autoCreateTime;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Project struct {
	gorm.Model
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Description string    `gorm:"not null" json:"description"`

  AllowedUrls []string `json:"allowedUrls"`

  UserID uuid.UUID `gorm:"type:uuid;not null" json:"userId"`
  User   User      `gorm:"constraint:OnDelete:CASCADE;not null" json:"user"`

	CreatedAt time.Time `gorm:"autoCreateTime;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
