package models

type CrawlerFunctionality struct {
	ID                         int64
	Duration                   float64
	CPUUsage                   float32
	RAMUsage                   float32
	NumberOfRequests           int
	NumberOfSuccessfulRequests int
}
