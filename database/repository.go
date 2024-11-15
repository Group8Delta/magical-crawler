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
	// Ads repository methods
	CreateAd(ad Dtos.AdDto) models.Ad
	UpdateAd(ad Dtos.AdDto) (models.Ad, error)
	BatchCreateAds(ads []Dtos.AdDto)
	DeleteAd(id int)
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

func (r *Repository) CreateAd(adDto Dtos.AdDto) models.Ad {
	ad := models.Ad{
		ID:            0,
		Link:          adDto.Link,
		PhotoUrl:      adDto.PhotoUrl,
		SellerName:    adDto.SellerName,
		SellerContact: adDto.SellerContact,
		Description:   adDto.Description,
		Price:         adDto.Price,
		RentPrice:     adDto.RentPrice,
		City:          adDto.City,
		Neighborhood:  adDto.Neighborhood,
		Size:          adDto.Size,
		Bedrooms:      adDto.Bedrooms,
		HasElevator:   adDto.HasElevator,
		HasStorage:    adDto.HasStorage,
		BuiltYear:     adDto.BuiltYear,
		ForRent:       adDto.ForRent,
		IsApartment:   adDto.IsApartment,
		Floor:         adDto.Floor,
		VisitCount:    0,
	}
	r.db.GetDb().Create(&ad)
	return ad
}
func (r *Repository) BatchCreateAds(adsReq []Dtos.AdDto) {
	ads := make([]models.Ad, 0, len(adsReq))
	for _, adDto := range adsReq {
		ad := models.Ad{
			ID:            0,
			Link:          adDto.Link,
			PhotoUrl:      adDto.PhotoUrl,
			SellerName:    adDto.SellerName,
			SellerContact: adDto.SellerContact,
			Description:   adDto.Description,
			Price:         adDto.Price,
			RentPrice:     adDto.RentPrice,
			City:          adDto.City,
			Neighborhood:  adDto.Neighborhood,
			Size:          adDto.Size,
			Bedrooms:      adDto.Bedrooms,
			HasElevator:   adDto.HasElevator,
			HasStorage:    adDto.HasStorage,
			BuiltYear:     adDto.BuiltYear,
			ForRent:       adDto.ForRent,
			IsApartment:   adDto.IsApartment,
			Floor:         adDto.Floor,
			VisitCount:    0,
		}
		ads = append(ads, ad)
	}
	r.db.GetDb().Create(&ads)
	return
}

func (r *Repository) UpdateAd(adDto Dtos.AdDto) (models.Ad, error) {
	ad := models.Ad{}
	res := r.db.GetDb().Where("ID = ?", adDto.ID).First(&ad)
	if res.Error != nil {
		return ad, res.Error
	}
	ad.Link = adDto.Link
	ad.City = adDto.City
	ad.Description = adDto.Description
	ad.PhotoUrl = adDto.PhotoUrl
	ad.SellerName = adDto.SellerName
	ad.SellerContact = adDto.SellerContact
	ad.Description = adDto.Description
	ad.Price = adDto.Price
	ad.RentPrice = adDto.RentPrice
	ad.IsApartment = adDto.IsApartment
	ad.HasElevator = adDto.HasElevator
	ad.HasStorage = adDto.HasStorage
	ad.BuiltYear = adDto.BuiltYear
	ad.ForRent = adDto.ForRent
	ad.IsApartment = adDto.IsApartment
	ad.Floor = adDto.Floor
	r.db.GetDb().Save(&ad)
	return ad, nil
}

func (r *Repository) DeleteAd(id int) {
	r.db.GetDb().Where("ID = ?", id).Delete(&models.Ad{})
}

func NewRepository(db DbService) *Repository {
	return &Repository{db: db}
}
