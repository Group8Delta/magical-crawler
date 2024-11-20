package models

import (
	"log"

	"gorm.io/gorm"
)

type User struct {
	ID           uint
	TelegramID   uint `gorm:"uniqueIndex;default 0"`
	FirstName    string
	LastName     string
	Username     string
	Email        *string
	PasswordHash *string
	RoleID       uint
	Role         Role `gorm:"constraint:OnDelete:SET NULL;"`
	Bookmarks    []Bookmark
	WatchLists   []WatchList
	FilteredAds  []FilteredAd
}

/*
	db := database.New()
	repo := repository.NewRepository(db)
	filterService := FilterServices.NewFilterServices(repo,notifier)
	//
	filterService.CreateFilter(...)


*/

func IsSuperAdmin(db *gorm.DB, userID uint) bool {
	var user User
	err := db.Preload("Role").First(&user, userID).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false
		}
		return false
	}
	return user.Role.Name == "Super Admin"
}

func IsAdmin(db *gorm.DB, userID uint) bool {
	var user User
	err := db.Preload("Role").First(&user, userID).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false
		}
		return false
	}
	return user.Role.Name == "Admin" || user.Role.Name == "Super Admin"
}

func FindOrCreateUser(db *gorm.DB, telegramID uint, firstName, lastName, username string) (*User, error) {
	var user User

	result := db.Preload("Role").First(&user, "telegram_id = ?", telegramID)
	if result.Error == gorm.ErrRecordNotFound {
		userRole, err := GetRoleByName(db, "User") // Adjust the case if your role name is "user"
		if err != nil {
			log.Printf("Error retrieving role: %v", err)
			return nil, err
		}

		user = User{
			TelegramID: telegramID,
			FirstName:  firstName,
			LastName:   lastName,
			RoleID:     userRole.ID,
			Username:   username,
		}
		if err := db.Create(&user).Error; err != nil {
			log.Printf("Error creating user: %v", err)
			return nil, err
		}

		if err := db.Preload("Role").First(&user, user.ID).Error; err != nil {
			log.Printf("Error reloading user with role: %v", err)
			return nil, err
		}
	} else if result.Error != nil {
		log.Printf("Database error: %v", result.Error)
		return nil, result.Error
	}

	return &user, nil
}
