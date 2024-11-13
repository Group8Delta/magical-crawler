package bot

import (
	"fmt"
	"magical-crwler/constants"
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

func (f *Filters) removeAllValue() {
	f.price.value = ""
	f.area.value = ""
	f.rooms.value = ""
	f.propertyType.value = ""
	f.buildingAge.value = ""
	f.floor.value = ""
	f.storage.value = ""
	f.elevator.value = ""
	f.adDate.value = ""
	f.location.value = ""
}

func (f *Filters) Message() (msg string) {
	msg += fmt.Sprintf("%s\t:\t%s\n", f.price.button.Text, f.price.value)
	msg += fmt.Sprintf("%s\t:\t%s\n", f.area.button.Text, f.area.value)
	msg += fmt.Sprintf("%s\t:\t%s\n", f.rooms.button.Text, f.rooms.value)
	msg += fmt.Sprintf("%s\t:\t%s\n", f.propertyType.button.Text, f.propertyType.value)
	msg += fmt.Sprintf("%s\t:\t%s\n", f.buildingAge.button.Text, f.buildingAge.value)
	msg += fmt.Sprintf("%s\t:\t%s\n", f.floor.button.Text, f.floor.value)
	msg += fmt.Sprintf("%s\t:\t%s\n", f.storage.button.Text, f.storage.value)
	msg += fmt.Sprintf("%s\t:\t%s\n", f.elevator.button.Text, f.elevator.value)
	msg += fmt.Sprintf("%s\t:\t%s\n", f.adDate.button.Text, f.adDate.value)
	msg += fmt.Sprintf("%s\t:\t%s\n", f.location.button.Text, f.location.value)

	return
}

func (f *Filters) startSearch() []interface{} {
	return nil
}

func FilterHandlers(b *Bot) func(ctx telebot.Context) error {

	var (
		menuSelector     = &telebot.ReplyMarkup{RemoveKeyboard: true}
		priceSelector    = &telebot.ReplyMarkup{RemoveKeyboard: true}
		areaSelector     = &telebot.ReplyMarkup{RemoveKeyboard: true}
		roomSelector     = &telebot.ReplyMarkup{RemoveKeyboard: true}
		propertySelector = &telebot.ReplyMarkup{RemoveKeyboard: true}
		ageSelector      = &telebot.ReplyMarkup{RemoveKeyboard: true}
		floorSelector    = &telebot.ReplyMarkup{RemoveKeyboard: true}
		ynSelector       = &telebot.ReplyMarkup{RemoveKeyboard: true}
		adDateSelector   = &telebot.ReplyMarkup{RemoveKeyboard: true}

		YNButtons = []telebot.Btn{
			ynSelector.Data(constants.Yes, "YesNo", constants.Yes, "1"),
			ynSelector.Data(constants.No, "YesNo", constants.No, "0"),
			ynSelector.Data(constants.Unknown, "YesNo", constants.Unknown, "-1"),
		}

		goBtn     = menuSelector.Data(constants.GoButton, "Filters", "Search")
		removeBtn = menuSelector.Data(constants.RemoveButton, "Filters", "Remove")

		filters = Filters{
			price: FilterValue{button: menuSelector.Data(constants.PriceFilter, "Filters", "Price"),
				value: "",
				subButton: []telebot.Btn{
					priceSelector.Data(constants.PriceUnder500M, "PriceRange", "0", "500"),
					priceSelector.Data(constants.Price500MTo700M, "PriceRange", "500", "700"),
					priceSelector.Data(constants.Price700MTo900M, "PriceRange", "700", "900"),
					priceSelector.Data(constants.Price900MTo1B, "PriceRange", "900", "1000"),
					priceSelector.Data(constants.Price1BTo1_5B, "PriceRange", "1000", "1500"),
					priceSelector.Data(constants.Price1_5BTo2B, "PriceRange", "1500", "2000"),
					priceSelector.Data(constants.Price2BTo3B, "PriceRange", "2000", "3000"),
					priceSelector.Data(constants.Price3BTo4B, "PriceRange", "3000", "4000"),
					priceSelector.Data(constants.Price4BTo5B, "PriceRange", "4000", "5000"),
					priceSelector.Data(constants.Price5BTo7B, "PriceRange", "5000", "7000"),
					priceSelector.Data(constants.Price7BTo10B, "PriceRange", "7000", "10000"),
					priceSelector.Data(constants.Price10BTo15B, "PriceRange", "10000", "15000"),
					priceSelector.Data(constants.Price15BTo20B, "PriceRange", "15000", "20000"),
					priceSelector.Data(constants.Price20BTo30B, "PriceRange", "20000", "30000"),
					priceSelector.Data(constants.Price30BTo40B, "PriceRange", "30000", "40000"),
					priceSelector.Data(constants.Price40BTo50B, "PriceRange", "40000", "50000"),
					priceSelector.Data(constants.Price50BTo75B, "PriceRange", "50000", "70000"),
					priceSelector.Data(constants.Price75BTo100B, "PriceRange", "70000", "90000"),
					priceSelector.Data(constants.Price100BTo200B, "PriceRange", "100000", "200000"),
					priceSelector.Data(constants.Price200BTo300B, "PriceRange", "200000", "300000"),
					priceSelector.Data(constants.Price300BTo500B, "PriceRange", "300000", "500000"),
					priceSelector.Data(constants.Price500BTo700B, "PriceRange", "500000", "700000"),
					priceSelector.Data(constants.Price700BTo900B, "PriceRange", "700000", "900000"),
					priceSelector.Data(constants.PriceOver900B, "PriceRange", "900000+"),
				},
			},
			area: FilterValue{button: menuSelector.Data(constants.AreaFilter, "Filters", "Area"),
				value: "",
				subButton: []telebot.Btn{
					areaSelector.Data(constants.AreaUnder50, "AreaRange", "0", "50"),
					areaSelector.Data(constants.Area50To75, "AreaRange", "50", "75"),
					areaSelector.Data(constants.Area75To100, "AreaRange", "75", "100"),
					areaSelector.Data(constants.Area100To150, "AreaRange", "100", "150"),

					areaSelector.Data(constants.Area150To200, "AreaRange", "150", "200"),
					areaSelector.Data(constants.Area200To250, "AreaRange", "200", "250"),
					areaSelector.Data(constants.Area250To300, "AreaRange", "250", "300"),
					areaSelector.Data(constants.Area300To400, "AreaRange", "300", "400"),
					areaSelector.Data(constants.Area400To500, "AreaRange", "400", "500"),
					areaSelector.Data(constants.Area500To750, "AreaRange", "500", "750"),
					areaSelector.Data(constants.Area750To1000, "AreaRange", "750", "1000"),

					areaSelector.Data(constants.Area1000To1500, "AreaRange", "1000", "1500"),
					areaSelector.Data(constants.Area1500To2000, "AreaRange", "1500", "2000"),
					areaSelector.Data(constants.Area2000To3000, "AreaRange", "2000", "3000"),
					areaSelector.Data(constants.Area3000To4000, "AreaRange", "3000", "4000"),
					areaSelector.Data(constants.Area4000To5000, "AreaRange", "4000", "5000"),
					areaSelector.Data(constants.Area5000To7500, "AreaRange", "5000", "7500"),
					areaSelector.Data(constants.Area7500To10000, "AreaRange", "7500", "10000"),
					areaSelector.Data(constants.AreaOver10000, "AreaRange", "1000+"),
				},
			},
			rooms: FilterValue{button: menuSelector.Data(constants.RoomsFilter, "Filters", "Rooms"),
				value: "",
				subButton: []telebot.Btn{
					roomSelector.Data(constants.Bedrooms0, "NumberOfRooms", constants.Bedrooms0, "0"),
					roomSelector.Data(constants.Bedrooms1, "NumberOfRooms", constants.Bedrooms1, "1"),
					roomSelector.Data(constants.Bedrooms2, "NumberOfRooms", constants.Bedrooms2, "2"),
					roomSelector.Data(constants.Bedrooms3, "NumberOfRooms", constants.Bedrooms3, "3"),
					roomSelector.Data(constants.Bedrooms4, "NumberOfRooms", constants.Bedrooms4, "4"),
					roomSelector.Data(constants.Bedrooms5, "NumberOfRooms", constants.Bedrooms5, "5"),
					roomSelector.Data(constants.Bedrooms6, "NumberOfRooms", constants.Bedrooms6, "6"),
					roomSelector.Data(constants.Bedrooms7, "NumberOfRooms", constants.Bedrooms7, "7"),
					roomSelector.Data(constants.Bedrooms8, "NumberOfRooms", constants.Bedrooms8, "8"),
					roomSelector.Data(constants.Bedrooms9, "NumberOfRooms", constants.Bedrooms9, "9"),
					roomSelector.Data(constants.Bedrooms10, "NumberOfRooms", constants.Bedrooms10, "10"),
					roomSelector.Data(constants.BedroomsOver10, "NumberOfRooms", constants.BedroomsOver10, "10+"),
				},
			},
			propertyType: FilterValue{button: menuSelector.Data(constants.PropertyTypeFilter, "Filters", "PropertyType"),
				value: "",
				subButton: []telebot.Btn{
					propertySelector.Data(constants.PropertyApartment, "Property", constants.PropertyApartment),
					propertySelector.Data(constants.PropertyVilla, "Property", constants.PropertyVilla),
					propertySelector.Data(constants.PropertyCommercial, "Property", constants.PropertyCommercial),
					propertySelector.Data(constants.PropertyOffice, "Property", constants.PropertyOffice),
					propertySelector.Data(constants.PropertyLand, "Property", constants.PropertyLand),
				},
			},
			buildingAge: FilterValue{button: menuSelector.Data(constants.BuildingAgeFilter, "Filters", "BuildingAge"),
				value: "",
				subButton: []telebot.Btn{
					ageSelector.Data(constants.BuildingAgeNew, "Age", constants.BuildingAgeNew),
					ageSelector.Data(constants.BuildingAge1Year, "Age", constants.BuildingAge1Year),
					ageSelector.Data(constants.BuildingAge2Years, "Age", constants.BuildingAge2Years),
					ageSelector.Data(constants.BuildingAge3Years, "Age", constants.BuildingAge3Years),
					ageSelector.Data(constants.BuildingAge4Years, "Age", constants.BuildingAge4Years),
					ageSelector.Data(constants.BuildingAge5Years, "Age", constants.BuildingAge5Years),
					ageSelector.Data(constants.BuildingAge6Years, "Age", constants.BuildingAge6Years),
					ageSelector.Data(constants.BuildingAge7Years, "Age", constants.BuildingAge7Years),
					ageSelector.Data(constants.BuildingAge8Years, "Age", constants.BuildingAge8Years),
					ageSelector.Data(constants.BuildingAge9Years, "Age", constants.BuildingAge9Years),
					ageSelector.Data(constants.BuildingAge10Years, "Age", constants.BuildingAge10Years),
					ageSelector.Data(constants.BuildingAgeOver10, "Age", constants.BuildingAgeOver10),
				},
			},
			floor: FilterValue{button: menuSelector.Data(constants.FloorFilter, "Filters", "Floor"),
				value: "",
				subButton: []telebot.Btn{
					floorSelector.Data(constants.Floor0, "Floor", constants.Floor0),
					floorSelector.Data(constants.Floor1, "Floor", constants.Floor1),
					floorSelector.Data(constants.Floor2, "Floor", constants.Floor2),
					floorSelector.Data(constants.Floor3, "Floor", constants.Floor3),
					floorSelector.Data(constants.Floor4, "Floor", constants.Floor4),
					floorSelector.Data(constants.Floor5, "Floor", constants.Floor5),
					floorSelector.Data(constants.Floor6, "Floor", constants.Floor6),
					floorSelector.Data(constants.Floor7, "Floor", constants.Floor7),
					floorSelector.Data(constants.Floor8, "Floor", constants.Floor8),
					floorSelector.Data(constants.Floor9, "Floor", constants.Floor9),
					floorSelector.Data(constants.Floor10, "Floor", constants.Floor10),
					floorSelector.Data(constants.FloorOver10, "Floor", constants.FloorOver10),
				},
			},
			storage: FilterValue{button: menuSelector.Data(constants.StorageFilter, "Filters", "Storage"),
				value:     "",
				subButton: YNButtons,
			},
			elevator: FilterValue{button: menuSelector.Data(constants.ElevatorFilter, "Filters", "Elevator"),
				value:     "",
				subButton: YNButtons,
			},
			adDate: FilterValue{button: menuSelector.Data(constants.AdDateFilter, "Filters", "AdDate"),
				value: "",
				subButton: []telebot.Btn{
					adDateSelector.Data(constants.TimeToday, "Time"),
					adDateSelector.Data(constants.Time1DayAgo, "Time"),
					adDateSelector.Data(constants.Time2DaysAgo, "Time"),
					adDateSelector.Data(constants.Time3DaysAgo, "Time"),
					adDateSelector.Data(constants.Time1WeekAgo, "Time"),
					adDateSelector.Data(constants.Time1MonthAgo, "Time"),
					adDateSelector.Data(constants.Time1YearAgo, "Time"),
				},
			},
			location: FilterValue{button: menuSelector.Data(constants.LocationFilter, "Filters", "Location"),
				value: "",
			},
		}
	)

	menuSelector.Inline(
		menuSelector.Row(filters.price.button),
		menuSelector.Row(filters.area.button, filters.rooms.button, filters.propertyType.button),
		menuSelector.Row(filters.buildingAge.button, filters.floor.button, filters.storage.button),
		menuSelector.Row(filters.elevator.button, filters.adDate.button, filters.location.button),
		menuSelector.Row(goBtn, removeBtn),
	)

	//Inline Styles
	priceSelector.Inline(menuSelector.Split(1, filters.price.subButton)...)
	areaSelector.Inline(menuSelector.Split(1, filters.area.subButton)...)
	roomSelector.Inline(menuSelector.Split(3, filters.rooms.subButton)...)
	propertySelector.Inline(menuSelector.Split(3, filters.propertyType.subButton)...)
	ageSelector.Inline(menuSelector.Split(3, filters.buildingAge.subButton)...)
	floorSelector.Inline(menuSelector.Split(3, filters.floor.subButton)...)
	ynSelector.Inline(menuSelector.Split(3, YNButtons)...)
	adDateSelector.Inline(menuSelector.Split(2, filters.adDate.subButton)...)

	// Buttons Handlers
	b.Bot.Handle(&telebot.InlineButton{Unique: "Filters"}, func(ctx telebot.Context) error {
		filterSelected := ctx.Data()
		switch filterSelected {
		case "Price":
			b.Bot.Handle(&telebot.InlineButton{Unique: "PriceRange"}, func(c telebot.Context) error {
				filters.price.value = c.Data()
				return c.EditOrSend(filters.Message(), menuSelector)
			})
			return ctx.EditOrSend(filters.Message(), priceSelector)
		case "Area":
			b.Bot.Handle(&telebot.InlineButton{Unique: "AreaRange"}, func(c telebot.Context) error {
				filters.area.value = c.Data()
				return c.EditOrSend(filters.Message(), menuSelector)
			})
			return ctx.EditOrSend(filters.Message(), areaSelector)
		case "Rooms":
			b.Bot.Handle(&telebot.InlineButton{Unique: "NumberOfRooms"}, func(c telebot.Context) error {
				filters.rooms.value = strings.Split(c.Data(), "|")[0]
				return c.EditOrSend(filters.Message(), menuSelector)
			})
			return ctx.EditOrSend(filters.Message(), roomSelector)
		case "PropertyType":
			b.Bot.Handle(&telebot.InlineButton{Unique: "Property"}, func(c telebot.Context) error {
				filters.propertyType.value = c.Data()
				return c.EditOrSend(filters.Message(), menuSelector)
			})
			return ctx.EditOrSend(filters.Message(), propertySelector)
		case "BuildingAge":
			b.Bot.Handle(&telebot.InlineButton{Unique: "Age"}, func(c telebot.Context) error {
				filters.buildingAge.value = c.Data()
				return c.EditOrSend(filters.Message(), menuSelector)
			})
			return ctx.EditOrSend(filters.Message(), ageSelector)
		case "Floor":
			b.Bot.Handle(&telebot.InlineButton{Unique: "Floor"}, func(c telebot.Context) error {
				filters.floor.value = c.Data()
				return c.EditOrSend(filters.Message(), menuSelector)
			})
			return ctx.EditOrSend(filters.Message(), floorSelector)
		case "Storage":
			b.Bot.Handle(&telebot.InlineButton{Unique: "YesNo"}, func(c telebot.Context) error {
				filters.storage.value = strings.Split(c.Data(), "|")[0]
				return c.EditOrSend(filters.Message(), menuSelector)
			})
			return ctx.EditOrSend(filters.Message(), ynSelector)
		case "Elevator":
			b.Bot.Handle(&telebot.InlineButton{Unique: "YesNo"}, func(c telebot.Context) error {
				filters.elevator.value = strings.Split(c.Data(), "|")[0]
				return c.EditOrSend(filters.Message(), menuSelector)
			})
			return ctx.EditOrSend(filters.Message(), ynSelector)
		case "AdDate":
			b.Bot.Handle(&telebot.InlineButton{Unique: "Time"}, func(c telebot.Context) error {
				filters.adDate.value = c.Data()
				return c.EditOrSend(filters.Message(), menuSelector)
			})
			return ctx.EditOrSend(filters.Message(), adDateSelector)
		case "Location":
			b.Bot.Handle(telebot.OnText, func(c telebot.Context) error {

				return c.EditOrSend(filters.Message(), menuSelector)
			})
			return ctx.EditOrSend(filters.Message(), priceSelector)
		case "Search":
			ctx.Send(constants.Loading)
			// ads := filters.startSearch()
			return ctx.Send(constants.SearchMsg)
		case "Remove":
			filters.removeAllValue()
			return ctx.EditOrSend(filters.Message(), menuSelector)
		default:
			filters.removeAllValue()
			return ctx.EditOrSend("/menu")
		}

	})

	return func(ctx telebot.Context) error {

		return ctx.EditOrSend(filters.Message(), menuSelector)
	}
}
