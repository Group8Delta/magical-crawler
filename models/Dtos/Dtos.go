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
	SearchedCount         int
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
	VisitCount    int
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

type PopularFiltersDto struct {
	FilterName string
	Value      string
	Count      int
}

type UserCrawlInfo struct {
	UserID    uint
	FirstName string
	LastName  string
	Filters   []FilterWithAds
}

type FilterWithAds struct {
	FilterID uint
	Filter   models.Filter
	Ads      []AdSummary
}

type AdSummary struct {
	ID          uint
	Link        string
	PhotoUrl    *string
	Description *string
	Price       *int64
}
type AccessDto struct {
	OwnerID       uint
	AccessedByID  uint
	AccessLevelID uint
}
type BookmarkDto struct {
	AdID     uint
	UserID   uint
	IsPublic bool
}
type BookmarkToShowDto struct {
	Ad       models.Ad
	IsPublic bool
}
