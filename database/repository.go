package database

import (
	"magical-crwler/models"
	"magical-crwler/models/Dtos"
	"time"
)

type IRepository interface {
	// Filter repository methods
	CreateFilter(filter Dtos.FilterDto) models.Filter
	GetFilterById(id int) (models.Filter, error)
	UpdateFilter(filter Dtos.FilterDto) (models.Filter, error)
	DeleteFilter(id int)
	GetAdByLink(link string) (*models.Ad, error)
	CreateAd(ad Dtos.AdDto) *models.Ad
	UpdateAd(ad Dtos.AdDto) (*models.Ad, error)
	CreatePriceHistory(ph Dtos.PriceHistoryDto) *models.PriceHistory
	GetAdminUsers() ([]*models.User, error)
	GetAdsByFilterId(filterId int) ([]models.Ad, error)
	// Log
	AddLog(log models.Log)
	//GetLogLevelByName(name string) (models.LogLevel, error)
	GetAllFilters() ([]models.Filter, error)
	SearchAds(filter models.Filter, args ...string) ([]models.Ad, error)
	GetExistingFiltersAds(filter models.Filter) ([]int, error)
	GetAFilterOwner(filter models.Filter) (models.User, error)
	GetAdsByIDs(ids []int) ([]models.Ad, error)
	SaveFilterAds(adIDs []int, userID uint, filterID uint) error
}

type Repository struct {
	db DbService
}

func (r *Repository) CreateFilter(filter Dtos.FilterDto) models.Filter {
	cf := models.Filter{
		SearchQuery:           filter.SearchQuery,
		PriceRange:            filter.PriceRange,
		RentPriceRange:        filter.RentPriceRange,
		ForRent:               filter.ForRent,
		City:                  filter.City,
		Neighborhood:          filter.Neighborhood,
		SizeRange:             filter.SizeRange,
		BedroomRange:          filter.BedroomRange,
		FloorRange:            filter.FloorRange,
		HasElevator:           filter.HasElevator,
		HasStorage:            filter.HasStorage,
		AgeRange:              filter.AgeRange,
		IsApartment:           filter.IsApartment,
		CreationTimeRangeFrom: filter.CreationTimeRangeFrom,
		CreationTimeRangeTo:   filter.CreationTimeRangeTo,
	}

	r.db.GetDb().Create(&cf)
	return cf
}

func (r *Repository) GetFilterById(id int) (models.Filter, error) {
	filter := models.Filter{}
	res := r.db.GetDb().Where("ID = ?", id).First(&filter)
	return filter, res.Error
}

func (r *Repository) UpdateFilter(filter Dtos.FilterDto) (models.Filter, error) {
	f := models.Filter{}
	res := r.db.GetDb().Where("ID = ?", filter.ID).First(&f)
	if res.Error != nil {
		return f, res.Error
	}

	f.City = filter.City
	f.Neighborhood = filter.Neighborhood
	f.SizeRange = filter.SizeRange
	f.BedroomRange = filter.BedroomRange
	f.FloorRange = filter.FloorRange
	f.HasElevator = filter.HasElevator
	f.HasStorage = filter.HasStorage
	f.AgeRange = filter.AgeRange
	f.IsApartment = filter.IsApartment
	f.CreationTimeRangeFrom = filter.CreationTimeRangeFrom
	f.CreationTimeRangeTo = filter.CreationTimeRangeTo
	f.SearchQuery = filter.SearchQuery
	f.PriceRange = filter.PriceRange
	f.RentPriceRange = filter.RentPriceRange
	f.ForRent = filter.ForRent
	r.db.GetDb().Save(&f)
	return f, nil
}

func (r *Repository) DeleteFilter(id int) {
	r.db.GetDb().Where("ID = ?", id).Delete(&models.Filter{})
}

func (r *Repository) GetAdByLink(link string) (*models.Ad, error) {
	ad := models.Ad{}
	res := r.db.GetDb().Where("link = ?", link).First(&ad)
	return &ad, res.Error
}

func (r *Repository) CreateAd(ad Dtos.AdDto) *models.Ad {
	adm := models.Ad{
		Link:          ad.Link,
		PhotoUrl:      ad.PhotoUrl,
		SellerName:    ad.SellerName,
		SellerContact: ad.SellerContact,
		Description:   ad.Description,
		Price:         ad.Price,
		RentPrice:     ad.RentPrice,
		City:          ad.City,
		Neighborhood:  ad.Neighborhood,
		Size:          ad.Size,
		Bedrooms:      ad.Bedrooms,
		HasElevator:   ad.HasElevator,
		HasStorage:    ad.HasStorage,
		BuiltYear:     ad.BuiltYear,
		ForRent:       ad.ForRent,
		IsApartment:   ad.IsApartment,
		Floor:         ad.Floor,
		CreationTime:  ad.CreationTime,
	}

	r.db.GetDb().Create(&adm)
	return &adm
}

func (r *Repository) UpdateAd(ad Dtos.AdDto) (*models.Ad, error) {
	a := models.Ad{}
	res := r.db.GetDb().Where("id = ?", ad.ID).First(&a)
	if res.Error != nil {
		return &a, res.Error
	}

	a.Link = ad.Link
	a.PhotoUrl = ad.PhotoUrl
	a.SellerName = ad.SellerName
	a.SellerContact = ad.SellerContact
	a.Description = ad.Description
	a.Price = ad.Price
	a.RentPrice = ad.RentPrice
	a.City = ad.City
	a.Neighborhood = ad.Neighborhood
	a.Size = ad.Size
	a.Bedrooms = ad.Bedrooms
	a.HasElevator = ad.HasElevator
	a.HasStorage = ad.HasStorage
	a.BuiltYear = ad.BuiltYear
	a.ForRent = ad.ForRent
	a.IsApartment = ad.IsApartment
	a.Floor = ad.Floor
	r.db.GetDb().Save(&a)
	return &a, nil
}

func (r *Repository) CreatePriceHistory(ph Dtos.PriceHistoryDto) *models.PriceHistory {
	p := models.PriceHistory{
		AdID:        ph.AdID,
		Price:       ph.Price,
		RentPrice:   ph.RentPrice,
		SubmittedAt: time.Now(),
	}

	r.db.GetDb().Create(&p)
	return &p
}

func (r *Repository) GetAdminUsers() ([]*models.User, error) {
	var users []*models.User

	result := r.db.GetDb().Where("role_id < ?", "3").Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
func (r *Repository) AddLog(log models.Log) {
	log.ID = 0
	r.db.GetDb().Create(&log)
}

func (r *Repository) GetAdsByFilterId(filterId int) ([]models.Ad, error) {
	filter, err := r.GetFilterById(filterId)
	if err != nil {
		return nil, err
	}

	ads, err := r.SearchAds(filter)
	if err != nil {
		return nil, err
	}
	return ads, nil
}

//func (r *Repository) GetLogLevelByName(name string) (models.LogLevel, error) {
//
//	r.db.GetDb().Where("name = ?", name).First(&models.LogLevel{})
//}

func NewRepository(db DbService) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAllFilters() ([]models.Filter, error) {
	var filters []models.Filter
	result := r.db.GetDb().Find(&filters)
	if result.Error != nil {
		return nil, result.Error
	}
	return filters, nil
}

func (r *Repository) GetExistingFiltersAds(filter models.Filter) ([]int, error) {

	var adIDs []int
	err := r.db.GetDb().
		Model(&models.FilteredAd{}).
		Where("filter_id=?", filter.ID).
		Select("ad_id").Find(&adIDs).Error
	return adIDs, err
}

func (r *Repository) GetAdsByIDs(ids []int) ([]models.Ad, error) {
	var ads []models.Ad
	err := r.db.GetDb().
		Model(&models.FilteredAd{}).
		Where("filter_id in ?", ids).Find(&ads).Error
	return ads, err
}

func (r *Repository) GetAFilterOwner(filter models.Filter) (models.User, error) {
	var userID int
	err := r.db.GetDb().
		Model(&models.FilteredAd{}).
		Where("filter_id=?", filter.ID).
		Select("user_id").First(&userID).Error
	if err != nil {
		return models.User{}, err
	}
	var user models.User
	err = r.db.GetDb().
		Model(&models.User{}).
		Where("id=?", filter.ID).First(&user).Error

	return user, err
}
func (r *Repository) SaveFilterAds(adIDs []int, userID uint, filterID uint) error {
	filterAds := make([]models.FilteredAd, 0)
	for i := 0; i < len(adIDs); i++ {
		filterAds = append(filterAds, models.FilteredAd{UserID: userID, FilterID: filterID, AdID: uint(adIDs[i])})
	}
	result := r.db.GetDb().Create(&filterAds)
	return result.Error
}

func (r *Repository) SearchAds(filter models.Filter, args ...string) ([]models.Ad, error) {
	var ads []models.Ad
	query := r.db.GetDb().Model(&models.Ad{})
	if args != nil {
		query = query.Select(args)
	}

	if filter.SearchQuery != nil && *filter.SearchQuery != "" {
		query = query.Where("description ILIKE ? OR seller_name ILIKE ?", "%"+*filter.SearchQuery+"%", "%"+*filter.SearchQuery+"%")
	}

	if filter.PriceRange != nil {
		if filter.PriceRange.Min != 0 {
			query = query.Where("price >= ?", filter.PriceRange.Min)
		}
		if filter.PriceRange.Max != 0 {
			query = query.Where("price <= ?", filter.PriceRange.Max)
		}
	}

	if filter.RentPriceRange != nil {
		if filter.RentPriceRange.Min != 0 {
			query = query.Where("rent_price >= ?", filter.RentPriceRange.Min)
		}
		if filter.RentPriceRange.Max != 0 {
			query = query.Where("rent_price <= ?", filter.RentPriceRange.Max)
		}
	}

	query = query.Where("for_rent = ?", filter.ForRent)

	if filter.City != nil {
		query = query.Where("city = ?", *filter.City)
	}

	if filter.Neighborhood != nil {
		query = query.Where("neighborhood = ?", *filter.Neighborhood)
	}

	if filter.SizeRange != nil {
		if filter.SizeRange.Min != 0 {
			query = query.Where("size >= ?", filter.SizeRange.Min)
		}
		if filter.SizeRange.Max != 0 {
			query = query.Where("size <= ?", filter.SizeRange.Max)
		}
	}

	if filter.BedroomRange != nil {
		if filter.BedroomRange.Min != 0 {
			query = query.Where("bedrooms >= ?", filter.BedroomRange.Min)
		}
		if filter.BedroomRange.Max != 0 {
			query = query.Where("bedrooms <= ?", filter.BedroomRange.Max)
		}
	}

	if filter.FloorRange != nil {
		if filter.FloorRange.Min != 0 {
			query = query.Where("floor >= ?", filter.FloorRange.Min)
		}
		if filter.FloorRange.Max != 0 {
			query = query.Where("floor <= ?", filter.FloorRange.Max)
		}
	}

	if filter.HasElevator != nil {
		query = query.Where("has_elevator = ?", *filter.HasElevator)
	}

	if filter.HasStorage != nil {
		query = query.Where("has_storage = ?", *filter.HasStorage)
	}

	if filter.AgeRange != nil {
		if filter.AgeRange.Min != 0 {
			query = query.Where("EXTRACT(YEAR FROM AGE(now(), built_year)) >= ?", filter.AgeRange.Min)
		}
		if filter.AgeRange.Max != 0 {
			query = query.Where("EXTRACT(YEAR FROM AGE(now(), built_year)) <= ?", filter.AgeRange.Max)
		}
	}

	if filter.IsApartment != nil {
		query = query.Where("is_apartment = ?", *filter.IsApartment)
	}

	if !filter.CreationTimeRangeFrom.IsZero() {
		query = query.Where("creation_time >= ?", filter.CreationTimeRangeFrom)
	}

	if !filter.CreationTimeRangeTo.IsZero() {
		query = query.Where("creation_time <= ?", filter.CreationTimeRangeTo)
	}

	err := query.Find(&ads).Error
	return ads, err
}
