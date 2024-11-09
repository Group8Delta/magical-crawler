package bot

import (
	"fmt"
	"magical-crwler/config"
	"strings"

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
	value     string
	button    telebot.Btn
	subButton []telebot.Btn
}

func FilterHandlers(b *Bot) func(ctx telebot.Context) error {

	var (
		selector         = &telebot.ReplyMarkup{RemoveKeyboard: true}
		priceSelector    = &telebot.ReplyMarkup{RemoveKeyboard: true}
		areaSelector     = &telebot.ReplyMarkup{RemoveKeyboard: true}
		roomSelector     = &telebot.ReplyMarkup{RemoveKeyboard: true}
		propertySelector = &telebot.ReplyMarkup{RemoveKeyboard: true}
		ageSelector      = &telebot.ReplyMarkup{RemoveKeyboard: true}
		floorSelector    = &telebot.ReplyMarkup{RemoveKeyboard: true}
		YNSelector       = &telebot.ReplyMarkup{RemoveKeyboard: true}

		YNButtons = []telebot.Btn{
			YNSelector.Data(config.Yes, "YesNo", config.Yes, "1"),
			YNSelector.Data(config.No, "YesNo", config.No, "0"),
			YNSelector.Data(config.Unknown, "YesNo", config.Unknown, "-1"),
		}

		filters = Filters{
			price: FilterValue{button: selector.Data(config.PriceFilter, "Filters", "Price"),
				value: "",
				subButton: []telebot.Btn{
					priceSelector.Data(config.PriceUnder500M, "PriceRange", "0", "500"),
					priceSelector.Data(config.Price500MTo700M, "PriceRange", "500", "700"),
					priceSelector.Data(config.Price700MTo900M, "PriceRange", "700", "900"),
					priceSelector.Data(config.Price900MTo1B, "PriceRange", "900", "1000"),
					priceSelector.Data(config.Price1BTo1_5B, "PriceRange", "1000", "1500"),
					priceSelector.Data(config.Price1_5BTo2B, "PriceRange", "1500", "2000"),
					priceSelector.Data(config.Price2BTo3B, "PriceRange", "2000", "3000"),
					priceSelector.Data(config.Price3BTo4B, "PriceRange", "3000", "4000"),
					priceSelector.Data(config.Price4BTo5B, "PriceRange", "4000", "5000"),
					priceSelector.Data(config.Price5BTo7B, "PriceRange", "5000", "7000"),
					priceSelector.Data(config.Price7BTo10B, "PriceRange", "7000", "10000"),
					priceSelector.Data(config.Price10BTo15B, "PriceRange", "10000", "15000"),
					priceSelector.Data(config.Price15BTo20B, "PriceRange", "15000", "20000"),
					priceSelector.Data(config.Price20BTo30B, "PriceRange", "20000", "30000"),
					priceSelector.Data(config.Price30BTo40B, "PriceRange", "30000", "40000"),
					priceSelector.Data(config.Price40BTo50B, "PriceRange", "40000", "50000"),
					priceSelector.Data(config.Price50BTo75B, "PriceRange", "50000", "70000"),
					priceSelector.Data(config.Price75BTo100B, "PriceRange", "70000", "90000"),
					priceSelector.Data(config.Price100BTo200B, "PriceRange", "100000", "200000"),
					priceSelector.Data(config.Price200BTo300B, "PriceRange", "200000", "300000"),
					priceSelector.Data(config.Price300BTo500B, "PriceRange", "300000", "500000"),
					priceSelector.Data(config.Price500BTo700B, "PriceRange", "500000", "700000"),
					priceSelector.Data(config.Price700BTo900B, "PriceRange", "700000", "900000"),
					priceSelector.Data(config.PriceOver900B, "PriceRange", "900000+"),
				},
			},
			area: FilterValue{button: selector.Data(config.AreaFilter, "Filters", "Area"),
				value: "",
				subButton: []telebot.Btn{
					areaSelector.Data(config.AreaUnder50, "AreaRange", "0", "50"),
					areaSelector.Data(config.Area50To75, "AreaRange", "50", "75"),
					areaSelector.Data(config.Area75To100, "AreaRange", "75", "100"),
					areaSelector.Data(config.Area100To150, "AreaRange", "100", "150"),

					areaSelector.Data(config.Area150To200, "AreaRange", "150", "200"),
					areaSelector.Data(config.Area200To250, "AreaRange", "200", "250"),
					areaSelector.Data(config.Area250To300, "AreaRange", "250", "300"),
					areaSelector.Data(config.Area300To400, "AreaRange", "300", "400"),
					areaSelector.Data(config.Area400To500, "AreaRange", "400", "500"),
					areaSelector.Data(config.Area500To750, "AreaRange", "500", "750"),
					areaSelector.Data(config.Area750To1000, "AreaRange", "750", "1000"),

					areaSelector.Data(config.Area1000To1500, "AreaRange", "1000", "1500"),
					areaSelector.Data(config.Area1500To2000, "AreaRange", "1500", "2000"),
					areaSelector.Data(config.Area2000To3000, "AreaRange", "2000", "3000"),
					areaSelector.Data(config.Area3000To4000, "AreaRange", "3000", "4000"),
					areaSelector.Data(config.Area4000To5000, "AreaRange", "4000", "5000"),
					areaSelector.Data(config.Area5000To7500, "AreaRange", "5000", "7500"),
					areaSelector.Data(config.Area7500To10000, "AreaRange", "7500", "10000"),
					areaSelector.Data(config.AreaOver10000, "AreaRange", "1000+"),
				},
			},
			rooms: FilterValue{button: selector.Data(config.RoomsFilter, "Filters", "Rooms"),
				value: "",
				subButton: []telebot.Btn{
					roomSelector.Data(config.Bedrooms0, "NumberOfRooms", config.Bedrooms0, "0"),
					roomSelector.Data(config.Bedrooms1, "NumberOfRooms", config.Bedrooms1, "1"),
					roomSelector.Data(config.Bedrooms2, "NumberOfRooms", config.Bedrooms2, "2"),
					roomSelector.Data(config.Bedrooms3, "NumberOfRooms", config.Bedrooms3, "3"),
					roomSelector.Data(config.Bedrooms4, "NumberOfRooms", config.Bedrooms4, "4"),
					roomSelector.Data(config.Bedrooms5, "NumberOfRooms", config.Bedrooms5, "5"),
					roomSelector.Data(config.Bedrooms6, "NumberOfRooms", config.Bedrooms6, "6"),
					roomSelector.Data(config.Bedrooms7, "NumberOfRooms", config.Bedrooms7, "7"),
					roomSelector.Data(config.Bedrooms8, "NumberOfRooms", config.Bedrooms8, "8"),
					roomSelector.Data(config.Bedrooms9, "NumberOfRooms", config.Bedrooms9, "9"),
					roomSelector.Data(config.Bedrooms10, "NumberOfRooms", config.Bedrooms10, "10"),
					roomSelector.Data(config.BedroomsOver10, "NumberOfRooms", config.BedroomsOver10, "10+"),
				},
			},
			propertyType: FilterValue{button: selector.Data(config.PropertyTypeFilter, "Filters", "PropertyType"),
				value: "",
				subButton: []telebot.Btn{
					propertySelector.Data(config.PropertyApartment, "Property", config.PropertyApartment),
					propertySelector.Data(config.PropertyVilla, "Property", config.PropertyVilla),
					propertySelector.Data(config.PropertyCommercial, "Property", config.PropertyCommercial),
					propertySelector.Data(config.PropertyOffice, "Property", config.PropertyOffice),
					propertySelector.Data(config.PropertyLand, "Property", config.PropertyLand),
				},
			},
			buildingAge: FilterValue{button: selector.Data(config.BuildingAgeFilter, "Filters", "BuildingAge"),
				value: "",
				subButton: []telebot.Btn{
					ageSelector.Data(config.BuildingAgeNew, "Age", config.BuildingAgeNew),
					ageSelector.Data(config.BuildingAge1Year, "Age", config.BuildingAge1Year),
					ageSelector.Data(config.BuildingAge2Years, "Age", config.BuildingAge2Years),
					ageSelector.Data(config.BuildingAge3Years, "Age", config.BuildingAge3Years),
					ageSelector.Data(config.BuildingAge4Years, "Age", config.BuildingAge4Years),
					ageSelector.Data(config.BuildingAge5Years, "Age", config.BuildingAge5Years),
					ageSelector.Data(config.BuildingAge6Years, "Age", config.BuildingAge6Years),
					ageSelector.Data(config.BuildingAge7Years, "Age", config.BuildingAge7Years),
					ageSelector.Data(config.BuildingAge8Years, "Age", config.BuildingAge8Years),
					ageSelector.Data(config.BuildingAge9Years, "Age", config.BuildingAge9Years),
					ageSelector.Data(config.BuildingAge10Years, "Age", config.BuildingAge10Years),
					ageSelector.Data(config.BuildingAgeOver10, "Age", config.BuildingAgeOver10),
				},
			},
			floor: FilterValue{button: selector.Data(config.FloorFilter, "Filters", "Floor"),
				value: "",
				subButton: []telebot.Btn{
					floorSelector.Data(config.Floor0, "Floor", config.Floor0),
					floorSelector.Data(config.Floor1, "Floor", config.Floor1),
					floorSelector.Data(config.Floor2, "Floor", config.Floor2),
					floorSelector.Data(config.Floor3, "Floor", config.Floor3),
					floorSelector.Data(config.Floor4, "Floor", config.Floor4),
					floorSelector.Data(config.Floor5, "Floor", config.Floor5),
					floorSelector.Data(config.Floor6, "Floor", config.Floor6),
					floorSelector.Data(config.Floor7, "Floor", config.Floor7),
					floorSelector.Data(config.Floor8, "Floor", config.Floor8),
					floorSelector.Data(config.Floor9, "Floor", config.Floor9),
					floorSelector.Data(config.Floor10, "Floor", config.Floor10),
					floorSelector.Data(config.FloorOver10, "Floor", config.FloorOver10),
				},
			},
			storage: FilterValue{button: selector.Data(config.StorageFilter, "Filters", "Storage"),
				value:     "",
				subButton: YNButtons,
			},
			elevator: FilterValue{button: selector.Data(config.ElevatorFilter, "Filters", "Elevator"),
				value:     "",
				subButton: YNButtons,
			},
			adDate: FilterValue{button: selector.Data(config.AdDateFilter, "Filters", "AdDate"),
				value: "",
			},
			location: FilterValue{button: selector.Data(config.LocationFilter, "Filters", "Location"),
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

	priceSelector.Inline(selector.Split(1, filters.price.subButton)...)
	areaSelector.Inline(selector.Split(1, filters.area.subButton)...)
	roomSelector.Inline(selector.Split(3, filters.rooms.subButton)...)
	propertySelector.Inline(selector.Split(3, filters.propertyType.subButton)...)
	ageSelector.Inline(selector.Split(3, filters.buildingAge.subButton)...)
	floorSelector.Inline(selector.Split(3, filters.floor.subButton)...)
	YNSelector.Inline(selector.Split(3, YNButtons)...)

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
	b.bot.Handle(&telebot.InlineButton{Unique: "Filters"}, func(ctx telebot.Context) error {
		filterSelected := ctx.Data()
		switch filterSelected {
		case "Price":
			b.bot.Handle(&telebot.InlineButton{Unique: "PriceRange"}, func(c telebot.Context) error {
				filters.price.value = c.Data()
				return c.EditOrSend(printFilter(), selector)
			})
			return ctx.EditOrSend(printFilter(), priceSelector)
		case "Area":
			b.bot.Handle(&telebot.InlineButton{Unique: "AreaRange"}, func(c telebot.Context) error {
				filters.area.value = c.Data()
				return c.EditOrSend(printFilter(), selector)
			})
			return ctx.EditOrSend(printFilter(), areaSelector)
		case "Rooms":
			b.bot.Handle(&telebot.InlineButton{Unique: "NumberOfRooms"}, func(c telebot.Context) error {
				filters.rooms.value = strings.Split(c.Data(), "|")[0]
				return c.EditOrSend(printFilter(), selector)
			})
			return ctx.EditOrSend(printFilter(), roomSelector)
		case "PropertyType":
			b.bot.Handle(&telebot.InlineButton{Unique: "Property"}, func(c telebot.Context) error {
				filters.propertyType.value = c.Data()
				return c.EditOrSend(printFilter(), selector)
			})
			return ctx.EditOrSend(printFilter(), propertySelector)
		case "BuildingAge":
			b.bot.Handle(&telebot.InlineButton{Unique: "Age"}, func(c telebot.Context) error {
				filters.buildingAge.value = c.Data()
				return c.EditOrSend(printFilter(), selector)
			})
			return ctx.EditOrSend(printFilter(), ageSelector)
		case "Floor":
			b.bot.Handle(&telebot.InlineButton{Unique: "Floor"}, func(c telebot.Context) error {
				filters.floor.value = c.Data()
				return c.EditOrSend(printFilter(), selector)
			})
			return ctx.EditOrSend(printFilter(), floorSelector)
		case "Storage":
			b.bot.Handle(&telebot.InlineButton{Unique: "YesNo"}, func(c telebot.Context) error {
				filters.storage.value = strings.Split(c.Data(), "|")[0]
				return c.EditOrSend(printFilter(), selector)
			})
			return ctx.EditOrSend(printFilter(), YNSelector)
		case "Elevator":
			b.bot.Handle(&telebot.InlineButton{Unique: "YesNo"}, func(c telebot.Context) error {
				filters.elevator.value = strings.Split(c.Data(), "|")[0]
				return c.EditOrSend(printFilter(), selector)
			})
			return ctx.EditOrSend(printFilter(), YNSelector)
		case "AdDate":
			b.bot.Handle(&telebot.InlineButton{Unique: "PriceRange"}, func(c telebot.Context) error {
				filters.adDate.value = c.Data()
				return c.EditOrSend(printFilter(), selector)
			})
			return ctx.EditOrSend(printFilter(), priceSelector)
		case "Location":
			b.bot.Handle(&telebot.InlineButton{Unique: "PriceRange"}, func(c telebot.Context) error {
				filters.location.value = c.Data()
				return c.EditOrSend(printFilter(), selector)
			})
			return ctx.EditOrSend(printFilter(), priceSelector)
		default:
			return ctx.EditOrSend("/menu")
		}

	})

	return func(ctx telebot.Context) error {

		return ctx.EditOrSend(printFilter(), selector)
	}
}
