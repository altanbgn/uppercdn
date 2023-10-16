package models

import (
	"gorm.io/gorm"
)

type App struct {
  gorm.Model
  UUID string `json:"uuid" gorm:"column:uuid primary_key unique"`
  Name string `json:"name" gorm:"column:name"`
  Description string `json:"description" gorm:"column:description"`
  APIKey []APIKey `json:"api_key" gorm:"foreignKey:AppUUID;references:UUID"`
}

type APIKey struct {
	gorm.Model
  AppUUID    string    `json:"app_uuid" gorm:"column:app_uuid"`
	Key        string    `json:"key" gorm:"column:key"`
	ExpiryDate string    `json:"expiry_date" gorm:"column:expiry_date"`
}
