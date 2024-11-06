package bot

import (
	"fmt"
	"magical-crwler/config"

	"gopkg.in/telebot.v4"
)

type Filters struct {
	price        FilterValue
	area         FilterValue
	rooms        FilterValue
	propertyType FilterValue
	buildingAge  FilterValue
	floor        FilterValue
	storage      FilterValue
	elevator     FilterValue
	adDate       FilterValue
	location     FilterValue
}
type FilterValue struct {
	value  string
	button telebot.Btn
}

func FilterHandlers(b *Bot) func(ctx telebot.Context) error {

	selector := &telebot.ReplyMarkup{RemoveKeyboard: true}

	var (
		filters = Filters{
			price: FilterValue{button: selector.Data(config.PriceFilter, "PriceFilter"),
				value: "",
			},
			area: FilterValue{button: selector.Data(config.AreaFilter, "AreaFilter"),
				value: "",
			},
			rooms: FilterValue{button: selector.Data(config.RoomsFilter, "RoomsFilter"),
				value: "",
			},
			propertyType: FilterValue{button: selector.Data(config.PropertyTypeFilter, "PropertyTypeFilter"),
				value: "",
			},
			buildingAge: FilterValue{button: selector.Data(config.BuildingAgeFilter, "BuildingAgeFilter"),
				value: "",
			},
			floor: FilterValue{button: selector.Data(config.FloorFilter, "FloorFilter"),
				value: "",
			},
			storage: FilterValue{button: selector.Data(config.StorageFilter, "StorageFilter"),
				value: "",
			},
			elevator: FilterValue{button: selector.Data(config.ElevatorFilter, "ElevatorFilter"),
				value: "",
			},
			adDate: FilterValue{button: selector.Data(config.AdDateFilter, "AdDateFilter"),
				value: "",
			},
			location: FilterValue{button: selector.Data(config.LocationFilter, "LocationFilter"),
				value: "",
			},
		}
	)

	selector.Inline(
		selector.Row(filters.price.button),
		selector.Row(filters.area.button, filters.rooms.button, filters.propertyType.button),
		selector.Row(filters.buildingAge.button, filters.floor.button, filters.storage.button),
		selector.Row(filters.elevator.button, filters.adDate.button, filters.location.button),
	)

	printFilter := func() string {
		var f string
		f += fmt.Sprintf("%s\t:\t%s\n", filters.price.button.Text, filters.price.value)
		f += fmt.Sprintf("%s\t:\t%s\n", filters.area.button.Text, filters.area.value)
		f += fmt.Sprintf("%s\t:\t%s\n", filters.rooms.button.Text, filters.rooms.value)
		f += fmt.Sprintf("%s\t:\t%s\n", filters.propertyType.button.Text, filters.propertyType.value)
		f += fmt.Sprintf("%s\t:\t%s\n", filters.buildingAge.button.Text, filters.buildingAge.value)
		f += fmt.Sprintf("%s\t:\t%s\n", filters.floor.button.Text, filters.floor.value)
		f += fmt.Sprintf("%s\t:\t%s\n", filters.storage.button.Text, filters.storage.value)
		f += fmt.Sprintf("%s\t:\t%s\n", filters.elevator.button.Text, filters.elevator.value)
		f += fmt.Sprintf("%s\t:\t%s\n", filters.adDate.button.Text, filters.adDate.value)
		f += fmt.Sprintf("%s\t:\t%s\n", filters.location.button.Text, filters.location.value)

		return f
	}

	// Buttons Handlers
	b.bot.Handle(&filters.price.button, func(ctx telebot.Context) error {
		b.bot.Handle(telebot.OnText, func(ctx telebot.Context) error {
			filters.price.value = ctx.Text()
			return ctx.EditOrSend(printFilter(), selector)
		})
		return ctx.EditOrSend("قیمت را بصورت انگلیسی درج کنید")
	})
	b.bot.Handle(&filters.area.button, func(ctx telebot.Context) error {
		b.bot.Handle(telebot.OnText, func(ctx telebot.Context) error {
			filters.area.value = ctx.Text()
			return ctx.EditOrSend(printFilter(), selector)
		})
		return ctx.EditOrSend("مساحت را بصورت انگلیسی درج کنید")
	})
	b.bot.Handle(&filters.rooms.button, func(ctx telebot.Context) error {
		b.bot.Handle(telebot.OnText, func(ctx telebot.Context) error {
			filters.rooms.value = ctx.Text()
			return ctx.EditOrSend(printFilter(), selector)
		})
		return ctx.EditOrSend("تعداد اتاق های را بصورت انگلیسی درج کنید")
	})
	b.bot.Handle(&filters.propertyType.button, func(ctx telebot.Context) error {
		b.bot.Handle(telebot.OnText, func(ctx telebot.Context) error {
			filters.propertyType.value = ctx.Text()
			return ctx.EditOrSend(printFilter(), selector)
		})
		return ctx.EditOrSend("نوع ساختمان را به فارسی بنویسید(آپارتمان | ویلایی)")
	})
	b.bot.Handle(&filters.buildingAge.button, func(ctx telebot.Context) error {
		b.bot.Handle(telebot.OnText, func(ctx telebot.Context) error {
			filters.buildingAge.value = ctx.Text()
			return ctx.EditOrSend(printFilter(), selector)
		})
		return ctx.EditOrSend("سن بنا را بصورت انگلیسی درج کنید")
	})
	b.bot.Handle(&filters.floor.button, func(ctx telebot.Context) error {
		b.bot.Handle(telebot.OnText, func(ctx telebot.Context) error {
			filters.floor.value = ctx.Text()
			return ctx.EditOrSend(printFilter(), selector)
		})
		return ctx.EditOrSend("طبقه مد نظر را بصورت انگلیسی درج کنید")
	})
	b.bot.Handle(&filters.storage.button, func(ctx telebot.Context) error {
		b.bot.Handle(telebot.OnText, func(ctx telebot.Context) error {
			filters.storage.value = ctx.Text()
			return ctx.EditOrSend(printFilter(), selector)
		})
		return ctx.EditOrSend("آیا انبار داشته باشد؟")
	})
	b.bot.Handle(&filters.elevator.button, func(ctx telebot.Context) error {
		b.bot.Handle(telebot.OnText, func(ctx telebot.Context) error {
			filters.elevator.value = ctx.Text()
			return ctx.EditOrSend(printFilter(), selector)
		})
		return ctx.EditOrSend("آیا آسانسور داشته باشد")
	})
	b.bot.Handle(&filters.adDate.button, func(ctx telebot.Context) error {
		b.bot.Handle(telebot.OnText, func(ctx telebot.Context) error {
			filters.adDate.value = ctx.Text()
			return ctx.EditOrSend(printFilter(), selector)
		})
		return ctx.EditOrSend("تاریخ درج آگهی را انگلیسی بنویسید")
	})
	b.bot.Handle(&filters.location.button, func(ctx telebot.Context) error {
		b.bot.Handle(telebot.OnText, func(ctx telebot.Context) error {
			filters.location.value = ctx.Text()
			return ctx.EditOrSend(printFilter(), selector)
		})
		return ctx.EditOrSend("محله را بنویسید")
	})

	return func(ctx telebot.Context) error {

		return ctx.EditOrSend(printFilter(), selector)
	}
}
