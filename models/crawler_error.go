package models

import "time"

type CrawlerError struct {
	ID          uint
	ErrorDetail string
	Completed   bool
	Timestamp   time.Time
	ReviewerID  uint
	User        User `gorm:"foreignKey:ReviewerID"`
}
