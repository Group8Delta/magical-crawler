package models

import "gorm.io/gorm"

type Bookmark struct {
	gorm.Model
	AdID     uint
	Ad       Ad
	UserID   uint
	IsPublic bool
}
