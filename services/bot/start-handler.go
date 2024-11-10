package bot

import (
	"log"
	"magical-crwler/config"

	"gopkg.in/telebot.v4"
)

func StartHandler(b *Bot) func(ctx telebot.Context) error {
	return func(ctx telebot.Context) error {
		user := ctx.Sender()
		log.Printf("%s %s | %d started bot", user.FirstName, user.LastName, user.ID)

		err := ctx.Send(config.WelcomeMsg)
		if err != nil {
			return err
		}

		return MenuHandler(b)(ctx)
	}
}
