package bot

import "magical-crwler/constants"

func RegisterHandlers(b *Bot) {
	b.bot.Handle("/start", StartHandler(b))
	b.bot.Handle(constants.SearchButton, FilterHandlers(b))
	b.bot.Handle("/exportFile", ExportHandler(b))
}
