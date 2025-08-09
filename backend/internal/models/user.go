package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID              uint          `gorm:"primaryKey"`
	Username        string        `gorm:"size:30;not null;unique" json:"username"`
	Email           string        `gorm:"size:100;not null;unique" json:"email"`
	Password        string        `json:"-" gorm:"not null"`
	Active          bool          `json:"active" gorm:"default:true"`
	LastLogin       time.Time     `json:"last_login" gorm:"index"`
	GroupID         uint          `json:"group_id"`
	Group           Group         `gorm:"foreignKey:GroupID"`
	PersonID        uint          `json:"person_id"`
	Person          Person        `gorm:"foreignKey:PersonID"`
	ConfigurationID uint          `json:"configuration_id"`
	Configuration   Configuration `gorm:"foreignKey:ConfigurationID"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

type Group struct {
	ID          uint         `gorm:"primaryKey"`
	Name        string       `gorm:"size:50;not null;unique"`
	Description *string      `gorm:"size:255"`
	Permissions []Permission `gorm:"many2many:group_permissions"`
}

type Permission struct {
	ID          uint    `gorm:"primaryKey"`
	Code        string  `gorm:"size:100;unique;not null"`
	Description *string `gorm:"size:255"`
}

type GroupPermission struct {
	GroupID      uint
	PermissionID uint
}
