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

	IsAdmin    bool `gorm:"default:false" json:"is_admin"`
	IsVerified bool `gorm:"default:false" json:"is_verified"`

  Projects []Project `gorm:"foreignKey:OwnerID" json:"projects"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type APIKey struct {
	gorm.Model
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
  Key       string    `gorm:"unique"`
  ExpireAt  time.Time `gorm:"not null"`

	ProjectID uuid.UUID `gorm:"type:uuid;not null"`
	Project   Project   `gorm:"foreignKey:ProjectID;not null"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type Project struct {
	gorm.Model
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Description string    `gorm:"not null" json:"description"`

  OwnerID uuid.UUID `gorm:"type:uuid;not null" json:"owner_id"`
  Owner   User      `gorm:"foreignKey:OwnerID" json:"owner"`

  APIKeys []APIKey `json:"api_keys"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
