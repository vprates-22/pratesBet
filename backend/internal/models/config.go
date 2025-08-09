package models

import (
	"time"

	"gorm.io/gorm"
)

type Configuration struct {
	ID            uint     `gorm:"primaryKey"`
	LanguageID    uint     `json:"language_id"`
	Language      Language `gorm:"foreignKey:LanguageID" json:"language"`
	Timezone      string   `gorm:"size:50"`
	Theme         string   `gorm:"size:20"`
	ReceiveEmails bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type Language struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"size:50;not null" json:"name"`
	Code string `gorm:"size:10;not null;unique" json:"code"`
}
