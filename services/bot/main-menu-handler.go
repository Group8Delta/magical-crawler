package bot

import (
	"magical-crwler/constants"
	"magical-crwler/models"

	"gopkg.in/telebot.v4"
	"gorm.io/gorm"
)

func MainMenuHandler(ctx telebot.Context, db *gorm.DB, user *models.User) error {
	var (
		menu = &telebot.ReplyMarkup{ResizeKeyboard: true}

		searchBtn       = menu.Text(constants.SearchButton)
		filtersBtn      = menu.Text(constants.FiltersButton)
		accountBtn      = menu.Text(constants.AccountManagementButton)
		exportBtn       = menu.Text(constants.ExportButton)
		bookmarkBtn     = menu.Text(constants.FavoritesButton)
		adminPnlBtn     = menu.Text(constants.AdminPanelButton)
		popularItemsBtn = menu.Text(constants.PopularItemsButton)
	)

	if models.IsSuperAdmin(db, user.ID) {
		menu.Reply(
			menu.Row(searchBtn, filtersBtn),
			menu.Row(exportBtn, bookmarkBtn),
			menu.Row(accountBtn, popularItemsBtn),
			menu.Row(adminPnlBtn),
		)
	} else {
		menu.Reply(
			menu.Row(searchBtn, filtersBtn),
			menu.Row(exportBtn, bookmarkBtn),
			menu.Row(accountBtn, popularItemsBtn),
		)
	}

	return ctx.Send(constants.MenuMsg, menu)

}
