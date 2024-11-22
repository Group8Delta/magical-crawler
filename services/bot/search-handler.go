package bot

import (
	"fmt"
	"log"
	"magical-crwler/constants"
	"magical-crwler/database"
	"magical-crwler/models"
	"magical-crwler/models/Dtos"
	"magical-crwler/utils"
	"strconv"
	"strings"
	"time"

	"gopkg.in/telebot.v4"
)

type Filters struct {
	adType       *FilterValue
	price        *FilterValue
	area         *FilterValue
	rooms        *FilterValue
	propertyType *FilterValue
	buildingAge  *FilterValue
	floor        *FilterValue
	storage      *FilterValue
	elevator     *FilterValue
	adDate       *FilterValue
	location     *FilterValue
}
type FilterValue struct {
	value     string
	data      Dtos.FilterDto
	button    telebot.Btn
	subButton []telebot.Btn
}

func (f *Filters) startSearch(repo database.IRepository) ([]models.Ad, models.Filter, error) {
	filters := Dtos.FilterDto{
		PriceRange:            f.price.data.PriceRange,
		RentPriceRange:        f.price.data.RentPriceRange,
		ForRent:               f.adType.data.ForRent,
		City:                  f.location.data.City,
		SizeRange:             f.area.data.SizeRange,
		BedroomRange:          f.rooms.data.BedroomRange,
		FloorRange:            f.floor.data.FloorRange,
		HasStorage:            f.storage.data.HasStorage,
		HasElevator:           f.elevator.data.HasElevator,
		AgeRange:              f.buildingAge.data.AgeRange,
		IsApartment:           f.propertyType.data.IsApartment,
		CreationTimeRangeFrom: f.adDate.data.CreationTimeRangeFrom,
		CreationTimeRangeTo:   time.Now(),
	}
	filter := repo.CreateFilter(filters)
	ads, err := repo.GetAdsByFilterId(int(filter.ID))
	if err != nil {
		return nil, filter, err
	}

	return ads, filter, nil
}
func (f *Filters) removeAllValue() {
	f.adType.value = ""
	f.adType.data.ForRent = false
	f.price.value = ""
	f.price.data.PriceRange = nil
	f.area.value = ""
	f.area.data.SizeRange = nil
	f.rooms.value = ""
	f.rooms.data.BedroomRange = nil
	f.propertyType.value = ""
	f.propertyType.data.IsApartment = nil
	f.buildingAge.value = ""
	f.buildingAge.data.AgeRange = nil
	f.floor.value = ""
	f.floor.data.FloorRange = nil
	f.storage.value = ""
	f.storage.data.HasStorage = nil
	f.elevator.value = ""
	f.elevator.data.HasElevator = nil
	f.adDate.value = ""
	f.adDate.data.CreationTimeRangeTo = time.Time{}
	f.adDate.data.CreationTimeRangeFrom = time.Time{}
	f.location.value = ""
	f.location.data.City = nil
	f.location.data.Neighborhood = nil
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

func SearchHandlers(b *Bot, user *models.User, db database.DbService) func(ctx telebot.Context) error {

	var (
		selector = map[string]*telebot.ReplyMarkup{
			"ad-type":       newReplyMarkup(),
			"menu":          newReplyMarkup(),
			"price":         newReplyMarkup(),
			"area":          newReplyMarkup(),
			"room":          newReplyMarkup(),
			"property":      newReplyMarkup(),
			"age":           newReplyMarkup(),
			"floor":         newReplyMarkup(),
			"yes-no":        newReplyMarkup(),
			"ad-date":       newReplyMarkup(),
			"location":      newReplyMarkup(),
			"price-history": newReplyMarkup(),
			"export":        newReplyMarkup(),
		}

		YNButtons = []telebot.Btn{
			selector["yes-no"].Data(constants.Yes, "YesNo", "Yes"),
			selector["yes-no"].Data(constants.No, "YesNo", "No"),
		}

		goBtn     = selector["menu"].Data(constants.GoButton, "Filters", "Search")
		removeBtn = selector["menu"].Data(constants.RemoveButton, "Filters", "Remove")
		wlBtn     = selector["menu"].Data(constants.WatchListButton, "Filters", "WatchListBtn")

		filters = Filters{
			adType: &FilterValue{button: selector["menu"].Data(constants.AdType, "Filters", "AdType"),
				value: "",
				data:  Dtos.FilterDto{},
				subButton: []telebot.Btn{
					selector["ad-type"].Data(constants.ForBuy, "AdType", "at-buy"),
					selector["ad-type"].Data(constants.ForRent, "AdType", "at-rent"),
				},
			},
			price: &FilterValue{button: selector["menu"].Data(constants.PriceFilter, "Filters", "Price"),
				value: "",
				data:  Dtos.FilterDto{},
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
			area: &FilterValue{button: selector["menu"].Data(constants.AreaFilter, "Filters", "Area"),
				value: "",
				data:  Dtos.FilterDto{},
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
			rooms: &FilterValue{button: selector["menu"].Data(constants.RoomsFilter, "Filters", "Rooms"),
				value: "",
				data:  Dtos.FilterDto{},
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
			propertyType: &FilterValue{button: selector["menu"].Data(constants.PropertyTypeFilter, "Filters", "PropertyType"),
				value: "",
				data:  Dtos.FilterDto{},
				subButton: []telebot.Btn{
					selector["property"].Data(constants.PropertyApartment, "Property", "PA-Apartment"),
					selector["property"].Data(constants.PropertyVilla, "Property", "PA-Villa"),
				},
			},
			buildingAge: &FilterValue{button: selector["menu"].Data(constants.BuildingAgeFilter, "Filters", "BuildingAge"),
				value: "",
				data:  Dtos.FilterDto{},
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
			floor: &FilterValue{button: selector["menu"].Data(constants.FloorFilter, "Filters", "Floor"),
				value: "",
				data:  Dtos.FilterDto{},
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
			storage: &FilterValue{button: selector["menu"].Data(constants.StorageFilter, "Filters", "Storage"),
				value:     "",
				subButton: YNButtons,
			},
			elevator: &FilterValue{button: selector["menu"].Data(constants.ElevatorFilter, "Filters", "Elevator"),
				value:     "",
				subButton: YNButtons,
			},
			adDate: &FilterValue{button: selector["menu"].Data(constants.AdDateFilter, "Filters", "AdDate"),
				value: "",
				data:  Dtos.FilterDto{},
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
			location: &FilterValue{button: selector["menu"].Data(constants.LocationFilter, "Filters", "Location"),
				value: "",
				data:  Dtos.FilterDto{},
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
		selector["menu"].Row(wlBtn),
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
	selector["yes-no"].Inline(selector["menu"].Split(2, YNButtons)...)
	selector["ad-date"].Inline(selector["menu"].Split(2, filters.adDate.subButton)...)
	selector["location"].Inline(selector["menu"].Split(4, filters.location.subButton)...)
	selector["export"].Inline(selector["export"].Split(2, YNButtons)...)

	// Buttons Handlers
	b.Bot.Handle(&telebot.InlineButton{Unique: "Filters"}, func(ctx telebot.Context) error {
		filterSelected := ctx.Data()
		switch filterSelected {
		case "AdType":
			b.Bot.Handle(&telebot.InlineButton{Unique: "AdType"}, func(c telebot.Context) error {
				data := getValue(c.Data())
				filters.adType.value = data[0]
				filters.adType.data.ForRent, _ = strconv.ParseBool(data[1])
				return c.EditOrSend(filters.message(), selector["menu"])
			})
			return ctx.EditOrSend(filters.message(), selector["ad-type"])
		case "Price":
			b.Bot.Handle(&telebot.InlineButton{Unique: "PriceRange"}, func(c telebot.Context) error {
				data := getValue(c.Data())
				filters.price.value = data[0]
				priceRange := strings.Split(data[1], ",")
				min, err := strconv.Atoi(priceRange[0])
				if err != nil {
					return err
				}
				max, err := strconv.Atoi(priceRange[1])
				if err != nil {
					return err
				}
				filters.price.data.PriceRange = &models.Range{Min: min, Max: max}
				return c.EditOrSend(filters.message(), selector["menu"])
			})
			return ctx.EditOrSend(filters.message(), selector["price"])
		case "Area":
			b.Bot.Handle(&telebot.InlineButton{Unique: "AreaRange"}, func(c telebot.Context) error {
				data := getValue(c.Data())
				filters.area.value = data[0]
				areaRange := strings.Split(data[1], ",")
				min, err := strconv.Atoi(areaRange[0])
				if err != nil {
					return err
				}
				max, err := strconv.Atoi(areaRange[1])
				if err != nil {
					return err
				}
				filters.area.data.SizeRange = &models.Range{Min: min, Max: max}
				return c.EditOrSend(filters.message(), selector["menu"])
			})
			return ctx.EditOrSend(filters.message(), selector["area"])
		case "Rooms":
			b.Bot.Handle(&telebot.InlineButton{Unique: "NumberOfRooms"}, func(c telebot.Context) error {
				data := getValue(c.Data())
				filters.rooms.value = data[0]
				rooms := strings.Split((data[1]), ",")
				min, err := strconv.Atoi(rooms[0])
				if err != nil {
					return err
				}
				max, err := strconv.Atoi(rooms[1])
				if err != nil {
					return err
				}
				filters.rooms.data.BedroomRange = &models.Range{Min: min, Max: max}
				return c.EditOrSend(filters.message(), selector["menu"])
			})
			return ctx.EditOrSend(filters.message(), selector["room"])
		case "PropertyType":
			b.Bot.Handle(&telebot.InlineButton{Unique: "Property"}, func(c telebot.Context) error {
				data := getValue(c.Data())
				filters.propertyType.value = data[0]
				b, err := strconv.ParseBool(data[1])
				if err != nil {
					return err
				}
				filters.location.data.IsApartment = &b
				return c.EditOrSend(filters.message(), selector["menu"])
			})
			return ctx.EditOrSend(filters.message(), selector["property"])
		case "BuildingAge":
			b.Bot.Handle(&telebot.InlineButton{Unique: "Age"}, func(c telebot.Context) error {
				data := getValue(c.Data())
				filters.buildingAge.value = data[0]
				age := strings.Split((data[1]), ",")
				min, err := strconv.Atoi(age[0])
				if err != nil {
					return err
				}
				max, err := strconv.Atoi(age[1])
				if err != nil {
					return err
				}
				filters.buildingAge.data.AgeRange = &models.Range{Min: min, Max: max}
				return c.EditOrSend(filters.message(), selector["menu"])
			})
			return ctx.EditOrSend(filters.message(), selector["age"])
		case "Floor":
			b.Bot.Handle(&telebot.InlineButton{Unique: "Floor"}, func(c telebot.Context) error {
				data := getValue(c.Data())
				filters.floor.value = data[0]
				floor := strings.Split((data[1]), ",")
				min, err := strconv.Atoi(floor[0])
				if err != nil {
					return err
				}
				max, err := strconv.Atoi(floor[1])
				if err != nil {
					return err
				}
				filters.floor.data.FloorRange = &models.Range{Min: min, Max: max}
				return c.EditOrSend(filters.message(), selector["menu"])
			})
			return ctx.EditOrSend(filters.message(), selector["floor"])
		case "Storage":
			b.Bot.Handle(&telebot.InlineButton{Unique: "YesNo"}, func(c telebot.Context) error {
				data := getValue(c.Data())
				filters.storage.value = data[0]
				b, err := strconv.ParseBool(data[1])
				if err != nil {
					return err
				}
				filters.storage.data.HasStorage = &b
				return c.EditOrSend(filters.message(), selector["menu"])
			})
			return ctx.EditOrSend(filters.message(), selector["yes-no"])
		case "Elevator":
			b.Bot.Handle(&telebot.InlineButton{Unique: "YesNo"}, func(c telebot.Context) error {
				data := getValue(c.Data())
				filters.elevator.value = data[0]
				b, err := strconv.ParseBool(data[1])
				if err != nil {
					return err
				}
				filters.elevator.data.HasElevator = &b
				return c.EditOrSend(filters.message(), selector["menu"])
			})
			return ctx.EditOrSend(filters.message(), selector["yes-no"])
		case "AdDate":
			b.Bot.Handle(&telebot.InlineButton{Unique: "Time"}, func(c telebot.Context) error {
				data := c.Data()
				now := time.Now()
				switch data {
				case "TM-Today":
					filters.adDate.data.CreationTimeRangeTo = now
					filters.adDate.value = constants.TimeToday
				case "TM-1DayAgo":
					filters.adDate.data.CreationTimeRangeTo = now.AddDate(0, 0, -1)
					filters.adDate.value = constants.Time1DayAgo
				case "TM-2DaysAgo":
					filters.adDate.data.CreationTimeRangeTo = now.AddDate(0, 0, -2)
					filters.adDate.value = constants.Time2DaysAgo
				case "TM-3DaysAgo":
					filters.adDate.data.CreationTimeRangeTo = now.AddDate(0, 0, -3)
					filters.adDate.value = constants.Time3DaysAgo
				case "TM-1WeekAgo":
					filters.adDate.data.CreationTimeRangeTo = now.AddDate(0, 0, -7)
					filters.adDate.value = constants.Time1WeekAgo
				case "TM-1MonthAgo":
					filters.adDate.data.CreationTimeRangeTo = now.AddDate(0, -1, 0)
					filters.adDate.value = constants.Time1MonthAgo
				case "TM-1YearAgo":
					filters.adDate.data.CreationTimeRangeTo = now.AddDate(-1, 0, 0)
					filters.adDate.value = constants.Time1YearAgo
				default:
					filters.adDate.data.CreationTimeRangeTo = now
				}
				filters.adDate.data.CreationTimeRangeFrom = now
				return c.EditOrSend(filters.message(), selector["menu"])
			})
			return ctx.EditOrSend(filters.message(), selector["ad-date"])
		case "Location":
			b.Bot.Handle(&telebot.InlineButton{Unique: "City"}, func(c telebot.Context) error {
				data := getValue(c.Data())
				filters.location.value = data[0]
				filters.location.data.City = &data[1]
				return c.EditOrSend(filters.message(), selector["menu"])
			})
			return ctx.EditOrSend(filters.message(), selector["location"])
		case "Search":
			ctx.Send(constants.Loading)
			ads, filter, err := filters.startSearch(b.repo)
			if err != nil {
				log.Println(err)
				return ctx.Send(err)
			}
			if len(ads) < 1 {
				return ctx.Send(constants.NoAdMsg)
			}
			for _, ad := range ads {
				photoURL := ad.PhotoUrl
				priceHistoryBtn := selector["price-history"].Data(constants.PriceHistory, "PriceHistory", fmt.Sprintf("%d", ad.ID))
				selector["price-history"].Inline(selector["prce-history"].Row(priceHistoryBtn))

				if photoURL != nil {
					pic := &telebot.Photo{File: telebot.FromURL(*photoURL)}
					ctx.SendAlbum(telebot.Album{pic})
				}
				ctx.Send(utils.GenerateFilterMessage(ad), telebot.ModeHTML, selector["price-history"])
				ctx.Set(fmt.Sprintf("%d", ad.ID), ad)
				b.Bot.Handle(&telebot.InlineButton{Unique: "PriceHistory"}, func(c telebot.Context) error {
					var data models.Ad
					if ctx.Get(c.Data()) != nil {
						data = ctx.Get(c.Data()).(models.Ad)
					}

					list, err := b.repo.GetPriceHistory(data.ID)
					if err != nil {
						return c.Send("Error in fetch data")
					}
					log.Println(list)
					return c.Send(utils.GeneratePriceHistory(list))
				})
			}

			adIDs := make([]uint, len(ads))
			for i, ad := range ads {
				adIDs[i] = ad.ID
			}
			repo := database.NewRepository(db)
			if err := repo.IncrementVisitCounts(adIDs); err != nil {
				return err
			}

			filteredAds := make([]models.FilteredAd, len(ads))
			timestamp := time.Now()
			for i, ad := range ads {
				filteredAds[i] = models.FilteredAd{
					UserID:    user.ID,
					FilterID:  filter.ID,
					AdID:      ad.ID,
					TimeStamp: timestamp,
				}
			}

			if err := repo.InsertFilteredAds(filteredAds); err != nil {
				return err
			}
			b.Bot.Handle(&telebot.InlineButton{Unique: "YesNo"}, exportFileHandler(ads, b))
			filters.removeAllValue()
			return ctx.Send(constants.ExportMsg, selector["export"])
		case "Remove":
			filters.removeAllValue()
			return ctx.EditOrSend(filters.message(), selector["menu"])
		case "WatchListBtn":
			ctx.EditOrSend(constants.TimeRangeMsg)
			b.Bot.Handle(telebot.OnText, func(c telebot.Context) error {
				// Retrieve the user message
				userMessage := c.Message().Text
				m, err := strconv.Atoi(userMessage)
				if err != nil {
					return ctx.Send(constants.TimeRangeError)

				}
				f := Dtos.FilterDto{
					PriceRange:            filters.price.data.PriceRange,
					RentPriceRange:        filters.price.data.RentPriceRange,
					ForRent:               filters.adType.data.ForRent,
					City:                  filters.location.data.City,
					SizeRange:             filters.area.data.SizeRange,
					BedroomRange:          filters.rooms.data.BedroomRange,
					FloorRange:            filters.floor.data.FloorRange,
					HasStorage:            filters.storage.data.HasStorage,
					HasElevator:           filters.elevator.data.HasElevator,
					AgeRange:              filters.buildingAge.data.AgeRange,
					IsApartment:           filters.propertyType.data.IsApartment,
					CreationTimeRangeFrom: filters.adDate.data.CreationTimeRangeFrom,
					CreationTimeRangeTo:   time.Now(),
				}
				u := ctx.Sender()
				user, _ := b.repo.GetUserByTelegramId(int(u.ID))
				filterModel := b.repo.CreateFilter(f)
				_, err = b.repo.CreateWatchList(Dtos.WatchListDto{UpdateCycle: m, FilterId: int(filterModel.ID), UserId: int(user.ID)})
				if err != nil {
					return err

				}
				return ctx.EditOrSend(constants.BookMarkAddedMsg)
			})

			// u := ctx.Sender()
			return nil

		default:
			filters.removeAllValue()
			return ctx.EditOrSend("/menu")
		}

	})

	return func(ctx telebot.Context) error {

		return ctx.EditOrSend(filters.message(), selector["menu"])
	}
}

func getValue(value string) []string {
	btnMap := map[string][]string{

		"PR-0":         {constants.PriceUnder500M, "0,500"},
		"PR-500":       {constants.Price500MTo700M, "500,700"},
		"PR-700":       {constants.Price700MTo900M, "700,900"},
		"PR-900":       {constants.Price900MTo1B, "900,1000"},
		"PR-1000":      {constants.Price1BTo1_5B, "1000,1500"},
		"PR-1500":      {constants.Price1_5BTo2B, "1500,2000"},
		"PR-2000":      {constants.Price2BTo3B, "2000,3000"},
		"PR-3000":      {constants.Price3BTo4B, "3000,4000"},
		"PR-4000":      {constants.Price4BTo5B, "4000,5000"},
		"PR-5000":      {constants.Price5BTo7B, "5000,7000"},
		"PR-7000":      {constants.Price7BTo10B, "7000,10000"},
		"PR-10000":     {constants.Price10BTo15B, "10000,15000"},
		"PR-15000":     {constants.Price15BTo20B, "15000,20000"},
		"PR-20000":     {constants.Price20BTo30B, "20000,30000"},
		"PR-30000":     {constants.Price30BTo40B, "30000,40000"},
		"PR-40000":     {constants.Price40BTo50B, "40000,50000"},
		"PR-50000":     {constants.Price50BTo75B, "50000,70000"},
		"PR-70000":     {constants.Price75BTo100B, "70000,100000"},
		"PR-100000":    {constants.Price100BTo200B, "100000,200000"},
		"PR-200000":    {constants.Price200BTo300B, "200000,300000"},
		"PR-300000":    {constants.Price300BTo500B, "300000,500000"},
		"PR-500000":    {constants.Price500BTo700B, "500000,700000"},
		"PR-700000":    {constants.Price700BTo900B, "700000,900000"},
		"PR-900000+":   {constants.PriceOver900B, "900000,0"},
		"at-buy":       {constants.ForBuy, "false"},
		"at-rent":      {constants.ForRent, "true"},
		"No":           {constants.No, "false"},
		"Yes":          {constants.Yes, "true"},
		"AR-0":         {constants.AreaUnder50, "0,50"},
		"AR-50":        {constants.Area50To75, "50,75"},
		"AR-75":        {constants.Area75To100, "75,100"},
		"AR-100":       {constants.Area100To150, "100,150"},
		"AR-150":       {constants.Area150To200, "150,200"},
		"AR-200":       {constants.Area200To250, "200,250"},
		"AR-250":       {constants.Area250To300, "250,300"},
		"AR-300":       {constants.Area300To400, "300,400"},
		"AR-400":       {constants.Area400To500, "400,500"},
		"AR-500":       {constants.Area500To750, "500,700"},
		"AR-750":       {constants.Area750To1000, "700,1000"},
		"AR-1000":      {constants.Area1000To1500, "1000,1500"},
		"AR-1500":      {constants.Area1500To2000, "1500,2000"},
		"AR-2000":      {constants.Area2000To3000, "2000,3000"},
		"AR-3000":      {constants.Area3000To4000, "3000,4000"},
		"AR-4000":      {constants.Area4000To5000, "4000,5000"},
		"AR-5000":      {constants.Area5000To7500, "5000,7500"},
		"AR-7500":      {constants.Area7500To10000, "75000,9000"},
		"AR-1000+":     {constants.AreaOver10000, "9000,0"},
		"NR-0":         {constants.Bedrooms0, "0,0"},
		"NR-1":         {constants.Bedrooms1, "1,1"},
		"NR-2":         {constants.Bedrooms2, "2,2"},
		"NR-3":         {constants.Bedrooms3, "3,3"},
		"NR-4":         {constants.Bedrooms4, "4,4"},
		"NR-5":         {constants.Bedrooms5, "5,5"},
		"NR-6":         {constants.Bedrooms6, "6,6"},
		"NR-7":         {constants.Bedrooms7, "7,7"},
		"NR-8":         {constants.Bedrooms8, "8,8"},
		"NR-9":         {constants.Bedrooms9, "9,9"},
		"NR-10":        {constants.Bedrooms10, "10,10"},
		"NR-10+":       {constants.BedroomsOver10, "11,11"},
		"PA-Apartment": {constants.PropertyApartment, "true"},
		"PA-Villa":     {constants.PropertyVilla, "false"},
		"BA-New":       {constants.BuildingAgeNew, "0,0"},
		"BA-1Year":     {constants.BuildingAge1Year, "1,1"},
		"BA-2Years":    {constants.BuildingAge2Years, "2,2"},
		"BA-3Years":    {constants.BuildingAge3Years, "3,3"},
		"BA-4Years":    {constants.BuildingAge4Years, "4,4"},
		"BA-5Years":    {constants.BuildingAge5Years, "5,5"},
		"BA-6Years":    {constants.BuildingAge6Years, "6,6"},
		"BA-7Years":    {constants.BuildingAge7Years, "7,7"},
		"BA-8Years":    {constants.BuildingAge8Years, "8,8"},
		"BA-9Years":    {constants.BuildingAge9Years, "9,9"},
		"BA-10Years":   {constants.BuildingAge10Years, "10,10"},
		"BA-Over10":    {constants.BuildingAgeOver10, "11,11"},
		"FLR-0":        {constants.Floor0, "0,0"},
		"FLR-1":        {constants.Floor1, "1,1"},
		"FLR-2":        {constants.Floor2, "2,2"},
		"FLR-3":        {constants.Floor3, "3,3"},
		"FLR-4":        {constants.Floor4, "4,4"},
		"FLR-5":        {constants.Floor5, "5,5"},
		"FLR-6":        {constants.Floor6, "6,6"},
		"FLR-7":        {constants.Floor7, "7,7"},
		"FLR-8":        {constants.Floor8, "8,8"},
		"FLR-9":        {constants.Floor9, "9,9"},
		"FLR-10":       {constants.Floor10, "10,10"},
		"FLR-10+":      {constants.FloorOver10, "11,11"},
		"Tehran":       {constants.Tehran, constants.Tehran},
		"Zanjan":       {constants.Zanjan, constants.Zanjan},
		"Khoram":       {constants.Khoram, constants.Khoram},
		"Mazandaran":   {constants.Mazandaran, constants.Mazandaran},
	}
	return btnMap[value]
}

func exportFileHandler(ads []models.Ad, b *Bot) func(telebot.Context) error {
	selector := newReplyMarkup()
	csvBtn := selector.Data("csv", "type", "csv")
	xlsxBtn := selector.Data("xlsx", "type", "xlsx")
	selector.Inline(selector.Row(csvBtn, xlsxBtn))
	return func(ctx telebot.Context) error {
		if ctx.Data() == "Yes" {
			b.Bot.Handle(&telebot.InlineButton{Unique: "type"}, func(ctx telebot.Context) error {
				ExportFileBot(ads, ctx.Data(), ctx)
				return ctx.Send(constants.SearchMsg)
			})
			return ctx.Send(constants.ExportTypeMsg, selector)
		} else {
			return ctx.Send(constants.SearchMsg)
		}
	}
}
