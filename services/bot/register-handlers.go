package bot

import (
	"magical-crwler/config"

	"gorm.io/gorm"
)

func RegisterHanlders(b *Bot, db *gorm.DB) {
	b.bot.Handle("/start", StartHandler(b, db))
	b.bot.Handle(config.SearchButton, FilterHandlers(b))
	b.bot.Handle(config.AdminPanelButton, AdminHandler(b))
	b.bot.Handle(config.AddAdminButton, AddAdminHandler(b, db))
	b.bot.Handle(config.RemoveAdminButton, RemoveAdminHandler(b, db))
	b.bot.Handle("/exportFile", ExportHandler(b))
}
