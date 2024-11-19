package utils

import (
	"fmt"
	"magical-crwler/constants"
	"magical-crwler/models"
	"strings"
)

func GenerateFilterMessage(ad models.Ad) string {

	var builder strings.Builder
	builder.WriteString("<b>آگهی‌های جدید:</b>\n\n")

	builder.WriteString(fmt.Sprintf("<b>%s</b> <a href='%s'>مشاهده آگهی</a>\n", ad.Title, ad.Link))
	builder.WriteString(fmt.Sprintf("<b>فروشنده:</b> %s\n", ad.SellerName))

	if ad.Description != nil {
		builder.WriteString(fmt.Sprintf("<b>توضیحات:</b> %s\n", *ad.Description))
	}
	if ad.Price != nil {
		builder.WriteString(fmt.Sprintf("<b>قیمت:</b> %d تومان\n", *ad.Price))
	}
	if ad.RentPrice != nil {
		builder.WriteString(fmt.Sprintf("<b>قیمت اجاره:</b> %d تومان\n", *ad.RentPrice))
	}
	if ad.Bedrooms != nil {
		builder.WriteString(fmt.Sprintf("<b>تعداد اتاق خواب:</b> %d\n", *ad.Bedrooms))
	}
	if ad.HasElevator != nil {
		if *ad.HasElevator {
			builder.WriteString(fmt.Sprintf("<b>آسانسور:</b> %s\n", "دارد"))
		} else {
			builder.WriteString(fmt.Sprintf("<b>آسانسور:</b> %s\n", "ندارد"))
		}
	}
	if ad.BuiltYear != nil {
		builder.WriteString(fmt.Sprintf("<b>سال ساخت:</b> %d\n", *ad.BuiltYear))
	}
	if ad.Floor != nil {
		builder.WriteString(fmt.Sprintf("<b>طبقه:</b> %d\n", *ad.Floor))
	}

	builder.WriteString("\n")
	return builder.String()
}

func GeneratePriceHistory(list []models.PriceHistory) string {

	var builder strings.Builder
	builder.WriteString("تغیرات قیمت:\n")

	for i := 0; i < len(list); i++ {
		item := list[i]
		if item.Price != 0 {
			builder.WriteString(fmt.Sprintf("%d. قیمت : %d \n", i+1, item.Price))
		}else if item.RentPrice != nil {
			builder.WriteString(fmt.Sprintf("%d. قیمت اجاره : %d \n", i+1, item.RentPrice))
		}
	}

	builder.WriteString("\n")
	return builder.String()
}

func GenerateCrawlerLog(logs []models.CrawlerFunctionality) string {

	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("%s:\n", constants.CrawlerStatusList))
	builder.WriteString("num date\t\t\tduration(s)\tcpu\tram(M)\tnumer of request\tsuccessful crawl\tfailed crawl\n")
	for index, log := range logs {
		formattedDate := log.Date.Format("2006-01-02 15:04:05") // Format the date
		builder.WriteString(fmt.Sprintf("%d.  %s\t\t%d\t\t%d\t%d\t%d\t\t\t%d\t\t\t%d\n", index+1, formattedDate, log.Duration, log.CPUUsage, log.RAMUsage, log.TotalRequests, log.SuccessfulRequests, log.FailedRequests))
	}

	return builder.String()
}
