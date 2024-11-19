package utils

import (
	"fmt"
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
