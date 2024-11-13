package models

import "gorm.io/gorm"

type Role struct {
	ID   uint
	Name string
}

func GetRoleByName(db *gorm.DB, roleName string) (Role, error) {
	var role Role
	if err := db.Where("name = ?", roleName).First(&role).Error; err != nil {
		return Role{}, err
	}
	return role, nil
}
