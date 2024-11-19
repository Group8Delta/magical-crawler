package models

import "time"

type FilteredAd struct {
	ID        uint
	UserID    uint `gorm:"not null;uniqueIndex:idx_user_filter_ad"`
	FilterID  uint `gorm:"not null;uniqueIndex:idx_user_filter_ad"`
	AdID      uint `gorm:"not null;uniqueIndex:idx_user_filter_ad"`
	TimeStamp time.Time
	Filter    Filter
	Ad        Ad
}
