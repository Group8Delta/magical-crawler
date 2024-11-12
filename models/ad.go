package models

import "time"

type Ad struct {
	ID            uint
	Link          string
	PhotoUrl      *string
	SellerName    string
	SellerContact string
	Description   *string
	Price         *int64
	RentPrice     *int
	City          *string
	Neighborhood  *string
	Size          *int
	Bedrooms      *int
	HasElevator   *bool
	HasStorage    *bool
	BuiltYear     *int
	ForRent       bool
	IsApartment   bool
	Floor         *int
	CreationTime  *time.Time
	VisitCount    int `gorm:"default:0"`
}
