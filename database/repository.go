package database

import (
	"errors"
	"fmt"
	"magical-crwler/models"
	"magical-crwler/models/Dtos"
	"sort"
	"time"

	"gorm.io/gorm"
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
	GetPriceHistory(adID uint) ([]models.PriceHistory, error)
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
	GetCurrentMinuteWatchLists() ([]models.WatchList, error)
	DeleteWatchList(id int) error
	UpdateWatchList(id int, wl Dtos.WatchListDto) error
	CreateWatchList(wl Dtos.WatchListDto) (*models.WatchList, error)
	GetUserById(id int) (*models.User, error)
	SaveCrawlerFunctionality(cf models.CrawlerFunctionality) error
	//User methods
	GetUserByUsername(username string) (*models.User, error)
	// bookmark
	CreateBookmark(bookmark Dtos.BookmarkDto) error
	DeleteBookmark(adid, userid uint) error
	GetBookmarksByUserID(userid uint) ([]Dtos.BookmarkToShowDto, error)
	GetPublicBookmarksByUserID(userid uint) ([]Dtos.BookmarkToShowDto, error)
	// access
	AddAccess(access Dtos.AccessDto) error
	GetAccessByIds(srcid, dstid uint) Dtos.AccessDto
	GetAllAccessLevels() []models.AccessLevel
	GetMostVisitedAds(count int) ([]models.Ad, error)
	GetMostSearchedFilters(count int) ([]models.Filter, error)
	GetMostSearchedSingleFilters(count int) ([]Dtos.PopularFiltersDto, error)
	GetWatchListFiltersByTelegramId(id int) ([]models.Filter, error)
	GetUserByTelegramId(id int) (*models.User, error)
	DeleteWatchListByFilterId(filterId int, userId int) error
	IncrementVisitCount(adID uint) error
}

type Repository struct {
	db DbService
}

func (r *Repository) GetWatchListFiltersByTelegramId(id int) ([]models.Filter, error) {
	filters := []models.Filter{}
	res := r.db.GetDb().Raw("select f.* from users u inner join watch_lists wl on wl.user_id =u.id inner join filters f on f.id=wl.filter_id where wl.deleted_at is null and u.telegram_id =?", id).Scan(&filters)

	if res.Error != nil {
		return nil, res.Error
	}
	return filters, nil
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
		SearchedCount:         filter.SearchedCount,
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

func (r *Repository) GetUserById(id int) (*models.User, error) {
	user := models.User{}
	res := r.db.GetDb().Where("id = ?", id).First(&user)
	return &user, res.Error
}

func (r *Repository) GetUserByTelegramId(id int) (*models.User, error) {
	user := models.User{}
	res := r.db.GetDb().Where("telegram_id = ?", id).First(&user)
	return &user, res.Error
}

func (r *Repository) CreateAd(ad Dtos.AdDto) *models.Ad {
	adm := models.Ad{
		Title:         ad.Title,
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
		VisitCount:    ad.VisitCount,
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

	a.Title = ad.Title
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

func (r *Repository) GetPriceHistory(adID uint) ([]models.PriceHistory, error) {
	var list []models.PriceHistory
	result := r.db.GetDb().Where("id=?",adID).Find(&list).Limit(10).Order("submitted_at DESC")
	return list, result.Error
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

func (r *Repository) GetMostVisitedAds(count int) ([]models.Ad, error) {
	var ads []models.Ad
	err := r.db.GetDb().Order("visit_count DESC").Limit(count).Find(&ads).Error
	if err != nil {
		return nil, err
	}
	return ads, nil
}

func (r *Repository) GetMostSearchedFilters(count int) ([]models.Filter, error) {
	var filters []models.Filter
	err := r.db.GetDb().Order("searched_count DESC").Limit(count).Find(&filters).Error
	if err != nil {
		return nil, err
	}
	return filters, nil
}

func (r *Repository) GetMostSearchedSingleFilters(count int) ([]Dtos.PopularFiltersDto, error) {
	var results []Dtos.PopularFiltersDto

	queryAndAppend := func(field string, filterName string, whereClause string) error {
		var tempResults []struct {
			Value string
			Count int
		}

		err := r.db.GetDb().
			Model(&models.Filter{}).
			Select(fmt.Sprintf("%s AS value, SUM(searched_count) AS count", field)).
			Where(whereClause).
			Group(field).
			Order("count DESC").
			Limit(count).
			Scan(&tempResults).Error
		if err != nil {
			return err
		}

		for _, temp := range tempResults {
			results = append(results, Dtos.PopularFiltersDto{
				FilterName: filterName,
				Value:      temp.Value,
				Count:      temp.Count,
			})
		}
		return nil
	}

	err := queryAndAppend("CONCAT(price_min, '-', price_max)", "PriceRange",
		"price_min IS NOT NULL AND price_max IS NOT NULL")
	if err != nil {
		return nil, err
	}

	err = queryAndAppend("CONCAT(rent_price_min, '-', rent_price_max)", "RentPriceRange",
		"rent_price_min IS NOT NULL AND rent_price_max IS NOT NULL")
	if err != nil {
		return nil, err
	}

	err = queryAndAppend("CASE for_rent WHEN true THEN 'Yes' ELSE 'No' END", "ForRent",
		"for_rent IS NOT NULL")
	if err != nil {
		return nil, err
	}

	err = queryAndAppend("city", "City", "city IS NOT NULL")
	if err != nil {
		return nil, err
	}

	err = queryAndAppend("neighborhood", "Neighborhood", "neighborhood IS NOT NULL")
	if err != nil {
		return nil, err
	}

	err = queryAndAppend("CONCAT(size_min, '-', size_max)", "SizeRange",
		"size_min IS NOT NULL AND size_max IS NOT NULL")
	if err != nil {
		return nil, err
	}

	err = queryAndAppend("CONCAT(bedroom_min, '-', bedroom_max)", "BedroomRange",
		"bedroom_min IS NOT NULL AND bedroom_max IS NOT NULL")
	if err != nil {
		return nil, err
	}

	err = queryAndAppend("CONCAT(floor_min, '-', floor_max)", "FloorRange",
		"floor_min IS NOT NULL AND floor_max IS NOT NULL")
	if err != nil {
		return nil, err
	}

	err = queryAndAppend("CASE has_elevator WHEN true THEN 'Yes' ELSE 'No' END", "HasElevator",
		"has_elevator IS NOT NULL")
	if err != nil {
		return nil, err
	}

	err = queryAndAppend("CASE has_storage WHEN true THEN 'Yes' ELSE 'No' END", "HasStorage",
		"has_storage IS NOT NULL")
	if err != nil {
		return nil, err
	}

	err = queryAndAppend("CONCAT(age_min, '-', age_max)", "AgeRange",
		"age_min IS NOT NULL AND age_max IS NOT NULL")
	if err != nil {
		return nil, err
	}

	err = queryAndAppend("CASE is_apartment WHEN true THEN 'Yes' ELSE 'No' END", "IsApartment",
		"is_apartment IS NOT NULL")
	if err != nil {
		return nil, err
	}

	err = queryAndAppend("CONCAT(creation_time_range_from, '-', creation_time_range_to)", "CreationTimeRange",
		"creation_time_range_from IS NOT NULL AND creation_time_range_to IS NOT NULL")
	if err != nil {
		return nil, err
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Count > results[j].Count
	})

	if len(results) > count {
		results = results[:count]
	}

	return results, nil
}

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

func (r *Repository) GetCurrentMinuteWatchLists() ([]models.WatchList, error) {
	now := time.Now()
	startOfMinute := now.Truncate(time.Minute)
	endOfMinute := startOfMinute.Add(time.Minute)
	var watchLists []models.WatchList
	if err := r.db.GetDb().Where("next_run_time >= ? AND next_run_time < ? and deleted_at is null", startOfMinute, endOfMinute).Find(&watchLists).Error; err != nil {
		return nil, err
	}

	return watchLists, nil
}

func (r *Repository) DeleteWatchList(id int) error {
	w := models.WatchList{}
	res := r.db.GetDb().Where("id = ?", id).First(&w)
	if res.Error != nil {
		return res.Error
	}
	now := time.Now()
	w.DeletedAt = &now

	return r.db.GetDb().Save(&w).Error
}

func (r *Repository) IncrementVisitCount(adID uint) error {
	res := r.db.GetDb().Model(&models.Ad{}).Where("id = ?", adID).Update("visit_count", gorm.Expr("visit_count + 1"))
	return res.Error
}

func (r *Repository) DeleteWatchListByFilterId(filterId int, userId int) error {
	w := models.WatchList{}
	res := r.db.GetDb().Where("filter_id = ? and user_id = ? and deleted_at is null", filterId, userId).First(&w)
	if res.Error != nil {
		return res.Error
	}
	now := time.Now()
	w.DeletedAt = &now

	return r.db.GetDb().Save(&w).Error
}

func (r *Repository) UpdateWatchList(id int, wl Dtos.WatchListDto) error {
	w := models.WatchList{}
	res := r.db.GetDb().Where("id = ?", id).First(&w)
	if res.Error != nil {
		return res.Error
	}

	if wl.FilterId > 0 {
		w.FilterID = uint(wl.FilterId)
	}
	if wl.UpdateCycle > 0 {
		w.UpdateCycle = wl.UpdateCycle
	}
	w.NextRunTime = time.Now().Add(time.Duration(w.UpdateCycle) * time.Minute)
	return r.db.GetDb().Save(&w).Error
}

func (r *Repository) CreateWatchList(wl Dtos.WatchListDto) (*models.WatchList, error) {

	w := models.WatchList{
		UserID:      uint(wl.UserId),
		FilterID:    uint(wl.FilterId),
		UpdateCycle: wl.UpdateCycle,
		NextRunTime: time.Now().Add(time.Duration(wl.UpdateCycle) * time.Minute),
		DeletedAt:   nil,
	}

	err := r.db.GetDb().Create(&w).Error
	return &w, err
}
func (r *Repository) SaveCrawlerFunctionality(cf models.CrawlerFunctionality) error {
	return r.db.GetDb().Save(&cf).Error
}

func (r *Repository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.GetDb().Where("username = ?", username).First(&user).Error
	return &user, err
}

func (r *Repository) CreateBookmark(bookmark Dtos.BookmarkDto) error {

	bm := models.Bookmark{
		AdID:     bookmark.AdID,
		UserID:   bookmark.UserID,
		IsPublic: bookmark.IsPublic,
	}

	result := r.db.GetDb().Where("ad_id = ? AND user_id = ?", bm.AdID, bm.UserID).First(&bm)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		r.db.GetDb().Create(&bm)
		return nil
	} else if result.Error != nil {
		return result.Error
	}
	return errors.New("bookmark already exists")
}

func (r *Repository) DeleteBookmark(adid, userid uint) error {
	bm := models.Bookmark{
		AdID:   adid,
		UserID: userid,
	}

	result := r.db.GetDb().Where("ad_id = ? AND user_id = ?", bm.AdID, bm.UserID).First(&bm)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New("bookmark not found")
	}
	r.db.GetDb().Where("ad_id = ? AND user_id = ?", bm.AdID, bm.UserID).Delete(&bm)
	return nil
}

func (r *Repository) GetBookmarksByUserID(userid uint) ([]Dtos.BookmarkToShowDto, error) {
	var bookmarks []models.Bookmark
	err := r.db.GetDb().Preload("Ad").Where("user_id = ?", userid).Find(&bookmarks).Error
	bmdtos := make([]Dtos.BookmarkToShowDto, 0, len(bookmarks))
	for _, bm := range bookmarks {
		bmdtos = append(bmdtos, Dtos.BookmarkToShowDto{
			Ad:       bm.Ad,
			IsPublic: bm.IsPublic,
		})
	}
	return bmdtos, err
}

func (r *Repository) GetPublicBookmarksByUserID(userid uint) ([]Dtos.BookmarkToShowDto, error) {
	var bookmarks []models.Bookmark
	err := r.db.GetDb().Preload("Ad").Where("user_id = ? AND is_public = true", userid).Find(&bookmarks).Error
	bmdtos := make([]Dtos.BookmarkToShowDto, 0, len(bookmarks))
	for _, bm := range bookmarks {
		bmdtos = append(bmdtos, Dtos.BookmarkToShowDto{
			Ad:       bm.Ad,
			IsPublic: bm.IsPublic,
		})
	}
	return bmdtos, err
}

func (r *Repository) AddAccess(access Dtos.AccessDto) error {

	ac := models.Access{
		OwnerID:      access.OwnerID,
		AccessedByID: access.AccessedByID,
	}

	result := r.db.GetDb().Where("owner_id = ? AND accessed_by_id = ?", ac.OwnerID, ac.AccessedByID).First(&ac)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		ac.AccessedByID = access.AccessedByID
		ac.OwnerID = access.OwnerID
		ac.AccessLevelID = access.AccessLevelID
		result = r.db.GetDb().Create(&ac)
		return result.Error
	} else if result.Error != nil {
		return result.Error
	}
	ac.AccessLevelID = access.AccessLevelID
	result = r.db.GetDb().Save(&ac)
	return result.Error
}
func (r *Repository) GetAccessByIds(srcid, dstid uint) Dtos.AccessDto {
	ac := models.Access{
		OwnerID:      srcid,
		AccessedByID: dstid,
	}

	result := r.db.GetDb().Where("owner_id = ? AND accessed_by_id = ?", ac.OwnerID, ac.AccessedByID).First(&ac)
	if result.Error != nil {
		return Dtos.AccessDto{
			OwnerID:       srcid,
			AccessedByID:  dstid,
			AccessLevelID: 2,
		}
	}
	return Dtos.AccessDto{
		OwnerID:       ac.OwnerID,
		AccessedByID:  ac.AccessedByID,
		AccessLevelID: ac.AccessLevelID,
	}
}

func (r *Repository) GetAllAccessLevels() []models.AccessLevel {
	lvls := make([]models.AccessLevel, 0)
	r.db.GetDb().Find(&lvls)
	return lvls
}
