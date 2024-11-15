package database

import (
	"magical-crwler/models"
	"magical-crwler/models/Dtos"
)

type IRepository interface {
	// Filter repository methods
	CreateFilter(filter Dtos.FilterDto) models.Filter
	GetFilterById(id int) (models.Filter, error)
	UpdateFilter(filter Dtos.FilterDto) (models.Filter, error)
	DeleteFilter(id int)
	// Log
	AddLog(log models.Log)
	//GetLogLevelByName(name string) (models.LogLevel, error)
	GetAllFilters() ([]models.Filter, error)
	SearchAdIDs(filter models.Filter) ([]int, error)
	GetExistingFiltersAds(filter models.Filter) ([]int, error)
	GetAFilterOwner(filter models.Filter) (models.User, error)
	GetAdsByIDs(ids []int) ([]models.Ad, error)
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

func (r *Repository) AddLog(log models.Log) {
	log.ID = 0
	r.db.GetDb().Create(&log)
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

func (r *Repository) SearchAdIDs(filter models.Filter) ([]int, error) {
	var adIDs []int
	query := r.db.GetDb().Model(&models.Ad{}).Select("id")
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

	err := query.Find(&adIDs).Error
	return adIDs, err
}
