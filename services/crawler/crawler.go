package crawler

import (
	"context"
	"errors"
	"magical-crwler/config"
	"magical-crwler/database"
	"magical-crwler/models/Dtos"
	"magical-crwler/services/alerting"
	"time"

	"gorm.io/gorm"
)

type CrawlerType string

const (
	DivarCrawlerType    CrawlerType = "divar_crawler"
	SheypoorCrawlerType CrawlerType = "sheypoor_crawler"
)
const numberOfCrawlerWorkers = 10

var CrawlerTypes []CrawlerType = []CrawlerType{
	DivarCrawlerType,
	SheypoorCrawlerType,
}

type CrawlerInterface interface {
	CrawlAdsLinks(ctx context.Context, searchUrl string) ([]string, int, error)
	CrawlPageUrl(ctx context.Context, pageUrl string) (*Ad, int, error)
	RunCrawler(timeout time.Duration)
}

func New(crawlerType CrawlerType, config *config.Config, d database.IRepository, maxDeepth int, alerter *alerting.Alerter) (CrawlerInterface, error) {
	switch crawlerType {
	case DivarCrawlerType:
		return &DivarCrawler{config: config, maxDeepth: maxDeepth, alerter: alerter, dbRepository: d}, nil
	case SheypoorCrawlerType:
		return &SheypoorCrawler{config: config, maxDeepth: maxDeepth, alerter: alerter, dbRepository: d}, nil
	default:
		return nil, errors.New("invalid crawler type")
	}
}
func SaveAdData(repo database.IRepository, ad *Ad) error {
	a, err := repo.GetAdByLink(ad.Link)

	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	if err == gorm.ErrRecordNotFound {
		price := int64(ad.Price)
		rprice := int(ad.RentPrice)
		size := int(ad.Size)
		betrooms := int(ad.Bedrooms)
		buildYear := int(ad.BuiltYear)
		floor := int(ad.Floor)
		nad := repo.CreateAd(Dtos.AdDto{Title: ad.Title, Link: ad.Link, PhotoUrl: &ad.PhotoUrl, SellerContact: ad.SellerContact, Description: &ad.Description, Price: &price, RentPrice: &rprice, City: &ad.City, Neighborhood: &ad.Neighborhood, Size: &size, Bedrooms: &betrooms, HasElevator: &ad.HasElevator, HasStorage: &ad.HasStorage, BuiltYear: &buildYear, ForRent: ad.ForRent, IsApartment: ad.IsApartment, Floor: &floor, CreationTime: &ad.CreationTime})
		repo.CreatePriceHistory(Dtos.PriceHistoryDto{AdID: uint(nad.ID), Price: *nad.Price, RentPrice: nad.RentPrice, SubmittedAt: time.Now()})
	} else {
		price := int64(ad.Price)
		rprice := int(ad.RentPrice)
		size := int(ad.Size)
		betrooms := int(ad.Bedrooms)
		buildYear := int(ad.BuiltYear)
		floor := int(ad.Floor)
		nad, err := repo.UpdateAd(Dtos.AdDto{Title: ad.Title, ID: a.ID, Link: ad.Link, PhotoUrl: &ad.PhotoUrl, SellerContact: ad.SellerContact, Description: &ad.Description, Price: &price, RentPrice: &rprice, City: &ad.City, Neighborhood: &ad.Neighborhood, Size: &size, Bedrooms: &betrooms, HasElevator: &ad.HasElevator, HasStorage: &ad.HasStorage, BuiltYear: &buildYear, ForRent: ad.ForRent, IsApartment: ad.IsApartment, Floor: &floor, CreationTime: &ad.CreationTime})
		if err != nil {
			return err
		}
		if ad.Price != uint(*a.Price) || ad.RentPrice != uint(*a.RentPrice) {
			repo.CreatePriceHistory(Dtos.PriceHistoryDto{AdID: uint(nad.ID), Price: *nad.Price, RentPrice: nad.RentPrice, SubmittedAt: time.Now()})
		}
	}
	return nil
}
