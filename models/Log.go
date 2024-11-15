package models

import "time"

type Log struct {
	ID         uint
	LogLevelID uint
	Message    string
	Time       time.Time
}
