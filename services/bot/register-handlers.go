package bot

import (
	"magical-crwler/constants"

	"gorm.io/gorm"
)

func RegisterHanlders(b *Bot, db *gorm.DB) {
	b.Bot.Handle("/start", StartHandler(b, db))
	b.Bot.Handle(constants.SearchButton, SearchHandlers(b))
	b.Bot.Handle(constants.AdminPanelButton, AdminHandler(b))
	b.Bot.Handle(constants.AddAdminButton, AddAdminHandler(b, db))
	b.Bot.Handle(constants.RemoveAdminButton, RemoveAdminHandler(b, db))
	b.Bot.Handle("/exportFile", ExportHandler(b))
}
