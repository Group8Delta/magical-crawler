package bot

import "magical-crwler/config"

func RegisterHanlders(b *Bot) {
	b.bot.Handle("/start", StartHandler(b))
	b.bot.Handle(config.SearchButton, FilterHandlers(b))
	b.bot.Handle("/exportFile", ExportHandler(b))
}
