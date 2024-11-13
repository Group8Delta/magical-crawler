package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID           uint
	TelegramID   uint `gorm:"uniqueIndex;not null"`
	FirstName    string
	LastName     string
	Email        *string
	PasswordHash *string
	RoleID       uint
	Role         Role `gorm:"constraint:OnDelete:SET NULL;"`
	Bookmarks    []Bookmark
	Filters      []UserFilter
	FilteredAds  []FilteredAd
}

func IsSuperAdmin(db *gorm.DB, userID int64) bool {
	var user User
	err := db.First(&user, userID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false
		}
		return false
	}
	return user.Role.Name == "Super Admin"
}