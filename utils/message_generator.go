package utils

import (
	"fmt"
	"magical-crwler/models"
	"strings"
)

func GenerateFilterMessage(ads []models.Ad) string {

	var builder strings.Builder
	builder.WriteString("<b>آگهی‌های جدید مطابق با جستجوی شما:</b>\n\n") // Bold title

	for _, ad := range ads {
		builder.WriteString(fmt.Sprintf("<b>فروشنده:</b> %s\n", ad.SellerName))

		if ad.PhotoUrl != nil {
			builder.WriteString(fmt.Sprintf("<b>عکس:</b>\n<img src='%s' alt='تصویر آگهی' style='max-width:100%%;height:auto;'/>\n", *ad.PhotoUrl))
		}

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
			builder.WriteString(fmt.Sprintf("<b>آسانسور:</b> %t\n", *ad.HasElevator))
		}
		if ad.BuiltYear != nil {
			builder.WriteString(fmt.Sprintf("<b>سال ساخت:</b> %d\n", *ad.BuiltYear))
		}
		if ad.Floor != nil {
			builder.WriteString(fmt.Sprintf("<b>طبقه:</b> %d\n", *ad.Floor))
		}

		builder.WriteString(fmt.Sprintf("<b>لینک مشاهده:</b> <a href='%s'>مشاهده آگهی</a>\n", ad.Link))
		builder.WriteString("\n")
	}

	return builder.String()
}
