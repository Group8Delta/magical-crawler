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

func (r *Repository) GetAdByLink(link string) (models.Ad, error) {
	ad := models.Ad{}
	res := r.db.GetDb().Where("link = ?", link).First(&ad)
	return ad, res.Error
}

func (r *Repository) CreateAd(ad Dtos.AdDto) models.Ad {
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
	return adm
}

func (r *Repository) UpdateAd(ad Dtos.AdDto) (models.Ad, error) {
	a := models.Ad{}
	res := r.db.GetDb().Where("id = ?", ad.ID).First(&a)
	if res.Error != nil {
		return a, res.Error
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
	return a, nil
}

func (r *Repository) CreatePriceHistory(ph Dtos.PriceHistoryDto) models.PriceHistory {
	p := models.PriceHistory{
		AdID:        ph.AdID,
		Price:       ph.Price,
		RentPrice:   ph.RentPrice,
		SubmittedAt: time.Now(),
	}

	r.db.GetDb().Create(&p)
	return p
}

func NewRepository(db DbService) *Repository {
	return &Repository{db: db}
}
