package models

import "time"

type Filter struct {
	ID                uint
	SearchQuery       *string
	PriceRange        *Range `gorm:"type:json"`
	RentPriceRange    *Range `gorm:"type:json"`
	ForRent           bool
	City              *string
	Neighborhood      *string
	SizeRange         *Range `gorm:"type:json"`
	BedroomRange      *Range `gorm:"type:json"`
	FloorRange        *Range `gorm:"type:json"`
	FilterRange       *Range `gorm:"type:json"`
	HasElevator       *bool
	HasStorage        *bool
	AgeRange          *Range `gorm:"type:json"`
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
