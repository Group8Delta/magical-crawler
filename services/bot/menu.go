package bot

// tasmim giri beshe

// import (
// 	"magical-crwler/constants"

// 	"gopkg.in/telebot.v4"
// )

// func MenuHandler(b *Bot) func(c telebot.Context) error {
// 	return func(c telebot.Context) error {
// 		menu := &telebot.ReplyMarkup{}

// 		searchBtn := menu.Data(constants.SearchButton, "search")
// 		exportBtn := menu.Data(constants.ExportButton, "export")
// 		bookmarkBtn := menu.Data(constants.FavoritesButton, "bookmark")
// 		accountManageBtn := menu.Data(constants.AccountManagementButton, "account")

// 		menu.Inline(
// 			menu.Row(searchBtn, accountManageBtn),
// 			menu.Row(exportBtn, bookmarkBtn),
// 		)

// 		return c.EditOrSend("یکی از گزینه های زیر را انتخاب کنید", menu)
// 	}
// }
