package bot

import (
	"magical-crwler/config"

	"gopkg.in/telebot.v4"
)

func MenuHandler(b *Bot) func(c telebot.Context) error {
	return func(c telebot.Context) error {
		menu := &telebot.ReplyMarkup{}

		searchBtn := menu.Data(config.SearchButton, "search")
		exportBtn := menu.Data(config.ExportButton, "export")
		bookmarkBtn := menu.Data(config.FavoritesButton, "bookmark")
		accountManageBtn := menu.Data(config.AccountManagementButton, "account")

		menu.Inline(
			menu.Row(searchBtn, accountManageBtn),
			menu.Row(exportBtn, bookmarkBtn),
		)

		return c.EditOrSend("یکی از گزینه های زیر را انتخاب کنید", menu)
	}
}
