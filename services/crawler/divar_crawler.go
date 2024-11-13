package crawler

import (
	"context"
	"errors"
	"fmt"
	"log"
	"magical-crwler/config"
	"magical-crwler/database"
	"magical-crwler/services/alerting"
	"magical-crwler/utils"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

var divar_search_urls = []string{
	"https://divar.ir/s/isfahan/buy-apartment",
	"https://divar.ir/s/isfahan/buy-villa",
	"https://divar.ir/s/isfahan/rent-apartment",
	"https://divar.ir/s/isfahan/rent-villa",
}

type DivarCrawler struct {
	config       *config.Config
	dbRepository database.IRepository
	maxDeepth    int
	alerter      *alerting.Alerter
}

func (c *DivarCrawler) CrawlAdsLinks(ctx context.Context, url string) ([]string, error) {
	ctx, cancel := chromedp.NewContext(ctx)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 100*time.Second)
	defer cancel()
	select {
	case <-ctx.Done():
		return make([]string, 0), errors.New("time-out")
	default:
		deepth := 0
		var lastHeight, newHeight int64
		var allHTMLContent strings.Builder

		err := chromedp.Run(ctx,
			chromedp.Navigate(url),
		)
		if err != nil {
			return nil, err
		}
		chromedp.Sleep(500 * time.Millisecond)
		for {

			err = chromedp.Run(ctx,
				chromedp.Evaluate(`document.body.scrollHeight`, &newHeight),
			)

			if err != nil {
				log.Println(err)
				continue
			}

			if newHeight == lastHeight {
				fmt.Println("No more content to load.")
				break
			}

			lastHeight = newHeight

			var buttonExists bool
			err = chromedp.Run(ctx,
				// This evaluates whether the button is visible on the page
				chromedp.Evaluate(`document.querySelector('.post-list__load-more-btn-be092') !== null`, &buttonExists),
			)
			if err != nil {
				log.Println(err)
				continue
			}

			if buttonExists {
				err = chromedp.Run(ctx,
					chromedp.Click(".post-list__load-more-btn-be092", chromedp.ByQuery),
					chromedp.Sleep(500*time.Millisecond), // Wait for ads to load
				)
				fmt.Println("click load more")
			} else {
				err = chromedp.Run(ctx,
					chromedp.Evaluate(`window.scrollTo(0, document.body.scrollHeight)`, nil),
					chromedp.Sleep(500*time.Millisecond),
				)
			}

			if err != nil {
				log.Println(err)
				continue
			}

			var html string
			err = chromedp.Run(ctx, chromedp.OuterHTML("html", &html))
			if err != nil {
				log.Println(err)
				continue
			}

			allHTMLContent.WriteString(html)
			deepth++
			fmt.Println("load page deepth : ", deepth)
			if deepth == c.maxDeepth {
				break
			}
		}

		doc, err := goquery.NewDocumentFromReader(strings.NewReader(allHTMLContent.String()))
		if err != nil {
			return nil, err
		}

		links := []string{}
		doc.Find("a.kt-post-card__action").Each(func(i int, s *goquery.Selection) {
			href, exists := s.Attr("href")
			if exists {
				links = append(links, "https://divar.ir"+href)
			} else {
				fmt.Println("No href found in h1 link.")
			}
		})
		return links, nil
	}
}

func (c *DivarCrawler) CrawlPageUrl(ctx context.Context, pageUrl string) (*Ad, error) {
	ctx, cancel := chromedp.NewContext(ctx)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 100*time.Second)
	defer cancel()
	select {
	case <-ctx.Done():
		return nil, errors.New("time-out")
	default:
		var ad Ad = Ad{}
		var panicErr error
		defer func() {
			if r := recover(); r != nil {
				// Recover from panic and set err to indicate the panic message
				panicErr = fmt.Errorf("recovered from panic in CrawlPageUrl: %v", r)
			}
		}()
		// Create a new context for Chrome
		ctx, cancel := chromedp.NewContext(context.Background())
		defer cancel()

		// Set timeout for the context
		ctx, cancel = context.WithTimeout(ctx, 100*time.Second)
		defer cancel()

		// Variables to store extracted data
		var Title string
		var Link string = pageUrl
		var PhotoUrl string
		var SellerContact string
		var Description string
		var Price string
		var RentPrice string
		var City string
		var Neighborhood string
		var Size string
		var Bedrooms string
		var HasElevator bool
		var HasStorage bool
		var HasParking bool
		var BuiltYear string
		var ForRent bool
		var IsApartment string
		var Floor string
		var CreationTime string

		var mapUrl string
		var Lat string
		var Lon string

		var attributes []string
		var secondAttributes []string

		var details []string
		var categories []string

		err := chromedp.Run(ctx,
			chromedp.ActionFunc(func(ctx context.Context) error {
				expiryTime := cdp.TimeSinceEpoch(time.Now().Add(24 * time.Hour))
				return network.SetCookie("token", c.config.DivarToken).
					WithDomain(".divar.ir").
					WithPath("/").
					WithExpires(&expiryTime).
					Do(ctx)
			}),

			chromedp.Navigate(pageUrl),
			chromedp.Sleep(500*time.Millisecond),

			chromedp.ActionFunc(func(ctx context.Context) error {
				var nodes []*cdp.Node
				err := chromedp.Run(ctx, chromedp.Nodes(`button.kt-button.kt-button--primary.post-actions__get-contact`, &nodes, chromedp.AtLeast(0)))
				if err == nil && len(nodes) > 0 {
					return chromedp.Click(`button.kt-button.kt-button--primary.post-actions__get-contact`, chromedp.NodeVisible).Do(ctx)
				}
				return nil // Ignore if element is not found
			}),

			chromedp.Sleep(500*time.Millisecond),

			chromedp.ActionFunc(func(ctx context.Context) error {
				var nodes []*cdp.Node
				err := chromedp.Run(ctx, chromedp.Nodes(`h1[class*="kt-page-title__title kt-page-title__title--responsive-sized"]`, &nodes, chromedp.AtLeast(0)))
				if err == nil && len(nodes) > 0 {
					return chromedp.Text(`h1[class*="kt-page-title__title kt-page-title__title--responsive-sized"]`, &Title, chromedp.NodeVisible).Do(ctx)
				}
				return nil
			}),

			chromedp.ActionFunc(func(ctx context.Context) error {
				var nodes []*cdp.Node
				err := chromedp.Run(ctx, chromedp.Nodes(`div[class*="kt-page-title__subtitle kt-page-title__subtitle--responsive-sized"]`, &nodes, chromedp.AtLeast(0)))
				if err == nil && len(nodes) > 0 {
					return chromedp.Text(`div[class*="kt-page-title__subtitle kt-page-title__subtitle--responsive-sized"]`, &CreationTime, chromedp.NodeVisible).Do(ctx)
				}
				return nil
			}),

			chromedp.ActionFunc(func(ctx context.Context) error {
				var nodes []*cdp.Node
				err := chromedp.Run(ctx, chromedp.Nodes(`p[class*="kt-description-row__text kt-description-row__text--primary"]`, &nodes, chromedp.AtLeast(0)))
				if err == nil && len(nodes) > 0 {
					return chromedp.Text(`p[class*="kt-description-row__text kt-description-row__text--primary"]`, &Description, chromedp.NodeVisible).Do(ctx)
				}
				return nil
			}),

			chromedp.ActionFunc(func(ctx context.Context) error {
				var nodes []*cdp.Node
				err := chromedp.Run(ctx, chromedp.Nodes(`img.kt-image-block__image.kt-image-block__image--fading`, &nodes, chromedp.AtLeast(0)))
				if err == nil && len(nodes) > 0 {
					return chromedp.AttributeValue(`img.kt-image-block__image.kt-image-block__image--fading`, "src", &PhotoUrl, nil).Do(ctx)
				}
				return nil
			}),

			chromedp.ActionFunc(func(ctx context.Context) error {
				var nodes []*cdp.Node
				err := chromedp.Run(ctx, chromedp.Nodes(`a.map-cm__attribution.map-cm__button`, &nodes, chromedp.AtLeast(0)))
				if err == nil && len(nodes) > 0 {
					return chromedp.AttributeValue(`a.map-cm__attribution.map-cm__button`, "href", &mapUrl, nil).Do(ctx)
				}
				return nil
			}),

			chromedp.ActionFunc(func(ctx context.Context) error {
				var nodes []*cdp.Node
				err := chromedp.Run(ctx, chromedp.Nodes(`.kt-group-row-item.kt-group-row-item__value.kt-group-row-item--info-row`, &nodes, chromedp.AtLeast(0)))
				if err == nil && len(nodes) > 0 {
					return chromedp.Evaluate(`Array.from(document.querySelectorAll(".kt-group-row-item.kt-group-row-item__value.kt-group-row-item--info-row")).map(el => el.innerText)`, &attributes).Do(ctx)
				}
				return nil
			}),

			chromedp.ActionFunc(func(ctx context.Context) error {
				var nodes []*cdp.Node
				err := chromedp.Run(ctx, chromedp.Nodes(`.kt-group-row-item.kt-group-row-item__value.kt-body.kt-body--stable`, &nodes, chromedp.AtLeast(0)))
				if err == nil && len(nodes) > 0 {
					return chromedp.Evaluate(`Array.from(document.querySelectorAll(".kt-group-row-item.kt-group-row-item__value.kt-body.kt-body--stable")).map(el => el.innerText)`, &secondAttributes).Do(ctx)
				}
				return nil
			}),

			chromedp.ActionFunc(func(ctx context.Context) error {
				var nodes []*cdp.Node
				err := chromedp.Run(ctx, chromedp.Nodes(`a.kt-unexpandable-row__action.kt-text-truncate`, &nodes, chromedp.AtLeast(0)))
				if err == nil && len(nodes) > 0 {
					return chromedp.Text(`a.kt-unexpandable-row__action.kt-text-truncate`, &SellerContact, chromedp.NodeVisible).Do(ctx)
				}
				return nil
			}),

			chromedp.ActionFunc(func(ctx context.Context) error {
				var nodes []*cdp.Node
				err := chromedp.Run(ctx, chromedp.Nodes(`div.kt-base-row.kt-base-row--large.kt-unexpandable-row`, &nodes, chromedp.AtLeast(0)))
				if err == nil && len(nodes) > 0 {
					return chromedp.Evaluate(`Array.from(document.querySelectorAll("div.kt-base-row.kt-base-row--large.kt-unexpandable-row")).map(el => el.innerText)`, &details).Do(ctx)
				}
				return nil
			}),

			chromedp.ActionFunc(func(ctx context.Context) error {
				var nodes []*cdp.Node
				err := chromedp.Run(ctx, chromedp.Nodes(`span.kt-breadcrumbs__action-text`, &nodes, chromedp.AtLeast(0)))
				if err == nil && len(nodes) > 0 {
					return chromedp.Evaluate(`Array.from(document.querySelectorAll("span.kt-breadcrumbs__action-text")).map(el => el.innerText)`, &categories).Do(ctx)
				}
				return nil
			}),
		)

		if err != nil {
			return nil, err
		}

		for _, v := range details {
			if strings.Contains(v, "قیمت کل") {
				Price = strings.Trim(strings.Split(v, "\n")[2], "تومان")

			}
			if strings.Contains(v, "ودیعه") && !strings.Contains(v, "اجاره") {
				Price = strings.Trim(strings.Split(v, "\n")[2], "تومان")

			}

			if strings.Contains(v, "اجارهٔ ماهانه") {
				RentPrice = strings.Trim(strings.Split(v, "\n")[2], "تومان")
				ForRent = true

			}
			if strings.Contains(v, "طبقه") {
				Floor = strings.Trim(strings.Split(v, "\n")[2], " ")
				if strings.Contains(Floor, "از") {
					Floor = strings.Split(Floor, "از")[0]
				}

			}
		}

		if len(attributes) > 0 {
			Size = attributes[0]
			BuiltYear = attributes[1]
			Bedrooms = attributes[2]
		}

		if len(secondAttributes) > 0 {
			for _, v := range secondAttributes {
				if strings.Contains(v, "پارکینگ") && !strings.Contains(v, "ندارد") {
					HasParking = true
				}
				if strings.Contains(v, "انباری") && !strings.Contains(v, "ندارد") {
					HasStorage = true
				}
				if strings.Contains(v, "آسانسور") && !strings.Contains(v, "ندارد") {
					HasElevator = true
				}
			}
		}

		City = strings.Split(CreationTime, "در")[1]
		if len(strings.Split(City, "،")) > 1 {
			Neighborhood = strings.Split(City, "،")[1]
			City = strings.Split(City, "،")[0]
		}
		CreationTime = strings.Split(CreationTime, "در")[0]
		IsApartment = categories[2]

		if mapUrl != "" {

			u, err := url.Parse(mapUrl)
			if err != nil {
				fmt.Println("Failed to parse map URL", err)
			}

			Lat = u.Query().Get("latitude")
			Lon = u.Query().Get("longitude")

		}

		ad.Title = Title
		ad.SellerContact = SellerContact

		pr, err := utils.PersianToEnglishDigits((Price))
		if err != nil {
			fmt.Println("invalid price ")
		}
		ad.Price = uint(pr)

		size, err := utils.PersianToEnglishDigits((Size))
		if err != nil {
			fmt.Println("invalid size ")
		}
		ad.Size = uint(size)

		buildYear, err := utils.PersianToEnglishDigits((BuiltYear))
		if err != nil {
			fmt.Println("invalid buildYear ")
		}
		ad.BuiltYear = uint(buildYear)

		bedrooms, err := utils.PersianToEnglishDigits((Bedrooms))
		if err != nil {
			fmt.Println("invalid bedrooms ")
		}
		ad.Bedrooms = uint(bedrooms)

		cr, err := utils.ParsePersianDate(strings.Trim(CreationTime, " "))
		if err != nil {
			fmt.Println(err)
		}
		ad.CreationTime = cr
		ad.HasElevator = HasElevator
		ad.HasParking = HasParking
		ad.HasStorage = HasStorage

		ad.Description = Description
		ad.City = City
		ad.Neighborhood = Neighborhood

		ad.Link = Link
		ad.PhotoUrl = PhotoUrl

		floor, err := utils.PersianToEnglishDigits((Floor))
		if err != nil {
			fmt.Println("invalid floor ")
		}
		ad.Floor = uint(floor)

		ad.IsApartment = strings.Contains(IsApartment, "آپارتمان")
		ad.ForRent = ForRent

		rpr, err := utils.PersianToEnglishDigits((RentPrice))
		if err != nil {
			fmt.Println("invalid RentPrice ")
		}
		ad.RentPrice = uint(rpr)

		lat, err := strconv.ParseFloat(Lat, 64)
		if err != nil {
			fmt.Println("invalid lat ")
		}
		ad.Lat = float32(lat)

		lon, err := strconv.ParseFloat(Lon, 64)
		if err != nil {
			fmt.Println("invalid Lon ")
		}
		ad.Lon = float32(lon)

		return &ad, panicErr
	}

}

// max deepth 0 means crawl infinity
func (c *DivarCrawler) RunCrawler() {
	go func() {
		for _, v := range divar_search_urls {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)

			wp := NewWorkerPool(v, numberOfCrawlerWorkers, c)

			wp.Start(ctx)
			results := wp.GetResults()
			errors := wp.GetErrors()

			for _, v := range errors {
				c.alerter.SendAlert(&alerting.Alert{Title: "divar crawler error", Content: v.String()})
				fmt.Println(v.Err.Error())
			}

			for _, v := range results {
				err := SaveAdData(c.dbRepository, v.Ad)
				if err != nil {
					log.Printf("error in save ad data: %s\n", err.Error())
				}
			}

			fmt.Printf("errors count:%v\n", len(errors))
			cancel()

		}
	}()
}
