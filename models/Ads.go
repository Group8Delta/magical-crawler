package models

type Category string

const (
	CategoryForRent Category = "Rent"
	CategoryForSell Category = "Sell"
)

type Ad struct {
	Id                  int64
	Link                string
	Photo_url           string
	Seller_Name         string
	Seller_Phone_Number string
	Description         string
	Price               float64
	City                string
	Neighborhood        string
	Size                float64
	No_Of_Rooms         int
	HasElevator         bool
	Has_Store           bool
	Age                 int
	Category            string // should be enum
}
