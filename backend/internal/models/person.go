package models

import "time"

type Person struct {
	ID         uint      `gorm:"primaryKey"`
	FirstName  string    `json:"first_name" gorm:"size:50"`
	MiddleName *string   `json:"middle_name" gorm:"size:50"`
	LastName   string    `json:"last_name" gorm:"size:50"`
	Birthdate  time.Time `json:"birthdate"`
	Sex        string    `json:"sex" gorm:"foreignKey:SexID"`
	Addresses  []Address `json:"adresses" gorm:"foreignKey:PersonID"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Address struct {
	ID           uint    `gorm:"primaryKey"`
	PersonID     uint    `gorm:"not null;index"`
	Person       Person  `gorm:"foreignKey:PersonID"`
	StreetName   string  `json:"street_name" gorm:"size:100"`
	StreetNumber uint    `json:"street_number"`
	StateID      uint    `json:"state_id"`
	State        State   `gorm:"foreignKey:StateID"`
	CountryID    uint    `json:"country_id"`
	Country      Country `gorm:"foreignKey:CountryID"`
	ZipCode      string  `json:"zip_code" gorm:"size:20"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
}

type State struct {
	ID      uint   `gorm:"primaryKey"`
	Name    string `gorm:"size:50"`
	Acronym string `gorm:"size:10"`
}

type Country struct {
	ID      uint   `gorm:"primaryKey"`
	Name    string `gorm:"size:50"`
	Acronym string `gorm:"size:10"`
}
