package models

import "time"

type WatchList struct {
	ID          uint
	UserID      uint
	FilterID    uint
	UpdateCycle int
	LastUpdated time.Time
}
