package Dtos

import (
	"magical-crwler/models"
	"time"
)

type FilterDto struct {
	ID                    uint
	SearchQuery           *string
	PriceRange            *models.Range
	RentPriceRange        *models.Range
	ForRent               bool
	City                  *string
	Neighborhood          *string
	SizeRange             *models.Range
	BedroomRange          *models.Range
	FloorRange            *models.Range
	HasElevator           *bool
	HasStorage            *bool
	AgeRange              *models.Range
	IsApartment           *bool
	CreationTimeRangeFrom time.Time
	CreationTimeRangeTo   time.Time
}

type AdDto struct {
	ID            uint
	Title         string
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
}

type PriceHistoryDto struct {
	ID          uint
	AdID        uint
	Price       int64
	RentPrice   *int
	SubmittedAt time.Time
}

type WatchListDto struct {
	UserId      int
	FilterId    int
	UpdateCycle int
}
