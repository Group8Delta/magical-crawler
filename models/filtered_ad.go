package models

import "time"

type FilteredAd struct {
	ID        uint
	UserID    uint
	FilterID  uint
	AdID      uint
	TimeStamp time.Time
	Filter    Filter
	Ad        Ad
}
