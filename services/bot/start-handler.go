package bot

import (
	"log"
	"magical-crwler/constants"

	"gopkg.in/telebot.v4"
)

func StartHandler(b *Bot) func(ctx telebot.Context) error {
	return func(ctx telebot.Context) error {
		user := ctx.Sender()
		log.Printf("%s %s | %d started bot", user.FirstName, user.LastName, user.ID)

		var (
			menu = &telebot.ReplyMarkup{ResizeKeyboard: true}

			searchBtn   = menu.Text(constants.SearchButton)
			filtersBtn  = menu.Text(constants.FiltersButton)
			accountBtn  = menu.Text(constants.AccountManagementButton)
			exportBtn   = menu.Text(constants.ExportButton)
			bookmarkBtn = menu.Text(constants.FavoritesButton)
		)

		menu.Reply(
			menu.Row(searchBtn, filtersBtn),
			menu.Row(exportBtn, bookmarkBtn),
			menu.Row(accountBtn),
		)

		return ctx.Send(constants.WelcomeMsg, menu)
	}
}
