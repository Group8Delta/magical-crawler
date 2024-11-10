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

		menu.Inline(
			menu.Row(searchBtn),
			menu.Row(exportBtn, bookmarkBtn),
		)

		return c.Send("یکی از گزینه های زیر را انتخاب کنید", menu)
	}
}
