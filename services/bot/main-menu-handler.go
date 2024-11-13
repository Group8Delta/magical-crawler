package bot

import (
	"magical-crwler/constants"

	"gopkg.in/telebot.v4"
)

func MainMenuHandler(ctx telebot.Context) error {
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

	return ctx.Send(constants.MenuMsg, menu)

}
