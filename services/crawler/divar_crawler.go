package crawler

import (
	"context"
	"fmt"
	"log"
	"magical-crwler/config"
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

type DivarCrawler struct {
	config *config.Config
	maxDeepth int
}

func (c *DivarCrawler) CrawlAdsLinks(url string) ([]string, error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 100*time.Second)
	defer cancel()
	deepth := 0
	var lastHeight, newHeight int64
	var allHTMLContent strings.Builder

	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
	)
	if err != nil {
		return nil, err
	}

	for {
		fmt.Println("load page deepth : ", deepth)
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

func (c *DivarCrawler) CrawlPageUrl(pageUrl string) (*Ad, error) {
	var ad Ad = Ad{}
	var err error
	defer func() {
		if r := recover(); r != nil {
			// Recover from panic and set err to indicate the panic message
			err = fmt.Errorf("recovered from panic in CrawlPageUrl: %v", r)
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
	var HasElevator string
	var HasStorage string
	var HasParking string
	var BuiltYear string
	var ForRent bool
	var IsApartment string
	var Floor string
	var CreationTime string

	var mapUrl string
	var Lat string
	var Lon string

	var attributes []string
	var details []string
	var categories []string

	err = chromedp.Run(ctx,
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
		chromedp.Click(`button.kt-button.kt-button--primary.post-actions__get-contact`, chromedp.NodeVisible),
		chromedp.Sleep(500*time.Millisecond),

		chromedp.Text(`h1[class*="kt-page-title__title kt-page-title__title--responsive-sized"]`, &Title, chromedp.NodeVisible),

		chromedp.Text(`div[class*="kt-page-title__subtitle kt-page-title__subtitle--responsive-sized"]`, &CreationTime, chromedp.NodeVisible),
		chromedp.Text(`p[class*="kt-description-row__text kt-description-row__text--primary"]`, &Description, chromedp.NodeVisible),

		chromedp.AttributeValue(`img.kt-image-block__image.kt-image-block__image--fading`, "src", &PhotoUrl, nil),

		chromedp.ActionFunc(func(ctx context.Context) error {
			var nodes []*cdp.Node
			err := chromedp.Run(ctx,
				chromedp.Nodes(`a.map-cm__attribution.map-cm__button`, &nodes, chromedp.AtLeast(0)),
			)
			if err != nil {
				return err
			}

			if len(nodes) > 0 {
				return chromedp.AttributeValue(`a.map-cm__attribution.map-cm__button`, "href", &mapUrl, nil).Do(ctx)
			}

			return nil
		}),

		chromedp.Evaluate(`Array.from(document.querySelectorAll(".kt-group-row__data-row")).map(el => el.innerText)`, &attributes),
		chromedp.Text(`a.kt-unexpandable-row__action.kt-text-truncate`, &SellerContact, chromedp.NodeVisible),
		chromedp.Evaluate(`Array.from(document.querySelectorAll("p.kt-unexpandable-row__value")).map(el => el.innerText)`, &details),
		chromedp.Evaluate(`Array.from(document.querySelectorAll("span.kt-breadcrumbs__action-text")).map(el => el.innerText)`, &categories),
	)

	if err != nil {
		return nil, err
	}

	firstAttributes := strings.Split(attributes[0], "\n")

	Size = firstAttributes[0]
	BuiltYear = firstAttributes[1]
	Bedrooms = firstAttributes[2]

	secondAttributes := strings.Split(attributes[1], "\n")
	HasElevator = secondAttributes[0]
	HasParking = secondAttributes[1]
	HasStorage = secondAttributes[2]
	City = strings.Split(CreationTime, "در")[1]
	if len(strings.Split(City, "،")) > 1 {
		Neighborhood = strings.Split(City, "،")[1]
		City = strings.Split(City, "،")[0]
	}
	CreationTime = strings.Split(CreationTime, "در")[0]
	IsApartment = categories[2]
	Price = strings.Trim(details[0], "تومان ")
	if strings.Contains(categories[1], "اجارهٔ") {
		ForRent = true
		RentPrice = strings.Trim(details[1], "تومان ")

	}

	if len(strings.Split(details[len(details)-1], " از")) > 1 {
		Floor = strings.Split(details[len(details)-1], " از")[0]
	} else {
		Floor = details[len(details)-1]

	}

	if mapUrl != "" {

		u, err := url.Parse(mapUrl)
		if err != nil {
			log.Println("Failed to parse map URL: %v", err)
		}

		Lat = u.Query().Get("latitude")
		Lon = u.Query().Get("longitude")

	}

	ad.Title = Title
	ad.SellerContact = SellerContact

	pr, err := utils.PersianToEnglishDigits((Price))
	if err != nil {
		log.Printf("error in converting price : %v", err)
	}
	ad.Price = uint(pr)

	size, err := utils.PersianToEnglishDigits((Size))
	if err != nil {
		log.Printf("error in converting size : %v", err)
	}
	ad.Size = uint(size)

	buildYear, err := utils.PersianToEnglishDigits((BuiltYear))
	if err != nil {
		log.Printf("error in converting buildYear : %v", err)
	}
	ad.BuiltYear = uint(buildYear)

	bedrooms, err := utils.PersianToEnglishDigits((Bedrooms))
	if err != nil {
		log.Printf("error in converting bedrooms : %v", err)
	}
	ad.Bedrooms = uint(bedrooms)

	cr, err := utils.ParsePersianDate(CreationTime)
	if err != nil {
		fmt.Println(err)
	}
	ad.CreationTime = cr
	ad.HasElevator = !strings.Contains(HasElevator, "ندارد")
	ad.HasParking = !strings.Contains(HasParking, "ندارد")
	ad.HasStorage = !strings.Contains(HasStorage, "ندارد")

	ad.Description = Description
	ad.City = City
	ad.Neighborhood = Neighborhood

	ad.Link = Link
	ad.PhotoUrl = PhotoUrl

	floor, err := utils.PersianToEnglishDigits((Floor))
	if err != nil {
		log.Printf("error in converting floor : %v", err)
	}
	ad.Floor = uint(floor)

	ad.IsApartment = strings.Contains(IsApartment, "آپارتمان")
	ad.ForRent = ForRent

	rpr, err := utils.PersianToEnglishDigits((RentPrice))
	if err != nil {
		log.Printf("error in converting RentPrice : %v", err)
	}
	ad.RentPrice = uint(rpr)

	lat, err := strconv.ParseFloat(Lat, 64)
	if err != nil {
		log.Printf("error in converting lat : %v", err)
	}
	ad.Lat = float32(lat)

	lon, err := strconv.ParseFloat(Lon, 64)
	if err != nil {
		log.Printf("error in converting Lon : %v", err)
	}
	ad.Lon = float32(lon)

	return &ad, err

}
