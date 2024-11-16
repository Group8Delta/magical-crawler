package crawler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"magical-crwler/config"
	"magical-crwler/database"
	"magical-crwler/services/alerting"
	"magical-crwler/utils"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

const CurrentYear = 1403

var sheypoor_search_urls = []string{
	"https://www.sheypoor.com/s/khorramdarreh/houses-apartments-for-sale",
	"https://www.sheypoor.com/s/khorramdarreh/house-apartment-for-rent",
	"https://www.sheypoor.com/s/khorramdarreh/villa-for-sale",
}

type SheypoorContractResponse struct {
	Data struct {
		Attributes struct {
			PhoneNumber string `json:"phoneNumber"`
		} `json:"attributes"`
	} `json:"data"`
}

type SheypoorCrawler struct {
	config       *config.Config
	dbRepository database.IRepository
	maxDeepth    int
	alerter      *alerting.Alerter
}

func (c *SheypoorCrawler) CrawlAdsLinks(ctx context.Context, url string) ([]string, int, error) {
	requestCount := 0
	ctx, cancel := chromedp.NewContext(ctx)
	defer cancel()
	select {
	case <-ctx.Done():
		return make([]string, 0), requestCount, errors.New("time-out")
	default:
		deepth := 0

		var lastHeight, newHeight int64
		var allHTMLContent strings.Builder

		err := chromedp.Run(ctx,
			chromedp.Navigate(url),
		)
		requestCount++
		if err != nil {
			return nil, requestCount, err
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

			err = chromedp.Run(ctx,
				chromedp.Evaluate(`window.scrollTo(0, document.body.scrollHeight)`, nil),
				chromedp.Sleep(500*time.Millisecond),
			)
			requestCount++
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
			return nil, requestCount, err
		}

		links := []string{}
		doc.Find(`a[class="flex h-auto desktop:flex-col desktop:border-b-0 desktop:pb-0 flex-row-reverse border-b-[1px] border-dark-4 pb-4 flex-col border-none"]`).Each(func(i int, s *goquery.Selection) {
			href, exists := s.Attr("href")
			if exists && strings.Contains(href, ".html") {
				links = append(links, "https://www.sheypoor.com"+href)
			} else {
				fmt.Println("No href found")
			}
		})
		return links, requestCount, nil
	}
}

func (c *SheypoorCrawler) CrawlPageUrl(ctx context.Context, pageUrl string) (*Ad, int, error) {
	var ad Ad = Ad{}
	ad.Link = pageUrl
	requestCount := 0

	var panicError error
	defer func() {
		if r := recover(); r != nil {
			// Recover from panic and set err to indicate the panic message
			panicError = fmt.Errorf("recovered from panic in CrawlPageUrl: %v", r)
		}
	}()
	ctx, cancel := chromedp.NewContext(ctx)
	defer cancel()
	select {
	case <-ctx.Done():
		return &ad, requestCount, errors.New("time-out")
	default:
		// Variables to store extracted data
		var Title string
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
		var IsApartment bool
		var Floor int
		var CreationTime string

		var attributes string

		err := chromedp.Run(ctx,
			chromedp.Navigate(pageUrl),
			chromedp.Sleep(500*time.Microsecond),
			// Check for Title existence and then retrieve it
			chromedp.ActionFunc(func(ctx context.Context) error {
				var nodes []*cdp.Node
				if err := chromedp.Nodes(`h1[class*="mjNIv"]`, &nodes, chromedp.AtLeast(0)).Do(ctx); err != nil {
					return err
				}
				if len(nodes) > 0 {
					return chromedp.Text(`h1[class*="mjNIv"]`, &Title, chromedp.NodeVisible).Do(ctx)
				}
				return nil
			}),

			// Check for Photo URL existence and then retrieve it
			chromedp.ActionFunc(func(ctx context.Context) error {
				var nodes []*cdp.Node
				if err := chromedp.Nodes(`div.swiper-slide.swiper-slide-active.U2WwT.ylynI img`, &nodes, chromedp.AtLeast(0)).Do(ctx); err != nil {
					return err
				}
				if len(nodes) > 0 {
					return chromedp.AttributeValue(`div.swiper-slide.swiper-slide-active.U2WwT.ylynI img`, "src", &PhotoUrl, nil, chromedp.ByQuery).Do(ctx)
				}
				return nil
			}),

			// Check for Description existence and then retrieve it
			chromedp.ActionFunc(func(ctx context.Context) error {
				var nodes []*cdp.Node
				if err := chromedp.Nodes(`div[class*="MQJ5W"]`, &nodes, chromedp.AtLeast(0)).Do(ctx); err != nil {
					return err
				}
				if len(nodes) > 0 {
					return chromedp.Text(`div[class*="MQJ5W"]`, &Description, chromedp.NodeVisible).Do(ctx)
				}
				return nil
			}),

			// Check for Price existence and then retrieve it
			chromedp.ActionFunc(func(ctx context.Context) error {
				var nodes []*cdp.Node
				if err := chromedp.Nodes(`span.l29r1 strong`, &nodes, chromedp.AtLeast(0)).Do(ctx); err != nil {
					return err
				}
				if len(nodes) > 0 {
					return chromedp.Text(`span.l29r1 strong`, &Price, chromedp.NodeVisible).Do(ctx)
				}
				return nil
			}),

			// Check for City existence and then retrieve it
			chromedp.ActionFunc(func(ctx context.Context) error {
				var nodes []*cdp.Node
				if err := chromedp.Nodes(`div._3oBho`, &nodes, chromedp.AtLeast(0)).Do(ctx); err != nil {
					return err
				}
				if len(nodes) > 0 {
					return chromedp.Text(`div._3oBho`, &City, chromedp.NodeVisible).Do(ctx)
				}
				return nil
			}),

			// Check for attributes existence and then retrieve them
			chromedp.ActionFunc(func(ctx context.Context) error {
				var nodes []*cdp.Node
				if err := chromedp.Nodes(`div.bWPjU`, &nodes, chromedp.AtLeast(0)).Do(ctx); err != nil {
					return err
				}
				if len(nodes) > 0 {
					return chromedp.Text(`div.bWPjU`, &attributes, chromedp.NodeVisible).Do(ctx)
				}
				return nil
			}),
		)
		requestCount++
		if err != nil {
			return &ad, requestCount, err
		}

		if Title == "" {
			return &ad, requestCount, errors.New("error occurred in crawling this page.")
		}

		CreationTime = strings.Split(City, "،")[0]
		Neighborhood = strings.Split(City, "،")[2]
		City = strings.Split(City, "،")[1]

		for k, v := range strings.Split(attributes, "\n") {
			if strings.Contains(v, "متراژ") {
				Size = strings.Split(attributes, "\n")[k+2]
			}

			if strings.Contains(v, "نوع ملک") {
				IsApartment = strings.Contains(strings.Split(attributes, "\n")[k+2], "آپارتمان")
			}

			if strings.Contains(v, "تعداد اتاق") {
				Bedrooms = strings.Split(attributes, "\n")[k+2]
			}

			if strings.Contains(v, "پارکینگ") {
				HasParking = !strings.Contains(strings.Split(attributes, "\n")[k+2], "ندارد")
			}

			if strings.Contains(v, "انباری") {
				HasStorage = !strings.Contains(strings.Split(attributes, "\n")[k+2], "ندارد")
			}

			if strings.Contains(v, "آسانسور") {
				HasElevator = !strings.Contains(strings.Split(attributes, "\n")[k+2], "ندارد")
			}
			if strings.Contains(v, "سن بنا") {
				BuiltYear = strings.Split(attributes, "\n")[k+2]
			}
			if strings.Contains(v, "اجاره") {
				ForRent = true
				RentPrice = strings.Split(attributes, "\n")[k+2]
			}
			if strings.Contains(v, "رهن") {
				Price = strings.Split(attributes, "\n")[k+2]
			}

		}

		id := strings.Trim(strings.Split(ad.Link, "-")[len(strings.Split(ad.Link, "-"))-1], ".html")
		SellerContact, err = c.getSellerPhone(id)
		requestCount++
		if err != nil {
			fmt.Println("error in getSellerPhone", err)
		}
		ad.Title = Title
		ad.PhotoUrl = PhotoUrl
		ad.Description = Description
		ad.City = strings.Trim(City, " ")
		ad.Neighborhood = strings.Trim(Neighborhood, " ")
		size, err := utils.PersianToEnglishDigits(Size)
		if err != nil {
			fmt.Println("invalid size")
		}
		ad.Size = uint(size)

		bedrooms, err := utils.PersianToEnglishDigits(Bedrooms)
		if err != nil {
			fmt.Println("invalid bedrooms")
		}
		ad.Bedrooms = uint(bedrooms)
		ad.Floor = uint(Floor)

		ad.IsApartment = IsApartment
		ad.ForRent = ForRent
		ad.HasElevator = HasElevator
		ad.HasParking = HasParking
		ad.HasStorage = HasStorage
		ad.SellerContact = SellerContact
		ad.Floor = uint(Floor)

		Price = strings.Trim(Price, " تومان")
		RentPrice = strings.Trim(RentPrice, " تومان")
		pr, err := utils.PersianToEnglishDigits(Price)
		if err != nil {
			fmt.Println("invalid price")
		}
		ad.Price = uint(pr)

		rpr, err := utils.PersianToEnglishDigits(RentPrice)
		if err != nil {
			fmt.Println("invalid RentPrice")
		}
		ad.RentPrice = uint(rpr)

		BuiltYear = strings.Trim(BuiltYear, "سال ")
		builtYear, err := utils.PersianToEnglishDigits(BuiltYear)
		if err != nil {
			fmt.Println("invalid builtYear")
		}
		ad.BuiltYear = uint(CurrentYear - builtYear)
		cr, err := utils.ParsePersianDate(strings.Trim(CreationTime, " "))
		if err != nil {
			fmt.Println("invalid CreationTime")
		}
		ad.CreationTime = cr
		return &ad, requestCount, panicError
	}

}

func (c *SheypoorCrawler) getSellerPhone(id string) (string, error) {
	url := "https://www.sheypoor.com/api/v10.0.0/listings/" + id + "/number"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	// Setting headers
	req.Header.Set("accept", "application/vnd.api+json")
	req.Header.Set("accept-language", "en-US,en-IN;q=0.9,en;q=0.8")
	req.Header.Set("cookie", "_ga=GA1.1.288547753.1730910489; geo=city; saved_items=%5B%5D; refresh_token="+c.config.SheypoorToken+"; user_logged_in=1; provinces=; cities=; provinceID=14; province=zanjan-province; cityID=532; city=khorramdarreh; _ga_RVTCLF1865=GS1.1.1731166652.3.1.1731172603.60.0.0")
	req.Header.Set("priority", "u=1, i")
	req.Header.Set("referer", "https://www.sheypoor.com/v/%D8%A7%D8%AC%D8%A7%D8%B1%D9%87-%D9%88%D8%A7%D8%AD%D8%AF-65-%D9%85%D8%AA%D8%B1-%D8%AE-%D8%A7%D8%A8%D8%A7%D9%86-%D8%B3-%D8%A7%D9%87-%D9%84%D8%A7-%D8%AF%D9%88%D8%B7%D8%A8%D9%82%D9%87-%D8%AF%D9%88-%D9%88%D8%A7%D8%AD%D8%AF-446072148.html")
	req.Header.Set("sec-ch-ua", `"Not/A)Brand";v="8", "Chromium";v="126", "Google Chrome";v="126"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Linux"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36")
	req.Header.Set("x-user-agent", "Sheypoorx/3.6.458 browser/Chrome.126.0.0.0 os/Linux.x86_64")

	// Sending the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err

	}
	defer resp.Body.Close()

	// Reading the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err

	}

	// Unmarshal the JSON response into the Response struct
	var result SheypoorContractResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err

	}

	return result.Data.Attributes.PhoneNumber, nil
}

// max deepth 0 means crawl infinity
func (c *SheypoorCrawler) RunCrawler() {
	for _, v := range sheypoor_search_urls {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*2000)

		wp := NewWorkerPool(v, numberOfCrawlerWorkers, c)

		wp.Start(ctx)
		results := wp.GetResults()
		errors := wp.GetErrors()
		if len(errors) > 0 { // handle error in crawl
			fmt.Printf("sheypoor crawler errors time (%s):%v\n", time.Now(), errors)
			for _, v := range errors {
				c.alerter.SendAlert(&alerting.Alert{Title: "sheypoor crawler error", Content: v.String()})
				fmt.Println(v.Err.Error())
			}
		}

		for _, v := range results {
			err := SaveAdData(c.dbRepository, v.Ad)
			if err != nil {
				log.Printf("error in save ad data: %s\n", err.Error())
			}

		}

		cf, err := wp.GetCrawlerFunctionalityReport()
		if err != nil {
			fmt.Println("Error in saving shapoor functionality report ", err)
		} else {
			c.dbRepository.SaveCrawlerFunctionality(cf) // save monitoring data in database
		}
		cancel()
	}

}
