package models

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type Filter struct {
	gorm.Model
	SearchQuery           *string
	PriceRange            *Range `gorm:"embedded;embeddedPrefix:price_"`
	RentPriceRange        *Range `gorm:"embedded;embeddedPrefix:rent_price_"`
	ForRent               bool
	City                  *string
	Neighborhood          *string
	SizeRange             *Range `gorm:"embedded;embeddedPrefix:size_"`
	BedroomRange          *Range `gorm:"embedded;embeddedPrefix:bedroom_"`
	FloorRange            *Range `gorm:"embedded;embeddedPrefix:floor_"`
	HasElevator           *bool
	HasStorage            *bool
	AgeRange              *Range `gorm:"embedded;embeddedPrefix:age_"`
	IsApartment           *bool
	CreationTimeRangeFrom time.Time
	CreationTimeRangeTo   time.Time
	SearchedCount         uint
}

type Range struct {
	Min int `json:"min,omitempty"`
	Max int `json:"max,omitempty"`
}

func CreateNewFilter(db *gorm.DB, filters *Filter) error {
	result := db.Create(filters)
	if result.Error != nil {
		return result.Error
	}
	log.Printf("| Added new filter to database, Id: %d |", filters.ID)
	return nil
}
