package models

import (
	"time"

	"gorm.io/gorm"
)

type Filter struct {
	gorm.Model
	//ID                uint
	SearchQuery           *string
	PriceRange            *Range `gorm:"embedded"`
	RentPriceRange        *Range `gorm:"embedded"`
	ForRent               bool
	City                  *string
	Neighborhood          *string
	SizeRange             *Range `gorm:"embedded"`
	BedroomRange          *Range `gorm:"embedded"`
	FloorRange            *Range `gorm:"embedded"`
	HasElevator           *bool
	HasStorage            *bool
	AgeRange              *Range `gorm:"embedded"`
	IsApartment           *bool
	CreationTimeRangeFrom time.Time
	CreationTimeRangeTo   time.Time
	SearchedCount         uint
}

type Range struct {
	Min int `json:"min,omitempty"`
	Max int `json:"max,omitempty"`
}

// type TimeRange struct {
// 	From time.Time `json:"from,omitempty"`
// 	To   time.Time `json:"to,omitempty"`
// }
