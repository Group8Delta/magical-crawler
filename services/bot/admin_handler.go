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

		menu.Reply(menu.Row(listAdminsBtn, removeAdminBtn, addAdminBtn))

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
		return ctx.Reply(constants.WrongUserIdFromat)
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
		return ctx.Reply(constants.WrongUserIdFromat)
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
			builder.WriteString(fmt.Sprintf("%s: %d, %s: %s %s\n", constants.AdminID, admin.ID, constants.AdminName, admin.FirstName, admin.LastName))
		}

		return ctx.Reply(builder.String())
	}
}
