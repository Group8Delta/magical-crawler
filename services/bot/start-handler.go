package bot

import (
	"log"
	"magical-crwler/config"

	"gopkg.in/telebot.v4"
)

func StartHandler(b *Bot) func(ctx telebot.Context) error {
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
		)

		menu.Reply(
			menu.Row(searchBtn, filtersBtn),
			menu.Row(exportBtn, bookmarkBtn),
			menu.Row(accountBtn),
		)

		return ctx.Send(config.WelcomeMsg, menu)
	}
}
