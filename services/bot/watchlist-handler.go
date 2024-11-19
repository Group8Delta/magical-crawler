package bot

import (
	"encoding/json"
	"magical-crwler/database"
	"strconv"

	"gopkg.in/telebot.v4"
)

func WatchListHandler(b *Bot, db database.DbService) func(ctx telebot.Context) error {

	buttons := &telebot.ReplyMarkup{ResizeKeyboard: true}
	removeButton := buttons.Data("حذف فیلتر منتخب", "remove_watchlist")
	getButton := buttons.Data("لیست فیلتر های منتخب", "get_watchlist")

	buttons.Inline(buttons.Row(getButton, removeButton))

	return func(ctx telebot.Context) error {
		b.Bot.Handle(&removeButton, RemoveWatchListHandler(b, db))
		b.Bot.Handle(&getButton, GetWatchListHandler(b, db))

		return ctx.EditOrSend("گزینه  هارا انتخاب کنید", buttons)
	}
}

func RemoveWatchListHandler(b *Bot, db database.DbService) func(ctx telebot.Context) error {
	return func(ctx telebot.Context) error {
		ctx.EditOrSend("شناسه فیلتر منتخب را وارد کنید:")
		b.Bot.Handle(telebot.OnText, func(c telebot.Context) error {
			// Retrieve the user message
			userMessage := c.Message().Text
			id, err := strconv.Atoi(userMessage)
			if err != nil {
				return ctx.Send("عدد وارد شده نامعتبر است")

			}
			u := ctx.Sender()
			user, err := b.repo.GetUserByTelegramId(int(u.ID))
			if err != nil {
				return err

			}
			err = b.repo.DeleteWatchListByFilterId(id, int(user.ID))
			if err != nil {
				return err

			}
			return ctx.EditOrSend("حذف فیلتر منتخب با موفقیت انجام شد")

		})

		// u := ctx.Sender()
		return nil
	}
}
func GetWatchListHandler(b *Bot, db database.DbService) func(ctx telebot.Context) error {
	return func(ctx telebot.Context) error {
		u := ctx.Sender()

		filters, err := b.repo.GetWatchListFiltersByTelegramId(int(u.ID))
		if err != nil {
			return err

		}
		jfilters, _ := json.MarshalIndent(filters, "", "    ")
		return ctx.EditOrSend(string(jfilters))
	}
}
