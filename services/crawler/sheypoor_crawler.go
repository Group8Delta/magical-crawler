package crawler

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
)

type SheypoorCrawler struct {
}

func (c *SheypoorCrawler) CrawlAdsLinks(url string) ([]string, error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 100*time.Second)
	defer cancel()

	max_deepth := 40
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

		err = chromedp.Run(ctx,
			chromedp.Evaluate(`window.scrollTo(0, document.body.scrollHeight)`, nil),
			chromedp.Sleep(500*time.Millisecond),
		)

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
		if deepth == max_deepth {
			break
		}
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(allHTMLContent.String()))
	if err != nil {
		return nil, err
	}

	links := []string{}
	doc.Find("a.qL9GS").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists && strings.Contains(href, ".html") {
			links = append(links, href)
		} else {
			fmt.Println("No href found")
		}
	})
	return links, nil
}

func (c *SheypoorCrawler) CrawlPageUrl(pageUrl string) (interface{}, error) {
	return nil, nil
}
