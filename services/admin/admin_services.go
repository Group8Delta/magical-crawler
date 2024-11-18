package admin

import (
	"fmt"
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
			return fmt.Errorf("user not found: %v", userID)
		}
		log.Println("Error retrieving user:", err)
		return err
	}

	adminRole, err := models.GetRoleByName(s.db, "Admin")
	if err != nil {
		return err
	}
	if user.RoleID == adminRole.ID {
		return nil
	}

	if err := s.db.Model(&user).Update("role_id", adminRole.ID).Error; err != nil {
		log.Println("Error updating user role:", err)
		return err
	}
	return nil
}

func (s *AdminService) RemoveAdmin(userID int64) error {
	var user models.User
	if err := s.db.First(&user, userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("user not found: %v", userID)
		}
		log.Println("Error retrieving user:", err)
		return err
	}

	userRole, err := models.GetRoleByName(s.db, "User")
	if err != nil {
		return err
	}
	if user.RoleID == userRole.ID {
		return nil
	}

	if err := s.db.Model(&user).Update("role_id", userRole.ID).Error; err != nil {
		log.Println("Error updating user role:", err)
		return err
	}
	return nil
}

func (s *AdminService) ListAdmins() ([]models.User, error) {
	var admins []models.User
	adminRole, err := models.GetRoleByName(s.db, "Admin")
	if err != nil {
		log.Println("Error retrieving admin role:", err)
		return nil, err
	}

	if err := s.db.Where("role_id = ?", adminRole.ID).Find(&admins).Error; err != nil {
		log.Println("Error retrieving admin users:", err)
		return nil, err
	}
	return admins, nil
}


func (s *AdminService) ListCrawlerStausLogs() ([]models.CrawlerFunctionality, error) {
	var logs []models.CrawlerFunctionality
	if err := s.db.Find(&logs).Limit(10).Error; err != nil {
		log.Println("Error retrieving crawler staus logs:", err)
		return nil, err
	}
	return logs, nil
}
