package admin

import (
	"log"
	"magical-crwler/models"

	"gorm.io/gorm"
)

type AdminService struct {
	db *gorm.DB
}

func NewAdminService(db *gorm.DB) *AdminService {
	return &AdminService{db: db}
}

func (s *AdminService) AddAdmin(userID int64) error {
	var user models.User
	if err := s.db.First(&user, userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}
		log.Println("Error retrieving user:", err)
		return err
	}

	if user.Role.Name == "Admin" {
		return nil
	}

	adminRole, err := models.GetRoleByName(s.db, "Admin")
	if err != nil {
		return err
	}

	user.RoleID = adminRole.ID
	if err := s.db.Save(&user).Error; err != nil {
		log.Println("Error updating user role:", err)
		return err
	}
	return nil
}

func (s *AdminService) RemoveAdmin(userID int64) error {
	var user models.User
	if err := s.db.First(&user, userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}
		log.Println("Error retrieving user:", err)
		return err
	}

	if user.Role.Name != "Admin" {
		return nil
	}

	userRole, err := models.GetRoleByName(s.db, "User")
	if err != nil {
		return err
	}

	user.RoleID = userRole.ID
	if err := s.db.Save(&user).Error; err != nil {
		log.Println("Error updating user role:", err)
		return err
	}
	return nil
}
