package crawler

import "time"

type Ad struct {
	Title         string
	Link          string
	PhotoUrl      string
	SellerContact string
	Description   string
	Price         uint
	RentPrice     uint
	City          string
	Lat           float32
	Lon           float32
	Neighborhood  string
	Size          uint
	Bedrooms      uint
	HasElevator   bool
	HasStorage    bool
	HasParking    bool
	BuiltYear     uint
	ForRent       bool
	IsApartment   bool
	Floor         uint
	CreationTime  time.Time
}
