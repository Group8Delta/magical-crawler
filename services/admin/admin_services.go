package admin

import (
	"fmt"
	"log"
	"magical-crwler/models"
	"magical-crwler/models/Dtos"

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

func (s *AdminService) ListUsers() ([]models.User, error) {
	var users []models.User
	adminRole, err := models.GetRoleByName(s.db, "User")
	if err != nil {
		log.Println("Error retrieving admin role:", err)
		return nil, err
	}

	if err := s.db.Where("role_id = ?", adminRole.ID).Find(&users).Error; err != nil {
		log.Println("Error retrieving users:", err)
		return nil, err
	}
	return users, nil
}

func (s *AdminService) GetUsersCrawlInfo() ([]Dtos.FilterWithAds, error) {
	var users []models.User

	err := s.db.Preload("FilteredAds").
		Preload("FilteredAds.Filter").
		Preload("FilteredAds.Ad").
		Find(&users).Error
	if err != nil {
		return nil, err
	}

	filterMap := make(map[uint]*Dtos.FilterWithAds)

	for _, user := range users {
		for _, filteredAd := range user.FilteredAds {
			if _, exists := filterMap[filteredAd.FilterID]; !exists {
				filterMap[filteredAd.FilterID] = &Dtos.FilterWithAds{
					FilterID: filteredAd.FilterID,
					Filter:   filteredAd.Filter,
					Ads:      []Dtos.AdSummary{},
				}
			}

			filterMap[filteredAd.FilterID].Ads = append(filterMap[filteredAd.FilterID].Ads, Dtos.AdSummary{
				ID:          filteredAd.Ad.ID,
				Link:        filteredAd.Ad.Link,
				PhotoUrl:    filteredAd.Ad.PhotoUrl,
				Description: filteredAd.Ad.Description,
				Price:       filteredAd.Ad.Price,
			})
		}
	}

	var filters []Dtos.FilterWithAds
	for _, filterWithAds := range filterMap {
		filters = append(filters, *filterWithAds)
	}

	return filters, nil
}

func (s *AdminService) GetSingleUserCrawlInfo(userID int64) (*Dtos.UserCrawlInfo, error) {
	var user models.User

	err := s.db.Preload("FilteredAds.Ad").
		Preload("FilteredAds.Filter").
		First(&user, userID).Error
	if err != nil {
		return nil, err
	}

	var filters []Dtos.FilterWithAds
	for _, filteredAd := range user.FilteredAds {
		filterIndex := -1
		for i, f := range filters {
			if f.FilterID == filteredAd.FilterID {
				filterIndex = i
				break
			}
		}

		if filterIndex == -1 {
			filters = append(filters, Dtos.FilterWithAds{
				FilterID: filteredAd.FilterID,
				Filter:   filteredAd.Filter,
				Ads:      []Dtos.AdSummary{},
			})
			filterIndex = len(filters) - 1
		}

		filters[filterIndex].Ads = append(filters[filterIndex].Ads, Dtos.AdSummary{
			ID:          filteredAd.Ad.ID,
			Link:        filteredAd.Ad.Link,
			PhotoUrl:    filteredAd.Ad.PhotoUrl,
			Description: filteredAd.Ad.Description,
			Price:       filteredAd.Ad.Price,
		})
	}

	userCrawlInfo := Dtos.UserCrawlInfo{
		UserID:    user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Filters:   filters,
	}

	return &userCrawlInfo, nil
}
