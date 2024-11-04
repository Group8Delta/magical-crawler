package models

import "time"

type Filter struct {
	ID                uint
	SearchQuery       string
	PriceRange        *Range `gorm:"type:json"`
	RentPriceRange    *Range `gorm:"type:json"`
	ForRent           bool
	City              string
	Neighborhood      string
	SizeRange         *Range `gorm:"type:json"`
	BedroomRange      *Range `gorm:"type:json"`
	FloorRange        *Range `gorm:"type:json"`
	FilterRange       *Range `gorm:"type:json"`
	HasElevator       bool
	HasStorage        bool
	AgeRange          *Range `gorm:"type:json"`
	IsApartment       bool
	CreationTimeRange *TimeRange
}

type Range struct {
	Max int
	Min int
}

type TimeRange struct {
	From time.Time
	To   time.Time
}
