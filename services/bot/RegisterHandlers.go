package bot

import "magical-crwler/config"

func RegisterHanlders(b *Bot) {
	b.bot.Handle("/menu", MenuHandler(b))
	b.bot.Handle("/start", StartHandler(b))
	b.bot.Handle(config.FiltersButton, FilterHandlers(b))
	b.bot.Handle("/exportFile", ExportHandler(b))
}
