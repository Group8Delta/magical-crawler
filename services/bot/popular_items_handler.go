package bot

import (
	"fmt"
	"magical-crwler/constants"
	"magical-crwler/database"
	"strconv"

	"strings"

	"gopkg.in/telebot.v4"
)

func PopularItemsHandler(b *Bot) func(ctx telebot.Context) error {
	return func(ctx telebot.Context) error {
		var menu = &telebot.ReplyMarkup{ResizeKeyboard: true}

		popularAdsBtn := menu.Text(constants.PopularAdsButton)
		popularSingleFiltersBtn := menu.Text(constants.PopularSingleFiltersButton)
		popularFiltersBtn := menu.Text(constants.PopularFiltersButton)

		menu.Reply(menu.Row(popularFiltersBtn, popularSingleFiltersBtn, popularAdsBtn))

		return ctx.Send(constants.PopularItemsActionMsg, menu)
	}
}

func PopularAdsHandler(b *Bot, db database.DbService) func(ctx telebot.Context) error {
	return func(ctx telebot.Context) error {
		ctx.Send(constants.NumberOfItemsQuestion)
		b.Bot.Handle(telebot.OnText, func(ctx telebot.Context) error {
			return handlePopularAds(ctx, db)
		})
		return nil
	}
}

func handlePopularAds(ctx telebot.Context, db database.DbService) error {
	n := ctx.Text()
	count, err := strconv.Atoi(n)
	if err != nil {
		return ctx.Reply(constants.WrongNumberFormat)
	}

	repo := database.NewRepository(db)

	ads, err := repo.GetMostVisitedAds(count)
	if err != nil {
		return ctx.Reply("An error occurred while getting ads.")
	}
	if len(ads) == 0 {
		return ctx.Reply(constants.EmptyAdList)
	}

	for _, ad := range ads {
		// If PhotoUrl exists, send it as an image
		if ad.PhotoUrl != nil && *ad.PhotoUrl != "" {
			photo := &telebot.Photo{
				File: telebot.FromURL(*ad.PhotoUrl),
				Caption: fmt.Sprintf(
					`Ad ID: %d
					Link: %s
					Price: %d
					Rent Price: %d
					City: %s
					Neighborhood: %s
					Size: %d sqm
					Bedrooms: %d
					For Rent: %t
					Visit Count: %d`,
					ad.ID, ad.Link, *ad.Price, ad.RentPrice, *ad.City,
					*ad.Neighborhood, *ad.Size, *ad.Bedrooms,
					ad.ForRent, ad.VisitCount,
				),
			}
			if err := ctx.Send(photo); err != nil {
				return ctx.Reply(fmt.Sprintf("Failed to send photo for ad ID %d.", ad.ID))
			}
		} else {
			if err := ctx.Reply(fmt.Sprintf(
				`Ad ID: %d
				Link: %s
				Price: %d
				Rent Price: %d
				City: %s
				Neighborhood: %s
				Size: %d sqm
				Bedrooms: %d
				For Rent: %t
				Visit Count: %d`,
				ad.ID, ad.Link, *ad.Price, ad.RentPrice, *ad.City,
				*ad.Neighborhood, *ad.Size, *ad.Bedrooms,
				ad.ForRent, ad.VisitCount,
			)); err != nil {
				return ctx.Reply(fmt.Sprintf("Failed to send text for ad ID %d.", ad.ID))
			}
		}
	}

	return nil
}

func PopularSingleFiltersHandler(b *Bot, db database.DbService) func(ctx telebot.Context) error {
	return func(ctx telebot.Context) error {
		ctx.Send(constants.NumberOfItemsQuestion)
		b.Bot.Handle(telebot.OnText, func(ctx telebot.Context) error {
			return handlePopularSingleFilters(ctx, db)
		})
		return nil
	}
}

func handlePopularSingleFilters(ctx telebot.Context, db database.DbService) error {
	n := ctx.Text()
	count, err := strconv.Atoi(n)
	if err != nil {
		return ctx.Reply(constants.WrongNumberFormat)
	}

	repo := database.NewRepository(db)

	filters, err := repo.GetMostSearchedSingleFilters(count)
	if err != nil {
		return ctx.Reply("An error occurred while getting filters.")
	}
	if len(filters) == 0 {
		return ctx.Reply(constants.EmptyFilterList)
	}

	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("%s:\n", constants.FilterList))
	for _, filter := range filters {
		builder.WriteString(fmt.Sprintf("- %s: %s (searched count: %d)\n", filter.FilterName, filter.Value, filter.Count))
	}

	return ctx.Reply(builder.String())
}

func PopularFiltersHandler(b *Bot, db database.DbService) func(ctx telebot.Context) error {
	return func(ctx telebot.Context) error {
		ctx.Send(constants.NumberOfItemsQuestion)
		b.Bot.Handle(telebot.OnText, func(ctx telebot.Context) error {
			return handlePopularFilters(ctx, db)
		})
		return nil
	}
}

func handlePopularFilters(ctx telebot.Context, db database.DbService) error {
	n := ctx.Text()
	count, err := strconv.Atoi(n)
	if err != nil {
		return ctx.Reply(constants.WrongNumberFormat)
	}

	repo := database.NewRepository(db)

	filters, err := repo.GetMostSearchedFilters(count)
	if err != nil {
		return ctx.Reply("An error occurred while getting filters.")
	}
	if len(filters) == 0 {
		return ctx.Reply(constants.EmptyFilterList)
	}

	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("%s:\n", constants.FilterList))
	for _, filter := range filters {
		builder.WriteString(fmt.Sprintf("Filter ID: %d\n", filter.ID))

		if filter.SearchQuery != nil {
			builder.WriteString(fmt.Sprintf("Search Query: %s\n", *filter.SearchQuery))
		}
		if filter.PriceRange != nil {
			builder.WriteString(fmt.Sprintf("Price Range: %d - %d\n", filter.PriceRange.Min, filter.PriceRange.Max))
		}
		if filter.RentPriceRange != nil {
			builder.WriteString(fmt.Sprintf("Rent Price Range: %d - %d\n", filter.RentPriceRange.Min, filter.RentPriceRange.Max))
		}
		builder.WriteString(fmt.Sprintf("For Rent: %t\n", filter.ForRent))
		if filter.City != nil {
			builder.WriteString(fmt.Sprintf("City: %s\n", *filter.City))
		}
		if filter.Neighborhood != nil {
			builder.WriteString(fmt.Sprintf("Neighborhood: %s\n", *filter.Neighborhood))
		}
		if filter.SizeRange != nil {
			builder.WriteString(fmt.Sprintf("Size Range: %d - %d sqm\n", filter.SizeRange.Min, filter.SizeRange.Max))
		}
		if filter.BedroomRange != nil {
			builder.WriteString(fmt.Sprintf("Bedroom Range: %d - %d\n", filter.BedroomRange.Min, filter.BedroomRange.Max))
		}
		if filter.FloorRange != nil {
			builder.WriteString(fmt.Sprintf("Floor Range: %d - %d\n", filter.FloorRange.Min, filter.FloorRange.Max))
		}
		if filter.HasElevator != nil {
			builder.WriteString(fmt.Sprintf("Has Elevator: %t\n", *filter.HasElevator))
		}
		if filter.HasStorage != nil {
			builder.WriteString(fmt.Sprintf("Has Storage: %t\n", *filter.HasStorage))
		}
		if filter.AgeRange != nil {
			builder.WriteString(fmt.Sprintf("Age Range: %d - %d years\n", filter.AgeRange.Min, filter.AgeRange.Max))
		}
		if filter.IsApartment != nil {
			builder.WriteString(fmt.Sprintf("Is Apartment: %t\n", *filter.IsApartment))
		}
		if !filter.CreationTimeRangeFrom.IsZero() {
			builder.WriteString(fmt.Sprintf("Created From: %s\n", filter.CreationTimeRangeFrom.Format("2006-01-02")))
		}
		if !filter.CreationTimeRangeTo.IsZero() {
			builder.WriteString(fmt.Sprintf("Created To: %s\n", filter.CreationTimeRangeTo.Format("2006-01-02")))
		}
		builder.WriteString(fmt.Sprintf("Searched Count: %d\n\n", filter.SearchedCount))
	}

	return ctx.Reply(builder.String())
}
