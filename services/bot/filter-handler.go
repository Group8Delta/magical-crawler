package bot

// TODO: change name to search-handler

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

// TODO: return type must change to Ads models
func (f *Filters) startSearch() []interface{} {

	return nil
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

func (f *Filters) message() string {
	var msg string
	msg += fmt.Sprintf("💰%s\t:\n%s\n\n", f.price.button.Text, f.price.value)
	msg += fmt.Sprintf("📏%s\t:\n%s\n\n", f.area.button.Text, f.area.value)
	msg += fmt.Sprintf("🛏️%s\t:\n%s\n\n", f.rooms.button.Text, f.rooms.value)
	msg += fmt.Sprintf("🏘️%s\t:\n%s\n\n", f.propertyType.button.Text, f.propertyType.value)
	msg += fmt.Sprintf("🏚️%s\t:\n%s\n\n", f.buildingAge.button.Text, f.buildingAge.value)
	msg += fmt.Sprintf("🔃%s\t:\n%s\n\n", f.floor.button.Text, f.floor.value)
	msg += fmt.Sprintf("🚪%s\t:\n%s\n\n", f.storage.button.Text, f.storage.value)
	msg += fmt.Sprintf("🛗%s\t:\n%s\n\n", f.elevator.button.Text, f.elevator.value)
	msg += fmt.Sprintf("🗓️%s\t:\n%s\n\n", f.adDate.button.Text, f.adDate.value)
	msg += fmt.Sprintf("📍%s\t:\n%s\n\n", f.location.button.Text, f.location.value)

	return msg
}

func newReplyMarkup() *telebot.ReplyMarkup {
	return &telebot.ReplyMarkup{RemoveKeyboard: true, ResizeKeyboard: true}
}

func FilterHandlers(b *Bot) func(ctx telebot.Context) error {

	var (
		selector = map[string]*telebot.ReplyMarkup{
			"menu":     newReplyMarkup(),
			"price":    newReplyMarkup(),
			"area":     newReplyMarkup(),
			"room":     newReplyMarkup(),
			"property": newReplyMarkup(),
			"age":      newReplyMarkup(),
			"floor":    newReplyMarkup(),
			"yes-no":   newReplyMarkup(),
			"ad-date":  newReplyMarkup(),
			"location": newReplyMarkup(),
		}

		YNButtons = []telebot.Btn{
			selector["yes-no"].Data(constants.Yes, "YesNo", constants.Yes, "1"),
			selector["yes-no"].Data(constants.No, "YesNo", constants.No, "0"),
			selector["yes-no"].Data(constants.Unknown, "YesNo", constants.Unknown, "-1"),
		}

		goBtn     = selector["menu"].Data(constants.GoButton, "Filters", "Search")
		removeBtn = selector["menu"].Data(constants.RemoveButton, "Filters", "Remove")

		filters = Filters{
			price: FilterValue{button: selector["menu"].Data(constants.PriceFilter, "Filters", "Price"),
				value: "",
				subButton: []telebot.Btn{
					selector["price"].Data(constants.PriceUnder500M, "PriceRange", "0", "500"),
					selector["price"].Data(constants.Price500MTo700M, "PriceRange", "500", "700"),
					selector["price"].Data(constants.Price700MTo900M, "PriceRange", "700", "900"),
					selector["price"].Data(constants.Price900MTo1B, "PriceRange", "900", "1000"),
					selector["price"].Data(constants.Price1BTo1_5B, "PriceRange", "1000", "1500"),
					selector["price"].Data(constants.Price1_5BTo2B, "PriceRange", "1500", "2000"),
					selector["price"].Data(constants.Price2BTo3B, "PriceRange", "2000", "3000"),
					selector["price"].Data(constants.Price3BTo4B, "PriceRange", "3000", "4000"),
					selector["price"].Data(constants.Price4BTo5B, "PriceRange", "4000", "5000"),
					selector["price"].Data(constants.Price5BTo7B, "PriceRange", "5000", "7000"),
					selector["price"].Data(constants.Price7BTo10B, "PriceRange", "7000", "10000"),
					selector["price"].Data(constants.Price10BTo15B, "PriceRange", "10000", "15000"),
					selector["price"].Data(constants.Price15BTo20B, "PriceRange", "15000", "20000"),
					selector["price"].Data(constants.Price20BTo30B, "PriceRange", "20000", "30000"),
					selector["price"].Data(constants.Price30BTo40B, "PriceRange", "30000", "40000"),
					selector["price"].Data(constants.Price40BTo50B, "PriceRange", "40000", "50000"),
					selector["price"].Data(constants.Price50BTo75B, "PriceRange", "50000", "70000"),
					selector["price"].Data(constants.Price75BTo100B, "PriceRange", "70000", "90000"),
					selector["price"].Data(constants.Price100BTo200B, "PriceRange", "100000", "200000"),
					selector["price"].Data(constants.Price200BTo300B, "PriceRange", "200000", "300000"),
					selector["price"].Data(constants.Price300BTo500B, "PriceRange", "300000", "500000"),
					selector["price"].Data(constants.Price500BTo700B, "PriceRange", "500000", "700000"),
					selector["price"].Data(constants.Price700BTo900B, "PriceRange", "700000", "900000"),
					selector["price"].Data(constants.PriceOver900B, "PriceRange", "900000+"),
				},
			},
			area: FilterValue{button: selector["menu"].Data(constants.AreaFilter, "Filters", "Area"),
				value: "",
				subButton: []telebot.Btn{
					selector["area"].Data(constants.AreaUnder50, "AreaRange", "0", "50"),
					selector["area"].Data(constants.Area50To75, "AreaRange", "50", "75"),
					selector["area"].Data(constants.Area75To100, "AreaRange", "75", "100"),
					selector["area"].Data(constants.Area100To150, "AreaRange", "100", "150"),

					selector["area"].Data(constants.Area150To200, "AreaRange", "150", "200"),
					selector["area"].Data(constants.Area200To250, "AreaRange", "200", "250"),
					selector["area"].Data(constants.Area250To300, "AreaRange", "250", "300"),
					selector["area"].Data(constants.Area300To400, "AreaRange", "300", "400"),
					selector["area"].Data(constants.Area400To500, "AreaRange", "400", "500"),
					selector["area"].Data(constants.Area500To750, "AreaRange", "500", "750"),
					selector["area"].Data(constants.Area750To1000, "AreaRange", "750", "1000"),

					selector["area"].Data(constants.Area1000To1500, "AreaRange", "1000", "1500"),
					selector["area"].Data(constants.Area1500To2000, "AreaRange", "1500", "2000"),
					selector["area"].Data(constants.Area2000To3000, "AreaRange", "2000", "3000"),
					selector["area"].Data(constants.Area3000To4000, "AreaRange", "3000", "4000"),
					selector["area"].Data(constants.Area4000To5000, "AreaRange", "4000", "5000"),
					selector["area"].Data(constants.Area5000To7500, "AreaRange", "5000", "7500"),
					selector["area"].Data(constants.Area7500To10000, "AreaRange", "7500", "10000"),
					selector["area"].Data(constants.AreaOver10000, "AreaRange", "1000+"),
				},
			},
			rooms: FilterValue{button: selector["menu"].Data(constants.RoomsFilter, "Filters", "Rooms"),
				value: "",
				subButton: []telebot.Btn{
					selector["room"].Data(constants.Bedrooms0, "NumberOfRooms", constants.Bedrooms0, "0"),
					selector["room"].Data(constants.Bedrooms1, "NumberOfRooms", constants.Bedrooms1, "1"),
					selector["room"].Data(constants.Bedrooms2, "NumberOfRooms", constants.Bedrooms2, "2"),
					selector["room"].Data(constants.Bedrooms3, "NumberOfRooms", constants.Bedrooms3, "3"),
					selector["room"].Data(constants.Bedrooms4, "NumberOfRooms", constants.Bedrooms4, "4"),
					selector["room"].Data(constants.Bedrooms5, "NumberOfRooms", constants.Bedrooms5, "5"),
					selector["room"].Data(constants.Bedrooms6, "NumberOfRooms", constants.Bedrooms6, "6"),
					selector["room"].Data(constants.Bedrooms7, "NumberOfRooms", constants.Bedrooms7, "7"),
					selector["room"].Data(constants.Bedrooms8, "NumberOfRooms", constants.Bedrooms8, "8"),
					selector["room"].Data(constants.Bedrooms9, "NumberOfRooms", constants.Bedrooms9, "9"),
					selector["room"].Data(constants.Bedrooms10, "NumberOfRooms", constants.Bedrooms10, "10"),
					selector["room"].Data(constants.BedroomsOver10, "NumberOfRooms", constants.BedroomsOver10, "10+"),
				},
			},
			propertyType: FilterValue{button: selector["menu"].Data(constants.PropertyTypeFilter, "Filters", "PropertyType"),
				value: "",
				subButton: []telebot.Btn{
					selector["property"].Data(constants.PropertyApartment, "Property", constants.PropertyApartment),
					selector["property"].Data(constants.PropertyVilla, "Property", constants.PropertyVilla),
					selector["property"].Data(constants.PropertyCommercial, "Property", constants.PropertyCommercial),
					selector["property"].Data(constants.PropertyOffice, "Property", constants.PropertyOffice),
					selector["property"].Data(constants.PropertyLand, "Property", constants.PropertyLand),
				},
			},
			buildingAge: FilterValue{button: selector["menu"].Data(constants.BuildingAgeFilter, "Filters", "BuildingAge"),
				value: "",
				subButton: []telebot.Btn{
					selector["age"].Data(constants.BuildingAgeNew, "Age", constants.BuildingAgeNew),
					selector["age"].Data(constants.BuildingAge1Year, "Age", constants.BuildingAge1Year),
					selector["age"].Data(constants.BuildingAge2Years, "Age", constants.BuildingAge2Years),
					selector["age"].Data(constants.BuildingAge3Years, "Age", constants.BuildingAge3Years),
					selector["age"].Data(constants.BuildingAge4Years, "Age", constants.BuildingAge4Years),
					selector["age"].Data(constants.BuildingAge5Years, "Age", constants.BuildingAge5Years),
					selector["age"].Data(constants.BuildingAge6Years, "Age", constants.BuildingAge6Years),
					selector["age"].Data(constants.BuildingAge7Years, "Age", constants.BuildingAge7Years),
					selector["age"].Data(constants.BuildingAge8Years, "Age", constants.BuildingAge8Years),
					selector["age"].Data(constants.BuildingAge9Years, "Age", constants.BuildingAge9Years),
					selector["age"].Data(constants.BuildingAge10Years, "Age", constants.BuildingAge10Years),
					selector["age"].Data(constants.BuildingAgeOver10, "Age", constants.BuildingAgeOver10),
				},
			},
			floor: FilterValue{button: selector["menu"].Data(constants.FloorFilter, "Filters", "Floor"),
				value: "",
				subButton: []telebot.Btn{
					selector["floor"].Data(constants.Floor0, "Floor", constants.Floor0),
					selector["floor"].Data(constants.Floor1, "Floor", constants.Floor1),
					selector["floor"].Data(constants.Floor2, "Floor", constants.Floor2),
					selector["floor"].Data(constants.Floor3, "Floor", constants.Floor3),
					selector["floor"].Data(constants.Floor4, "Floor", constants.Floor4),
					selector["floor"].Data(constants.Floor5, "Floor", constants.Floor5),
					selector["floor"].Data(constants.Floor6, "Floor", constants.Floor6),
					selector["floor"].Data(constants.Floor7, "Floor", constants.Floor7),
					selector["floor"].Data(constants.Floor8, "Floor", constants.Floor8),
					selector["floor"].Data(constants.Floor9, "Floor", constants.Floor9),
					selector["floor"].Data(constants.Floor10, "Floor", constants.Floor10),
					selector["floor"].Data(constants.FloorOver10, "Floor", constants.FloorOver10),
				},
			},
			storage: FilterValue{button: selector["menu"].Data(constants.StorageFilter, "Filters", "Storage"),
				value:     "",
				subButton: YNButtons,
			},
			elevator: FilterValue{button: selector["menu"].Data(constants.ElevatorFilter, "Filters", "Elevator"),
				value:     "",
				subButton: YNButtons,
			},
			adDate: FilterValue{button: selector["menu"].Data(constants.AdDateFilter, "Filters", "AdDate"),
				value: "",
				subButton: []telebot.Btn{
					selector["ad-date"].Data(constants.TimeToday, "Time", constants.TimeToday),
					selector["ad-date"].Data(constants.Time1DayAgo, "Time", constants.Time1DayAgo),
					selector["ad-date"].Data(constants.Time2DaysAgo, "Time", constants.Time2DaysAgo),
					selector["ad-date"].Data(constants.Time3DaysAgo, "Time", constants.Time3DaysAgo),
					selector["ad-date"].Data(constants.Time1WeekAgo, "Time", constants.Time1WeekAgo),
					selector["ad-date"].Data(constants.Time1MonthAgo, "Time", constants.Time1MonthAgo),
					selector["ad-date"].Data(constants.Time1YearAgo, "Time", constants.Time1YearAgo),
				},
			},
			location: FilterValue{button: selector["menu"].Data(constants.LocationFilter, "Filters", "Location"),
				value: "",
				subButton: []telebot.Btn{
					selector["location"].Data(constants.Tehran, "City", constants.Tehran),
					selector["location"].Data(constants.Esfahan, "City", constants.Esfahan),
					selector["location"].Data(constants.Mashhad, "City", constants.Mashhad),
					selector["location"].Data(constants.Shiraz, "City", constants.Shiraz),
				},
			},
		}
	)

	selector["menu"].Inline(
		selector["menu"].Row(filters.price.button),
		selector["menu"].Row(filters.area.button, filters.rooms.button, filters.propertyType.button),
		selector["menu"].Row(filters.buildingAge.button, filters.floor.button, filters.storage.button),
		selector["menu"].Row(filters.elevator.button, filters.adDate.button, filters.location.button),
		selector["menu"].Row(goBtn, removeBtn),
	)

	//Inline Styles
	selector["price"].Inline(selector["menu"].Split(1, filters.price.subButton)...)
	selector["area"].Inline(selector["menu"].Split(1, filters.area.subButton)...)
	selector["room"].Inline(selector["menu"].Split(3, filters.rooms.subButton)...)
	selector["property"].Inline(selector["menu"].Split(3, filters.propertyType.subButton)...)
	selector["age"].Inline(selector["menu"].Split(3, filters.buildingAge.subButton)...)
	selector["floor"].Inline(selector["menu"].Split(3, filters.floor.subButton)...)
	selector["yes-no"].Inline(selector["menu"].Split(3, YNButtons)...)
	selector["ad-date"].Inline(selector["menu"].Split(2, filters.adDate.subButton)...)
	selector["location"].Inline(selector["menu"].Split(4, filters.location.subButton)...)

	// Buttons Handlers
	b.Bot.Handle(&telebot.InlineButton{Unique: "Filters"}, func(ctx telebot.Context) error {
		filterSelected := ctx.Data()
		switch filterSelected {
		case "Price":
			b.Bot.Handle(&telebot.InlineButton{Unique: "PriceRange"}, func(c telebot.Context) error {
				filters.price.value = c.Data()
				return c.EditOrSend(filters.message(), selector["menu"])
			})
			return ctx.EditOrSend(filters.message(), selector["price"])
		case "Area":
			b.Bot.Handle(&telebot.InlineButton{Unique: "AreaRange"}, func(c telebot.Context) error {
				filters.area.value = c.Data()
				return c.EditOrSend(filters.message(), selector["menu"])
			})
			return ctx.EditOrSend(filters.message(), selector["area"])
		case "Rooms":
			b.Bot.Handle(&telebot.InlineButton{Unique: "NumberOfRooms"}, func(c telebot.Context) error {
				filters.rooms.value = strings.Split(c.Data(), "|")[0]
				return c.EditOrSend(filters.message(), selector["menu"])
			})
			return ctx.EditOrSend(filters.message(), selector["room"])
		case "PropertyType":
			b.Bot.Handle(&telebot.InlineButton{Unique: "Property"}, func(c telebot.Context) error {
				filters.propertyType.value = c.Data()
				return c.EditOrSend(filters.message(), selector["menu"])
			})
			return ctx.EditOrSend(filters.message(), selector["property"])
		case "BuildingAge":
			b.Bot.Handle(&telebot.InlineButton{Unique: "Age"}, func(c telebot.Context) error {
				filters.buildingAge.value = c.Data()
				return c.EditOrSend(filters.message(), selector["menu"])
			})
			return ctx.EditOrSend(filters.message(), selector["age"])
		case "Floor":
			b.Bot.Handle(&telebot.InlineButton{Unique: "Floor"}, func(c telebot.Context) error {
				filters.floor.value = c.Data()
				return c.EditOrSend(filters.message(), selector["menu"])
			})
			return ctx.EditOrSend(filters.message(), selector["floor"])
		case "Storage":
			b.Bot.Handle(&telebot.InlineButton{Unique: "YesNo"}, func(c telebot.Context) error {
				filters.storage.value = strings.Split(c.Data(), "|")[0]
				return c.EditOrSend(filters.message(), selector["menu"])
			})
			return ctx.EditOrSend(filters.message(), selector["yes-no"])
		case "Elevator":
			b.Bot.Handle(&telebot.InlineButton{Unique: "YesNo"}, func(c telebot.Context) error {
				filters.elevator.value = strings.Split(c.Data(), "|")[0]
				return c.EditOrSend(filters.message(), selector["menu"])
			})
			return ctx.EditOrSend(filters.message(), selector["yes-no"])
		case "AdDate":
			b.Bot.Handle(&telebot.InlineButton{Unique: "Time"}, func(c telebot.Context) error {
				filters.adDate.value = c.Data()
				return c.EditOrSend(filters.message(), selector["menu"])
			})
			return ctx.EditOrSend(filters.message(), selector["ad-date"])
		case "Location":
			b.Bot.Handle(&telebot.InlineButton{Unique: "City"}, func(c telebot.Context) error {
				filters.location.value = c.Data()
				return c.EditOrSend(filters.message(), selector["menu"])
			})
			return ctx.EditOrSend(filters.message(), selector["location"])
		case "Search":
			ctx.Send(constants.Loading)
			ads := filters.startSearch()
			for _, ad := range ads {
				ctx.Send(ad)
			}
			return ctx.Send(constants.SearchMsg)
		case "Remove":
			filters.removeAllValue()
			return ctx.EditOrSend(filters.message(), selector["menu"])
		default:
			filters.removeAllValue()
			return ctx.EditOrSend("/menu")
		}

	})

	return func(ctx telebot.Context) error {

		return ctx.EditOrSend(filters.message(), selector["menu"])
	}
}
