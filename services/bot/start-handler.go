package bot

import (
	"log"
	"magical-crwler/config"
	"magical-crwler/models"

	"gopkg.in/telebot.v4"
	"gorm.io/gorm"
)

func StartHandler(b *Bot, db *gorm.DB) func(ctx telebot.Context) error {
	return func(ctx telebot.Context) error {
		user := ctx.Sender()
		log.Printf("%s %s | %d started bot", user.FirstName, user.LastName, user.ID)

		var (
			menu = &telebot.ReplyMarkup{ResizeKeyboard: true}

			searchBtn   = menu.Text(config.SearchButton)
			filtersBtn  = menu.Text(config.FiltersButton)
			accountBtn  = menu.Text(config.AccountManagementButton)
			exportBtn   = menu.Text(config.ExportButton)
			bookmarkBtn = menu.Text(config.FavoritesButton)
			adminPnlBtn = menu.Text(config.AdminPanelButton)
		)

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

		return ctx.Send(config.WelcomeMsg, menu)
	}
}
