package models

import "time"

type UserFilter struct {
	ID           uint
	UserID       uint
	FilterID     uint
	ShouldUpdate bool
	UpdateCycle  int
	LastUpdated  time.Time
}
