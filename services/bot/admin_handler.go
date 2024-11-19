package bot

import (
	"fmt"
	"log"
	"magical-crwler/constants"
	"magical-crwler/services/admin"
	"strconv"
	"strings"

	"gopkg.in/telebot.v4"
	"gorm.io/gorm"
)

func AdminHandler(b *Bot) func(ctx telebot.Context) error {
	return func(ctx telebot.Context) error {
		var menu = &telebot.ReplyMarkup{ResizeKeyboard: true}

		addAdminBtn := menu.Text(constants.AddAdminButton)
		removeAdminBtn := menu.Text(constants.RemoveAdminButton)
		listAdminsBtn := menu.Text(constants.ListAdminsButton)
		userListBtn := menu.Text(constants.UserList)
		crawlerStatusBtn := menu.Text(constants.CrawlerStatusButton)
		listCrawlInfoBtn := menu.Text(constants.ListCrawlInfoButton)
		crawlInfoBtn := menu.Text(constants.CrawlInfoButton)

		menu.Reply(
			menu.Row(removeAdminBtn, addAdminBtn),
			menu.Row(listAdminsBtn, userListBtn),
			menu.Row(listCrawlInfoBtn, crawlInfoBtn),
			menu.Row(crawlerStatusBtn),
		)

		return ctx.Send(constants.AdminActionMsg, menu)
	}
}

func AddAdminHandler(b *Bot, db *gorm.DB) func(ctx telebot.Context) error {
	return func(ctx telebot.Context) error {
		ctx.Send(constants.AddAdminQuestion)
		b.Bot.Handle(telebot.OnText, func(ctx telebot.Context) error {
			return handleAddAdmin(ctx, db)
		})
		return nil
	}
}

func handleAddAdmin(ctx telebot.Context, db *gorm.DB) error {
	userInput := ctx.Text()
	userID, err := strconv.ParseInt(userInput, 10, 64)
	if err != nil {
		log.Println("Error user ID:", userID)
		return ctx.Reply(constants.WrongUserIdFormat)
	}

	adminService := admin.NewAdminService(db)

	err = adminService.AddAdmin(userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Reply(constants.UserNotFound)
		}
		return ctx.Reply("An error occurred while updating the user role.")
	}
	return ctx.Reply(constants.AdminAddedMsg)
}

func RemoveAdminHandler(b *Bot, db *gorm.DB) func(ctx telebot.Context) error {
	return func(ctx telebot.Context) error {
		ctx.Send(constants.RemoveAdminQuestion)
		b.Bot.Handle(telebot.OnText, func(ctx telebot.Context) error {
			return handleRemoveAdmin(ctx, db)
		})
		return nil
	}
}

func handleRemoveAdmin(ctx telebot.Context, db *gorm.DB) error {
	userInput := ctx.Text()
	userID, err := strconv.ParseInt(userInput, 10, 64)
	if err != nil {
		log.Println("Error user ID:", userID)
		return ctx.Reply(constants.WrongUserIdFormat)
	}

	adminService := admin.NewAdminService(db)

	err = adminService.RemoveAdmin(userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Reply(constants.AdminNotFound)
		}
		return ctx.Reply("An error occurred while updating the user role.")
	}
	return ctx.Reply(constants.AdminRemovedMsg)
}

func AdminListHandler(b *Bot, db *gorm.DB) func(ctx telebot.Context) error {
	return func(ctx telebot.Context) error {
		adminService := admin.NewAdminService(db)
		admins, err := adminService.ListAdmins()
		if err != nil {
			log.Println("Error retrieving admin list:", err)
			return ctx.Reply("An error occurred while retrieving the admin list.")
		}

		if len(admins) == 0 {
			return ctx.Reply(constants.EmptyAdminList)
		}

		var builder strings.Builder
		builder.WriteString(fmt.Sprintf("%s:\n", constants.AdminList))
		for _, admin := range admins {
			builder.WriteString(fmt.Sprintf("%s: %d, %s: %s %s\n", constants.UserID, admin.ID, constants.UserName, admin.FirstName, admin.LastName))
		}

		return ctx.Reply(builder.String())
	}
}

func UserListHandler(b *Bot, db *gorm.DB) func(ctx telebot.Context) error {
	return func(ctx telebot.Context) error {
		adminService := admin.NewAdminService(db)
		users, err := adminService.ListUsers()
		if err != nil {
			log.Println("Error retrieving admin list:", err)
			return ctx.Reply("An error occurred while retrieving the admin list.")
		}

		if len(users) == 0 {
			return ctx.Reply(constants.EmptyUserList)
		}

		var builder strings.Builder
		builder.WriteString(fmt.Sprintf("%s:\n", constants.UserList))
		for _, user := range users {
			builder.WriteString(fmt.Sprintf("%s: %d, %s: %s %s\n", constants.UserID, user.ID, constants.UserName, user.FirstName, user.LastName))
		}

		return ctx.Reply(builder.String())
	}
}

func CrawlerStatusLogs(b *Bot, db *gorm.DB) func(ctx telebot.Context) error {
	return func(ctx telebot.Context) error {
		adminService := admin.NewAdminService(db)
		logs, err := adminService.ListCrawlerStausLogs()
		if err != nil {
			log.Println("Error retrieving admin list:", err)
			return ctx.Reply("An error occurred while retrieving the admin list.")
		}

		if len(logs) == 0 {
			return ctx.Reply(constants.EmptyCrawlerStatusList)
		}

		var builder strings.Builder
		builder.WriteString(fmt.Sprintf("%s:\n", constants.CrawlerStatusList))
		for _, log := range logs {
			builder.WriteString(fmt.Sprintf("%s duration: %s, cpu: %s, ram: %sM, numer of request: %s , successful crawl: %s, failed crawl: %s\n", log.Date, log.Duration, log.CPUUsage, log.RAMUsage, log.TotalRequests, log.SuccessfulRequests, log.FailedRequests))
		}

		return ctx.Reply(builder.String())
	}
}

func UsersCrawlInfoHandler(b *Bot, db *gorm.DB) func(ctx telebot.Context) error {
	return func(ctx telebot.Context) error {
		adminService := admin.NewAdminService(db)

		filtersWithAds, err := adminService.GetUsersCrawlInfo()
		if err != nil {
			log.Println("Error retrieving users' crawl info:", err)
			return ctx.Reply("An error occurred while retrieving users' crawl info.")
		}

		if len(filtersWithAds) == 0 {
			return ctx.Reply(constants.EmptyUserCrawlInfoList)
		}

		var builder strings.Builder
		builder.WriteString(fmt.Sprintf("%s:\n", constants.UserCrawlInfoList))

		for _, filterWithAds := range filtersWithAds {
			builder.WriteString(fmt.Sprintf("Filter ID: %d\n", filterWithAds.FilterID))

			if filterWithAds.Filter.SearchQuery != nil {
				builder.WriteString(fmt.Sprintf("Search Query: %s\n", *filterWithAds.Filter.SearchQuery))
			}
			if filterWithAds.Filter.PriceRange != nil {
				builder.WriteString(fmt.Sprintf("Price Range: %d - %d\n", filterWithAds.Filter.PriceRange.Min, filterWithAds.Filter.PriceRange.Max))
			}
			if filterWithAds.Filter.RentPriceRange != nil {
				builder.WriteString(fmt.Sprintf("Rent Price Range: %d - %d\n", filterWithAds.Filter.RentPriceRange.Min, filterWithAds.Filter.RentPriceRange.Max))
			}
			builder.WriteString(fmt.Sprintf("For Rent: %t\n", filterWithAds.Filter.ForRent))
			if filterWithAds.Filter.City != nil {
				builder.WriteString(fmt.Sprintf("City: %s\n", *filterWithAds.Filter.City))
			}
			if filterWithAds.Filter.Neighborhood != nil {
				builder.WriteString(fmt.Sprintf("Neighborhood: %s\n", *filterWithAds.Filter.Neighborhood))
			}
			if filterWithAds.Filter.SizeRange != nil {
				builder.WriteString(fmt.Sprintf("Size Range: %d - %d sqm\n", filterWithAds.Filter.SizeRange.Min, filterWithAds.Filter.SizeRange.Max))
			}
			if filterWithAds.Filter.BedroomRange != nil {
				builder.WriteString(fmt.Sprintf("Bedroom Range: %d - %d\n", filterWithAds.Filter.BedroomRange.Min, filterWithAds.Filter.BedroomRange.Max))
			}
			if filterWithAds.Filter.FloorRange != nil {
				builder.WriteString(fmt.Sprintf("Floor Range: %d - %d\n", filterWithAds.Filter.FloorRange.Min, filterWithAds.Filter.FloorRange.Max))
			}
			if filterWithAds.Filter.HasElevator != nil {
				builder.WriteString(fmt.Sprintf("Has Elevator: %t\n", *filterWithAds.Filter.HasElevator))
			}
			if filterWithAds.Filter.HasStorage != nil {
				builder.WriteString(fmt.Sprintf("Has Storage: %t\n", *filterWithAds.Filter.HasStorage))
			}
			if filterWithAds.Filter.AgeRange != nil {
				builder.WriteString(fmt.Sprintf("Age Range: %d - %d years\n", filterWithAds.Filter.AgeRange.Min, filterWithAds.Filter.AgeRange.Max))
			}
			if filterWithAds.Filter.IsApartment != nil {
				builder.WriteString(fmt.Sprintf("Is Apartment: %t\n", *filterWithAds.Filter.IsApartment))
			}
			if !filterWithAds.Filter.CreationTimeRangeFrom.IsZero() {
				builder.WriteString(fmt.Sprintf("Created From: %s\n", filterWithAds.Filter.CreationTimeRangeFrom.Format("2006-01-02")))
			}
			if !filterWithAds.Filter.CreationTimeRangeTo.IsZero() {
				builder.WriteString(fmt.Sprintf("Created To: %s\n", filterWithAds.Filter.CreationTimeRangeTo.Format("2006-01-02")))
			}
			builder.WriteString(fmt.Sprintf("Searched Count: %d\n\n", filterWithAds.Filter.SearchedCount))

			builder.WriteString(fmt.Sprintf("- Ads: %d\n", len(filterWithAds.Ads)))

			for _, ad := range filterWithAds.Ads {
				builder.WriteString(fmt.Sprintf("  * Ad Link: %s\n", ad.Link))
				if ad.Price != nil {
					builder.WriteString(fmt.Sprintf("    Price: %d\n", *ad.Price))
				}
				if ad.Description != nil {
					builder.WriteString(fmt.Sprintf("    Description: %s\n", *ad.Description))
				}
				if ad.PhotoUrl != nil {
					builder.WriteString(fmt.Sprintf("    Photo: %s\n", *ad.PhotoUrl))
				}
			}
			builder.WriteString("\n")
		}

		return ctx.Reply(builder.String())
	}
}

func SingleUserCrawlInfoHandler(b *Bot, db *gorm.DB) func(ctx telebot.Context) error {
	return func(ctx telebot.Context) error {
		ctx.Send(constants.UserCrawlInfoQuestion)
		b.Bot.Handle(telebot.OnText, func(ctx telebot.Context) error {
			return HandleSingleUserCrawlInfo(ctx, db)
		})
		return nil
	}
}

func HandleSingleUserCrawlInfo(ctx telebot.Context, db *gorm.DB) error {
	userInput := ctx.Text()
	userID, err := strconv.ParseInt(userInput, 10, 64)
	if err != nil {
		log.Println("Error user ID:", userID)
		return ctx.Reply(constants.WrongUserIdFormat)
	}

	adminService := admin.NewAdminService(db)

	userCrawlInfo, err := adminService.GetSingleUserCrawlInfo(userID)
	if err != nil {
		log.Println("Error retrieving users' crawl info:", err)
		return ctx.Reply("An error occurred while retrieving users' crawl info.")
	}

	if len(userCrawlInfo.Filters) == 0 {
		return ctx.Reply(constants.EmptyUserCrawlInfoList)
	}

	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("%s:\n", constants.UserCrawlInfoList))

	for _, filterWithAds := range userCrawlInfo.Filters {
		builder.WriteString(fmt.Sprintf("Filter ID: %d\n", filterWithAds.FilterID))

		if filterWithAds.Filter.SearchQuery != nil {
			builder.WriteString(fmt.Sprintf("Search Query: %s\n", *filterWithAds.Filter.SearchQuery))
		}
		if filterWithAds.Filter.PriceRange != nil {
			builder.WriteString(fmt.Sprintf("Price Range: %d - %d\n", filterWithAds.Filter.PriceRange.Min, filterWithAds.Filter.PriceRange.Max))
		}
		if filterWithAds.Filter.RentPriceRange != nil {
			builder.WriteString(fmt.Sprintf("Rent Price Range: %d - %d\n", filterWithAds.Filter.RentPriceRange.Min, filterWithAds.Filter.RentPriceRange.Max))
		}
		builder.WriteString(fmt.Sprintf("For Rent: %t\n", filterWithAds.Filter.ForRent))
		if filterWithAds.Filter.City != nil {
			builder.WriteString(fmt.Sprintf("City: %s\n", *filterWithAds.Filter.City))
		}
		if filterWithAds.Filter.Neighborhood != nil {
			builder.WriteString(fmt.Sprintf("Neighborhood: %s\n", *filterWithAds.Filter.Neighborhood))
		}
		if filterWithAds.Filter.SizeRange != nil {
			builder.WriteString(fmt.Sprintf("Size Range: %d - %d sqm\n", filterWithAds.Filter.SizeRange.Min, filterWithAds.Filter.SizeRange.Max))
		}
		if filterWithAds.Filter.BedroomRange != nil {
			builder.WriteString(fmt.Sprintf("Bedroom Range: %d - %d\n", filterWithAds.Filter.BedroomRange.Min, filterWithAds.Filter.BedroomRange.Max))
		}
		if filterWithAds.Filter.FloorRange != nil {
			builder.WriteString(fmt.Sprintf("Floor Range: %d - %d\n", filterWithAds.Filter.FloorRange.Min, filterWithAds.Filter.FloorRange.Max))
		}
		if filterWithAds.Filter.HasElevator != nil {
			builder.WriteString(fmt.Sprintf("Has Elevator: %t\n", *filterWithAds.Filter.HasElevator))
		}
		if filterWithAds.Filter.HasStorage != nil {
			builder.WriteString(fmt.Sprintf("Has Storage: %t\n", *filterWithAds.Filter.HasStorage))
		}
		if filterWithAds.Filter.AgeRange != nil {
			builder.WriteString(fmt.Sprintf("Age Range: %d - %d years\n", filterWithAds.Filter.AgeRange.Min, filterWithAds.Filter.AgeRange.Max))
		}
		if filterWithAds.Filter.IsApartment != nil {
			builder.WriteString(fmt.Sprintf("Is Apartment: %t\n", *filterWithAds.Filter.IsApartment))
		}
		if !filterWithAds.Filter.CreationTimeRangeFrom.IsZero() {
			builder.WriteString(fmt.Sprintf("Created From: %s\n", filterWithAds.Filter.CreationTimeRangeFrom.Format("2006-01-02")))
		}
		if !filterWithAds.Filter.CreationTimeRangeTo.IsZero() {
			builder.WriteString(fmt.Sprintf("Created To: %s\n", filterWithAds.Filter.CreationTimeRangeTo.Format("2006-01-02")))
		}
		builder.WriteString(fmt.Sprintf("Searched Count: %d\n\n", filterWithAds.Filter.SearchedCount))

		builder.WriteString(fmt.Sprintf("- Ads: %d\n", len(filterWithAds.Ads)))

		for _, ad := range filterWithAds.Ads {
			builder.WriteString(fmt.Sprintf("  * Ad Link: %s\n", ad.Link))
			if ad.Price != nil {
				builder.WriteString(fmt.Sprintf("    Price: %d\n", *ad.Price))
			}
			if ad.Description != nil {
				builder.WriteString(fmt.Sprintf("    Description: %s\n", *ad.Description))
			}
			if ad.PhotoUrl != nil {
				builder.WriteString(fmt.Sprintf("    Photo: %s\n", *ad.PhotoUrl))
			}
		}
		builder.WriteString("\n")
	}

	return ctx.Reply(builder.String())
}
