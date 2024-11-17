package models

import "time"

type WatchList struct {
	ID          uint
	UserID      uint
	FilterID    uint
	UpdateCycle int
	NextRunTime time.Time
	DeletedAt   *time.Time `gorm:"default:null"`
}
