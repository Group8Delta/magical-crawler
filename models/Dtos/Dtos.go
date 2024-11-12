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
