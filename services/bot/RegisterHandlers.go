package bot

import (
	"magical-crwler/config"

	"gopkg.in/telebot.v4"
)

func RegisterHanlders(b *Bot) {
	b.bot.Handle("/menu", MenuHandler(b))
	b.bot.Handle("/start", StartHandler(b))
	b.bot.Handle(&telebot.Btn{Unique: "export"}, ExportHandler(b))
	b.bot.Handle(config.FiltersButton, FilterHandlers(b))
	b.bot.Handle("/exportFile", ExportHandler(b))
	b.bot.Handle(&telebot.Btn{Unique: "export"}, ExportHandler(b))
	b.bot.Handle(&telebot.Btn{Unique: "export_csv"}, export_csv_Handler(b))
	b.bot.Handle(&telebot.Btn{Unique: "export_xlsx"}, export_xlsx_Handler(b))
}
