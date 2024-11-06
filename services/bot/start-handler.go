package bot

import (
	"log"
	"magical-crwler/config"

	"gopkg.in/telebot.v4"
)

func StartHandler(b *Bot) func(ctx telebot.Context) error {
	return func(ctx telebot.Context) error {
		menu := &telebot.ReplyMarkup{
			ResizeKeyboard: true,
		}

		filterBtn := menu.Text(config.FiltersButton)

		menu.Reply(
			menu.Row(filterBtn),
		)

		user := ctx.Sender()
		log.Printf("%s %s | %d started bot", user.FirstName, user.LastName, user.ID)

		return ctx.Send(config.WelcomeMsg, menu)
	}
}
