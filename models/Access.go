package models

type Access struct {
	ID            uint
	OwnerID       uint
	AccessedByID  uint
	AccessLevelID uint
	Owner         User `gorm:"foreignkey:OwnerID"`
	AccessedBy    User `gorm:"foreignkey:OwnerID"`
}
