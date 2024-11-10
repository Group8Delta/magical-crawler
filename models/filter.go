package models

import (
	"gorm.io/gorm"
	"time"
)

type Filter struct {
	gorm.Model
	//ID                uint
	SearchQuery       *string
	PriceRange        *Range `gorm:"type:json"`
	RentPriceRange    *Range `gorm:"type:json"`
	ForRent           bool
	City              *string
	Neighborhood      *string
	SizeRange         *Range `gorm:"embedded"`
	BedroomRange      *Range `gorm:"embedded"`
	FloorRange        *Range `gorm:"embedded"`
	HasElevator       *bool
	HasStorage        *bool
	AgeRange          *Range `gorm:"embedded"`
	IsApartment       *bool
	CreationTimeRange *TimeRange
}

type Range struct {
	Min int `json:"min,omitempty"`
	Max int `json:"max,omitempty"`
}

type TimeRange struct {
	From time.Time `json:"from,omitempty"`
	To   time.Time `json:"to,omitempty"`
}
