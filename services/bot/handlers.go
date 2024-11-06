package bot

import (
	"log"
	"magical-crwler/config"

	"gopkg.in/telebot.v4"
)

func StartHandler(ctx telebot.Context) error {
	menu := &telebot.ReplyMarkup{
		ResizeKeyboard: true,
	}

	filterBtn := menu.Text(config.FiltersButton)

	menu.Reply(
		menu.Row(filterBtn),
	)

	user := NewUser(*ctx.Sender())
	log.Printf("%s %s | %d started bot", user.info.FirstName, user.info.LastName, user.info.ID)

	return ctx.Send(config.WelcomeMsg, menu)
}

func FilterHandlers(ctx telebot.Context) error {
	selector := &telebot.ReplyMarkup{RemoveKeyboard: true}
	menu := &telebot.ReplyMarkup{}

	applyFiltersButton := menu.Text(config.ApplyFiltersButton)
	removeFiltersButton := menu.Text(config.RemoveFiltersButton)

	priceFilterSelector := selector.Data(config.PriceFilter, "PriceFilter")
	areaFilterSelector := selector.Data(config.AreaFilter, "AreaFilter")
	roomsFilterSelector := selector.Data(config.RoomsFilter, "RoomsFilter")
	propertyTypeFilterSelector := selector.Data(config.PropertyTypeFilter, "PropertyTypeFilter")
	buildingAgeFilterSelector := selector.Data(config.BuildingAgeFilter, "BuildingAgeFilter")
	floorFilterSelector := selector.Data(config.FloorFilter, "FloorFilter")
	storageFilterSelector := selector.Data(config.StorageFilter, "StorageFilter")
	elevatorFilterSelector := selector.Data(config.ElevatorFilter, "ElevatorFilter")
	adDateFilterSelector := selector.Data(config.AdDateFilter, "AdDateFilter")
	locationFilterSelector := selector.Data(config.LocationFilter, "LocationFilter")

	selector.Inline(
		selector.Row(priceFilterSelector),
		selector.Row(areaFilterSelector, roomsFilterSelector, propertyTypeFilterSelector),
		selector.Row(buildingAgeFilterSelector, floorFilterSelector, storageFilterSelector),
		selector.Row(elevatorFilterSelector, adDateFilterSelector, locationFilterSelector),
	)

	menu.Reply(
		menu.Row(applyFiltersButton, removeFiltersButton),
	)
	return ctx.Send("فیلتر ها", selector)
}
