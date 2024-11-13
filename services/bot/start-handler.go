package bot

import (
	"log"
	"magical-crwler/constants"
	"magical-crwler/models"

	"gopkg.in/telebot.v4"
	"gorm.io/gorm"
)

func StartHandler(b *Bot, db *gorm.DB) func(ctx telebot.Context) error {
	return func(ctx telebot.Context) error {
		user := ctx.Sender()
		log.Printf("%s %s | %d started bot", user.FirstName, user.LastName, user.ID)
		ctx.Send(constants.WelcomeMsg)
		return MainMenuHandler(ctx)

		var (
			menu = &telebot.ReplyMarkup{ResizeKeyboard: true}

			searchBtn   = menu.Text(constants.SearchButton)
			filtersBtn  = menu.Text(constants.FiltersButton)
			accountBtn  = menu.Text(constants.AccountManagementButton)
			exportBtn   = menu.Text(constants.ExportButton)
			bookmarkBtn = menu.Text(constants.FavoritesButton)
			adminPnlBtn = menu.Text(constants.AdminPanelButton)
		)
		//TODO: move to main
		if models.IsSuperAdmin(db, user.ID) {
			menu.Reply(
				menu.Row(searchBtn, filtersBtn),
				menu.Row(exportBtn, bookmarkBtn),
				menu.Row(accountBtn, adminPnlBtn),
			)
		} else {
			menu.Reply(
				menu.Row(searchBtn, filtersBtn),
				menu.Row(exportBtn, bookmarkBtn),
				menu.Row(accountBtn),
			)
		}

		return ctx.Send(constants.WelcomeMsg, menu)
	}
}
