package models

type User struct {
	ID           uint
	FirstName    string
	LastName     string
	Email        *string
	PasswordHash *string
	RoleId       *uint
	Role         *Role `gorm:"constraint:OnDelete:SET NULL;"`
	Bookmarks    []Bookmark
	Filters      []UserFilter
	FilteredAds  []FilteredAd
}
