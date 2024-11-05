package models

import "time"

type CrawlerFunctionality struct {
	ID                 uint
	Date               time.Time
	Duration           int
	CPUUsage           int
	RAMUsage           int
	TotalRequests      int
	SuccessfulRequests int
}
