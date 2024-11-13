package bot

import (
	"log"
	"magical-crwler/config"
	"magical-crwler/services/admin"
	"strconv"

	"gopkg.in/telebot.v4"
	"gorm.io/gorm"
)

func AdminHandler(b *Bot) func(ctx telebot.Context) error {
	return func(ctx telebot.Context) error {
		var menu = &telebot.ReplyMarkup{ResizeKeyboard: true}

		addAdminBtn := menu.Text(config.AddAdminButton)
		removeAdminBtn := menu.Text(config.RemoveAdminButton)

		menu.Reply(menu.Row(addAdminBtn, removeAdminBtn))

		return ctx.Send(config.AdminActionMsg, menu)
	}
}

func AddAdminHandler(b *Bot, db *gorm.DB) func(ctx telebot.Context) error {
	return func(ctx telebot.Context) error {
		ctx.Send(config.AddAdminQuestion)
		b.bot.Handle(telebot.OnText, func(ctx telebot.Context) error {
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
		return ctx.Reply(config.WrongUserIdFromat)
	}

	adminService := admin.NewAdminService(db)

	err = adminService.AddAdmin(userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Reply(config.UserNotFound)
		}
		return ctx.Reply("An error occurred while updating the user role.")
	}
	return ctx.Reply(config.AdminAddedMsg)
}

func RemoveAdminHandler(b *Bot, db *gorm.DB) func(ctx telebot.Context) error {
	return func(ctx telebot.Context) error {
		ctx.Send(config.RemoveAdminQuestion)
		b.bot.Handle(telebot.OnText, func(ctx telebot.Context) error {
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
		return ctx.Reply(config.WrongUserIdFromat)
	}

	adminService := admin.NewAdminService(db)

	err = adminService.RemoveAdmin(userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Reply(config.AdminNotFound)
		}
		return ctx.Reply("An error occurred while updating the user role.")
	}
	return ctx.Reply(config.AdminRemovedMsg)
}
