package bot

import (
	"fmt"
	"magical-crwler/constants"
	"magical-crwler/models"
	"time"

	"gopkg.in/telebot.v4"
)

type Filters struct {
	adType       FilterValue
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
	data      interface{}
	button    telebot.Btn
	subButton []telebot.Btn
}

func (f *Filters) startSearch() []models.Ad {

	return nil
}
func (f *Filters) removeAllValue() {
	f.adType.value = ""
	f.adType.data = nil
	f.price.value = ""
	f.price.data = nil
	f.area.value = ""
	f.area.data = nil
	f.rooms.value = ""
	f.rooms.data = nil
	f.propertyType.value = ""
	f.propertyType.data = nil
	f.buildingAge.value = ""
	f.buildingAge.data = nil
	f.floor.value = ""
	f.floor.data = nil
	f.storage.value = ""
	f.storage.data = nil
	f.elevator.value = ""
	f.elevator.data = nil
	f.adDate.value = ""
	f.adDate.data = nil
	f.location.value = ""
	f.location.data = nil
}

func (f *Filters) message() string {
	var msg string
	msg += fmt.Sprintf("🏘️%s\t:\n%s\n", f.adType.button.Text, f.adType.value)
	msg += fmt.Sprintf("💰%s\t:\n%s\n", f.price.button.Text, f.price.value)
	msg += fmt.Sprintf("📏%s\t:\n%s\n", f.area.button.Text, f.area.value)
	msg += fmt.Sprintf("🛏️%s\t:\n%s\n", f.rooms.button.Text, f.rooms.value)
	msg += fmt.Sprintf("🏘️%s\t:\n%s\n", f.propertyType.button.Text, f.propertyType.value)
	msg += fmt.Sprintf("🏚️%s\t:\n%s\n", f.buildingAge.button.Text, f.buildingAge.value)
	msg += fmt.Sprintf("🔃%s\t:\n%s\n", f.floor.button.Text, f.floor.value)
	msg += fmt.Sprintf("🚪%s\t:\n%s\n", f.storage.button.Text, f.storage.value)
	msg += fmt.Sprintf("🛗%s\t:\n%s\n", f.elevator.button.Text, f.elevator.value)
	msg += fmt.Sprintf("🗓️%s\t:\n%s\n", f.adDate.button.Text, f.adDate.value)
	msg += fmt.Sprintf("📍%s\t:\n%s\n", f.location.button.Text, f.location.value)

	return msg
}

func newReplyMarkup() *telebot.ReplyMarkup {
	return &telebot.ReplyMarkup{RemoveKeyboard: true, ResizeKeyboard: true}
}

func SearchHandlers(b *Bot) func(ctx telebot.Context) error {

	var (
		selector = map[string]*telebot.ReplyMarkup{
			"ad-type":  newReplyMarkup(),
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
			selector["yes-no"].Data(constants.Yes, "YesNo", "1"),
			selector["yes-no"].Data(constants.No, "YesNo", "0"),
			selector["yes-no"].Data(constants.Unknown, "YesNo", "-1"),
		}

		goBtn     = selector["menu"].Data(constants.GoButton, "Filters", "Search")
		removeBtn = selector["menu"].Data(constants.RemoveButton, "Filters", "Remove")

		filters = Filters{
			adType: FilterValue{button: selector["menu"].Data(constants.AdType, "Filters", "AdType"),
				value: "",
				subButton: []telebot.Btn{
					selector["ad-type"].Data(constants.ForBuy, "AdType", "at-buy"),
					selector["ad-type"].Data(constants.ForRent, "AdType", "at-rent"),
				},
			},
			price: FilterValue{button: selector["menu"].Data(constants.PriceFilter, "Filters", "Price"),
				value: "",
				subButton: []telebot.Btn{
					selector["price"].Data(constants.PriceUnder500M, "PriceRange", "PR-0"),
					selector["price"].Data(constants.Price500MTo700M, "PriceRange", "PR-500"),
					selector["price"].Data(constants.Price700MTo900M, "PriceRange", "PR-700"),
					selector["price"].Data(constants.Price900MTo1B, "PriceRange", "PR-900"),
					selector["price"].Data(constants.Price1BTo1_5B, "PriceRange", "PR-1000"),
					selector["price"].Data(constants.Price1_5BTo2B, "PriceRange", "PR-1500"),
					selector["price"].Data(constants.Price2BTo3B, "PriceRange", "PR-2000"),
					selector["price"].Data(constants.Price3BTo4B, "PriceRange", "PR-3000"),
					selector["price"].Data(constants.Price4BTo5B, "PriceRange", "PR-4000"),
					selector["price"].Data(constants.Price5BTo7B, "PriceRange", "PR-5000"),
					selector["price"].Data(constants.Price7BTo10B, "PriceRange", "PR-7000"),
					selector["price"].Data(constants.Price10BTo15B, "PriceRange", "PR-10000"),
					selector["price"].Data(constants.Price15BTo20B, "PriceRange", "PR-15000"),
					selector["price"].Data(constants.Price20BTo30B, "PriceRange", "PR-20000"),
					selector["price"].Data(constants.Price30BTo40B, "PriceRange", "PR-30000"),
					selector["price"].Data(constants.Price40BTo50B, "PriceRange", "PR-40000"),
					selector["price"].Data(constants.Price50BTo75B, "PriceRange", "PR-50000"),
					selector["price"].Data(constants.Price75BTo100B, "PriceRange", "PR-70000"),
					selector["price"].Data(constants.Price100BTo200B, "PriceRange", "PR-100000"),
					selector["price"].Data(constants.Price200BTo300B, "PriceRange", "PR-200000"),
					selector["price"].Data(constants.Price300BTo500B, "PriceRange", "PR-300000"),
					selector["price"].Data(constants.Price500BTo700B, "PriceRange", "PR-500000"),
					selector["price"].Data(constants.Price700BTo900B, "PriceRange", "PR-700000"),
					selector["price"].Data(constants.PriceOver900B, "PriceRange", "PR-900000+"),
				},
			},
			area: FilterValue{button: selector["menu"].Data(constants.AreaFilter, "Filters", "Area"),
				value: "",
				subButton: []telebot.Btn{
					selector["area"].Data(constants.AreaUnder50, "AreaRange", "AR-0"),
					selector["area"].Data(constants.Area50To75, "AreaRange", "AR-50"),
					selector["area"].Data(constants.Area75To100, "AreaRange", "AR-75"),
					selector["area"].Data(constants.Area100To150, "AreaRange", "AR-100"),

					selector["area"].Data(constants.Area150To200, "AreaRange", "AR-150"),
					selector["area"].Data(constants.Area200To250, "AreaRange", "AR-200"),
					selector["area"].Data(constants.Area250To300, "AreaRange", "AR-250"),
					selector["area"].Data(constants.Area300To400, "AreaRange", "AR-300"),
					selector["area"].Data(constants.Area400To500, "AreaRange", "AR-400"),
					selector["area"].Data(constants.Area500To750, "AreaRange", "AR-500"),
					selector["area"].Data(constants.Area750To1000, "AreaRange", "AR-750"),

					selector["area"].Data(constants.Area1000To1500, "AreaRange", "AR-1000"),
					selector["area"].Data(constants.Area1500To2000, "AreaRange", "AR-1500"),
					selector["area"].Data(constants.Area2000To3000, "AreaRange", "AR-2000"),
					selector["area"].Data(constants.Area3000To4000, "AreaRange", "AR-3000"),
					selector["area"].Data(constants.Area4000To5000, "AreaRange", "AR-4000"),
					selector["area"].Data(constants.Area5000To7500, "AreaRange", "AR-5000"),
					selector["area"].Data(constants.Area7500To10000, "AreaRange", "AR-7500"),
					selector["area"].Data(constants.AreaOver10000, "AreaRange", "AR-1000+"),
				},
			},
			rooms: FilterValue{button: selector["menu"].Data(constants.RoomsFilter, "Filters", "Rooms"),
				value: "",
				subButton: []telebot.Btn{
					selector["room"].Data(constants.Bedrooms0, "NumberOfRooms", "NR-0"),
					selector["room"].Data(constants.Bedrooms1, "NumberOfRooms", "NR-1"),
					selector["room"].Data(constants.Bedrooms2, "NumberOfRooms", "NR-2"),
					selector["room"].Data(constants.Bedrooms3, "NumberOfRooms", "NR-3"),
					selector["room"].Data(constants.Bedrooms4, "NumberOfRooms", "NR-4"),
					selector["room"].Data(constants.Bedrooms5, "NumberOfRooms", "NR-5"),
					selector["room"].Data(constants.Bedrooms6, "NumberOfRooms", "NR-6"),
					selector["room"].Data(constants.Bedrooms7, "NumberOfRooms", "NR-7"),
					selector["room"].Data(constants.Bedrooms8, "NumberOfRooms", "NR-8"),
					selector["room"].Data(constants.Bedrooms9, "NumberOfRooms", "NR-9"),
					selector["room"].Data(constants.Bedrooms10, "NumberOfRooms", "NR-10"),
					selector["room"].Data(constants.BedroomsOver10, "NumberOfRooms", "NR-10+"),
				},
			},
			propertyType: FilterValue{button: selector["menu"].Data(constants.PropertyTypeFilter, "Filters", "PropertyType"),
				value: "",
				subButton: []telebot.Btn{
					selector["property"].Data(constants.PropertyApartment, "Property", "PA-Apartment"),
					selector["property"].Data(constants.PropertyVilla, "Property", "PA-Villa"),
					selector["property"].Data(constants.PropertyCommercial, "Property", "PA-Commercial"),
					selector["property"].Data(constants.PropertyOffice, "Property", "PA-Office"),
					selector["property"].Data(constants.PropertyLand, "Property", "PA-Land"),
				},
			},
			buildingAge: FilterValue{button: selector["menu"].Data(constants.BuildingAgeFilter, "Filters", "BuildingAge"),
				value: "",
				subButton: []telebot.Btn{
					selector["age"].Data(constants.BuildingAgeNew, "Age", "BA-New"),
					selector["age"].Data(constants.BuildingAge1Year, "Age", "BA-1Year"),
					selector["age"].Data(constants.BuildingAge2Years, "Age", "BA-2Years"),
					selector["age"].Data(constants.BuildingAge3Years, "Age", "BA-3Years"),
					selector["age"].Data(constants.BuildingAge4Years, "Age", "BA-4Years"),
					selector["age"].Data(constants.BuildingAge5Years, "Age", "BA-5Years"),
					selector["age"].Data(constants.BuildingAge6Years, "Age", "BA-6Years"),
					selector["age"].Data(constants.BuildingAge7Years, "Age", "BA-7Years"),
					selector["age"].Data(constants.BuildingAge8Years, "Age", "BA-8Years"),
					selector["age"].Data(constants.BuildingAge9Years, "Age", "BA-9Years"),
					selector["age"].Data(constants.BuildingAge10Years, "Age", "BA-10Years"),
					selector["age"].Data(constants.BuildingAgeOver10, "Age", "BA-Over10"),
				},
			},
			floor: FilterValue{button: selector["menu"].Data(constants.FloorFilter, "Filters", "Floor"),
				value: "",
				subButton: []telebot.Btn{
					selector["floor"].Data(constants.Floor0, "Floor", "FLR-0"),
					selector["floor"].Data(constants.Floor1, "Floor", "FLR-1"),
					selector["floor"].Data(constants.Floor2, "Floor", "FLR-2"),
					selector["floor"].Data(constants.Floor3, "Floor", "FLR-3"),
					selector["floor"].Data(constants.Floor4, "Floor", "FLR-4"),
					selector["floor"].Data(constants.Floor5, "Floor", "FLR-5"),
					selector["floor"].Data(constants.Floor6, "Floor", "FLR-6"),
					selector["floor"].Data(constants.Floor7, "Floor", "FLR-7"),
					selector["floor"].Data(constants.Floor8, "Floor", "FLR-8"),
					selector["floor"].Data(constants.Floor9, "Floor", "FLR-9"),
					selector["floor"].Data(constants.Floor10, "Floor", "FLR-10"),
					selector["floor"].Data(constants.FloorOver10, "Floor", "FLR-10+"),
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
					selector["ad-date"].Data(constants.TimeToday, "Time", "TM-Today"),
					selector["ad-date"].Data(constants.Time1DayAgo, "Time", "TM-1DayAgo"),
					selector["ad-date"].Data(constants.Time2DaysAgo, "Time", "TM-2DaysAgo"),
					selector["ad-date"].Data(constants.Time3DaysAgo, "Time", "TM-3DaysAgo"),
					selector["ad-date"].Data(constants.Time1WeekAgo, "Time", "TM-1WeekAgo"),
					selector["ad-date"].Data(constants.Time1MonthAgo, "Time", "TM-1MonthAgo"),
					selector["ad-date"].Data(constants.Time1YearAgo, "Time", "TM-1YearAgo"),
				},
			},
			location: FilterValue{button: selector["menu"].Data(constants.LocationFilter, "Filters", "Location"),
				value: "",
				subButton: []telebot.Btn{
					selector["location"].Data(constants.Tehran, "City", "Tehran"),
					selector["location"].Data(constants.Zanjan, "City", "Zanjan"),
					selector["location"].Data(constants.Khoram, "City", "Khoram"),
					selector["location"].Data(constants.Mazandaran, "City", "Mazandaran"),
				},
			},
		}
	)

	selector["menu"].Inline(
		selector["menu"].Row(filters.adType.button, filters.price.button),
		selector["menu"].Row(filters.area.button, filters.rooms.button, filters.propertyType.button),
		selector["menu"].Row(filters.buildingAge.button, filters.floor.button, filters.storage.button),
		selector["menu"].Row(filters.elevator.button, filters.adDate.button, filters.location.button),
		selector["menu"].Row(goBtn, removeBtn),
	)

	//Inline Styles
	selector["ad-type"].Inline(selector["menu"].Split(2, filters.adType.subButton)...)
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
		case "AdType":
			b.Bot.Handle(&telebot.InlineButton{Unique: "AdType"}, func(c telebot.Context) error {
				data := getValue(c.Data())
				filters.adType.value = fmt.Sprintf("%v", data[0])
				filters.location.data = data[1]
				return c.EditOrSend(filters.message(), selector["menu"])
			})
			return ctx.EditOrSend(filters.message(), selector["ad-type"])
		case "Price":
			b.Bot.Handle(&telebot.InlineButton{Unique: "PriceRange"}, func(c telebot.Context) error {
				data := getValue(c.Data())
				filters.price.value = fmt.Sprintf("%v", data[0])
				filters.location.data = data[1]
				return c.EditOrSend(filters.message(), selector["menu"])
			})
			return ctx.EditOrSend(filters.message(), selector["price"])
		case "Area":
			b.Bot.Handle(&telebot.InlineButton{Unique: "AreaRange"}, func(c telebot.Context) error {
				data := getValue(c.Data())
				filters.area.value = fmt.Sprintf("%v", data[0])
				filters.location.data = data[1]
				return c.EditOrSend(filters.message(), selector["menu"])
			})
			return ctx.EditOrSend(filters.message(), selector["area"])
		case "Rooms":
			b.Bot.Handle(&telebot.InlineButton{Unique: "NumberOfRooms"}, func(c telebot.Context) error {
				data := getValue(c.Data())
				filters.rooms.value = fmt.Sprintf("%v", data[0])
				filters.location.data = data[1]
				return c.EditOrSend(filters.message(), selector["menu"])
			})
			return ctx.EditOrSend(filters.message(), selector["room"])
		case "PropertyType":
			b.Bot.Handle(&telebot.InlineButton{Unique: "Property"}, func(c telebot.Context) error {
				data := getValue(c.Data())
				filters.propertyType.value = fmt.Sprintf("%v", data[0])
				filters.location.data = data[1]
				return c.EditOrSend(filters.message(), selector["menu"])
			})
			return ctx.EditOrSend(filters.message(), selector["property"])
		case "BuildingAge":
			b.Bot.Handle(&telebot.InlineButton{Unique: "Age"}, func(c telebot.Context) error {
				data := getValue(c.Data())
				filters.buildingAge.value = fmt.Sprintf("%v", data[0])
				filters.location.data = data[1]
				return c.EditOrSend(filters.message(), selector["menu"])
			})
			return ctx.EditOrSend(filters.message(), selector["age"])
		case "Floor":
			b.Bot.Handle(&telebot.InlineButton{Unique: "Floor"}, func(c telebot.Context) error {
				data := getValue(c.Data())
				filters.floor.value = fmt.Sprintf("%v", data[0])
				filters.location.data = data[1]
				return c.EditOrSend(filters.message(), selector["menu"])
			})
			return ctx.EditOrSend(filters.message(), selector["floor"])
		case "Storage":
			b.Bot.Handle(&telebot.InlineButton{Unique: "YesNo"}, func(c telebot.Context) error {
				data := getValue(c.Data())
				filters.storage.value = fmt.Sprintf("%v", data[0])
				filters.location.data = data[1]
				return c.EditOrSend(filters.message(), selector["menu"])
			})
			return ctx.EditOrSend(filters.message(), selector["yes-no"])
		case "Elevator":
			b.Bot.Handle(&telebot.InlineButton{Unique: "YesNo"}, func(c telebot.Context) error {
				data := getValue(c.Data())
				filters.elevator.value = fmt.Sprintf("%v", data[0])
				filters.location.data = data[1]
				return c.EditOrSend(filters.message(), selector["menu"])
			})
			return ctx.EditOrSend(filters.message(), selector["yes-no"])
		case "AdDate":
			b.Bot.Handle(&telebot.InlineButton{Unique: "Time"}, func(c telebot.Context) error {
				data := getValue(c.Data())
				filters.adDate.value = fmt.Sprintf("%v", data[0])
				filters.location.data = data[1]
				return c.EditOrSend(filters.message(), selector["menu"])
			})
			return ctx.EditOrSend(filters.message(), selector["ad-date"])
		case "Location":
			b.Bot.Handle(&telebot.InlineButton{Unique: "City"}, func(c telebot.Context) error {
				data := getValue(c.Data())
				filters.location.value = fmt.Sprintf("%v", data[0])
				filters.location.data = data[1]
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

func getValue(value string) []interface{} {
	btnMap := map[string][]interface{}{

		"PR-0":          {constants.PriceUnder500M, models.Range{Min: 0, Max: 500}},
		"PR-500":        {constants.Price500MTo700M, models.Range{Min: 500, Max: 700}},
		"PR-700":        {constants.Price700MTo900M, models.Range{Min: 700, Max: 900}},
		"PR-900":        {constants.Price900MTo1B, models.Range{Min: 900, Max: 1_000}},
		"PR-1000":       {constants.Price1BTo1_5B, models.Range{Min: 1_000, Max: 1_500}},
		"PR-1500":       {constants.Price1_5BTo2B, models.Range{Min: 1_500, Max: 2_000}},
		"PR-2000":       {constants.Price2BTo3B, models.Range{Min: 2_000, Max: 3_000}},
		"PR-3000":       {constants.Price3BTo4B, models.Range{Min: 3_000, Max: 4_000}},
		"PR-4000":       {constants.Price4BTo5B, models.Range{Min: 4_000, Max: 5_000}},
		"PR-5000":       {constants.Price5BTo7B, models.Range{Min: 5_000, Max: 7_000}},
		"PR-7000":       {constants.Price7BTo10B, models.Range{Min: 7_000, Max: 10_000}},
		"PR-10000":      {constants.Price10BTo15B, models.Range{Min: 10_000, Max: 15_000}},
		"PR-15000":      {constants.Price15BTo20B, models.Range{Min: 15_000, Max: 20_000}},
		"PR-20000":      {constants.Price20BTo30B, models.Range{Min: 20_000, Max: 30_000}},
		"PR-30000":      {constants.Price30BTo40B, models.Range{Min: 30_000, Max: 40_000}},
		"PR-40000":      {constants.Price40BTo50B, models.Range{Min: 40_000, Max: 50_000}},
		"PR-50000":      {constants.Price50BTo75B, models.Range{Min: 50_000, Max: 70_000}},
		"PR-70000":      {constants.Price75BTo100B, models.Range{Min: 70_000, Max: 100_000}},
		"PR-100000":     {constants.Price100BTo200B, models.Range{Min: 100_000, Max: 200_000}},
		"PR-200000":     {constants.Price200BTo300B, models.Range{Min: 200_000, Max: 300_000}},
		"PR-300000":     {constants.Price300BTo500B, models.Range{Min: 300_000, Max: 500_000}},
		"PR-500000":     {constants.Price500BTo700B, models.Range{Min: 500_000, Max: 700_000}},
		"PR-700000":     {constants.Price700BTo900B, models.Range{Min: 700_000, Max: 900_000}},
		"PR-900000+":    {constants.PriceOver900B, models.Range{Min: 900_000, Max: 0}},
		"at-buy":        {constants.ForBuy, false},
		"at-rent":       {constants.ForRent, true},
		"-1":            {constants.Unknown, nil},
		"0":             {constants.No, false},
		"1":             {constants.Yes, true},
		"AR-0":          {constants.AreaUnder50, models.Range{Min: 0, Max: 50}},
		"AR-50":         {constants.Area50To75, models.Range{Min: 50, Max: 75}},
		"AR-75":         {constants.Area75To100, models.Range{Min: 75, Max: 100}},
		"AR-100":        {constants.Area100To150, models.Range{Min: 100, Max: 150}},
		"AR-150":        {constants.Area150To200, models.Range{Min: 150, Max: 200}},
		"AR-200":        {constants.Area200To250, models.Range{Min: 200, Max: 250}},
		"AR-250":        {constants.Area250To300, models.Range{Min: 250, Max: 300}},
		"AR-300":        {constants.Area300To400, models.Range{Min: 300, Max: 400}},
		"AR-400":        {constants.Area400To500, models.Range{Min: 400, Max: 500}},
		"AR-500":        {constants.Area500To750, models.Range{Min: 500, Max: 700}},
		"AR-750":        {constants.Area750To1000, models.Range{Min: 700, Max: 1_000}},
		"AR-1000":       {constants.Area1000To1500, models.Range{Min: 1_000, Max: 1_500}},
		"AR-1500":       {constants.Area1500To2000, models.Range{Min: 1_500, Max: 2_000}},
		"AR-2000":       {constants.Area2000To3000, models.Range{Min: 2_000, Max: 3_000}},
		"AR-3000":       {constants.Area3000To4000, models.Range{Min: 3_000, Max: 4_000}},
		"AR-4000":       {constants.Area4000To5000, models.Range{Min: 4_000, Max: 5_000}},
		"AR-5000":       {constants.Area5000To7500, models.Range{Min: 5_000, Max: 7_500}},
		"AR-7500":       {constants.Area7500To10000, models.Range{Min: 7_5000, Max: 9_000}},
		"AR-1000+":      {constants.AreaOver10000, models.Range{Min: 9_000, Max: 0}},
		"NR-0":          {constants.Bedrooms0, models.Range{Min: 0, Max: 0}},
		"NR-1":          {constants.Bedrooms1, models.Range{Min: 0, Max: 1}},
		"NR-2":          {constants.Bedrooms2, models.Range{Min: 0, Max: 2}},
		"NR-3":          {constants.Bedrooms3, models.Range{Min: 0, Max: 3}},
		"NR-4":          {constants.Bedrooms4, models.Range{Min: 0, Max: 4}},
		"NR-5":          {constants.Bedrooms5, models.Range{Min: 0, Max: 5}},
		"NR-6":          {constants.Bedrooms6, models.Range{Min: 0, Max: 6}},
		"NR-7":          {constants.Bedrooms7, models.Range{Min: 0, Max: 7}},
		"NR-8":          {constants.Bedrooms8, models.Range{Min: 0, Max: 8}},
		"NR-9":          {constants.Bedrooms9, models.Range{Min: 0, Max: 9}},
		"NR-10":         {constants.Bedrooms10, models.Range{Min: 0, Max: 10}},
		"NR-10+":        {constants.BedroomsOver10, models.Range{Min: 0, Max: 11}},
		"PA-Apartment":  {constants.PropertyApartment, "Apartment"},
		"PA-Villa":      {constants.PropertyVilla, "Villa"},
		"PA-Commercial": {constants.PropertyCommercial, "Commercial"},
		"PA-Office":     {constants.PropertyOffice, "Office"},
		"PA-Land":       {constants.PropertyLand, "Land"},
		"BA-New":        {constants.BuildingAgeNew, 0},
		"BA-1Year":      {constants.BuildingAge1Year, 1},
		"BA-2Years":     {constants.BuildingAge2Years, 2},
		"BA-3Years":     {constants.BuildingAge3Years, 3},
		"BA-4Years":     {constants.BuildingAge4Years, 4},
		"BA-5Years":     {constants.BuildingAge5Years, 5},
		"BA-6Years":     {constants.BuildingAge6Years, 6},
		"BA-7Years":     {constants.BuildingAge7Years, 7},
		"BA-8Years":     {constants.BuildingAge8Years, 8},
		"BA-9Years":     {constants.BuildingAge9Years, 9},
		"BA-10Years":    {constants.BuildingAge10Years, 10},
		"BA-Over10":     {constants.BuildingAgeOver10, 11},
		"FLR-0":         {constants.Floor0, models.Range{Min: 0, Max: 0}},
		"FLR-1":         {constants.Floor1, models.Range{Min: 0, Max: 1}},
		"FLR-2":         {constants.Floor2, models.Range{Min: 0, Max: 2}},
		"FLR-3":         {constants.Floor3, models.Range{Min: 0, Max: 3}},
		"FLR-4":         {constants.Floor4, models.Range{Min: 0, Max: 4}},
		"FLR-5":         {constants.Floor5, models.Range{Min: 0, Max: 5}},
		"FLR-6":         {constants.Floor6, models.Range{Min: 0, Max: 6}},
		"FLR-7":         {constants.Floor7, models.Range{Min: 0, Max: 7}},
		"FLR-8":         {constants.Floor8, models.Range{Min: 0, Max: 8}},
		"FLR-9":         {constants.Floor9, models.Range{Min: 0, Max: 9}},
		"FLR-10":        {constants.Floor10, models.Range{Min: 0, Max: 10}},
		"FLR-10+":       {constants.FloorOver10, models.Range{Min: 0, Max: 11}},
		"TM-Today":      {constants.TimeToday, time.Now()},
		"TM-1DayAgo":    {constants.Time1DayAgo, time.Now()},
		"TM-2DaysAgo":   {constants.Time2DaysAgo, time.Now()},
		"TM-3DaysAgo":   {constants.Time3DaysAgo, time.Now()},
		"TM-1WeekAgo":   {constants.Time1WeekAgo, time.Now()},
		"TM-1MonthAgo":  {constants.Time1MonthAgo, time.Now()},
		"TM-1YearAgo":   {constants.Time1YearAgo, time.Now()},
		"Tehran":        {constants.Tehran, "Tehran"},
		"Zanjan":        {constants.Zanjan, "Zanjan"},
		"Khoram":        {constants.Khoram, "Khoram"},
		"Mazandaran":    {constants.Mazandaran, "Mazandaran"},
	}
	return btnMap[value]
}
