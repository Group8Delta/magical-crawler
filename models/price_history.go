package models

import "time"

type PriceHistory struct {
	ID          uint
	AdID        uint
	Price       int64
	RentPrice   *int
	SubmittedAt time.Time
}
