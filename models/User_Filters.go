package models

type UserFilter struct {
	ID           int64
	UserID       int64
	FilterID     int64
	ShouldUpdate bool
}
