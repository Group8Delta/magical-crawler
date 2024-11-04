package models

import "time"

type PriceHistory struct {
	ID          uint
	AddId       uint
	Price       int
	SubmittedAt time.Time
}
