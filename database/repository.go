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
