package bot

import (
	"log"
	"magical-crwler/constants"

	"gopkg.in/telebot.v4"
)

func StartHandler(b *Bot) func(ctx telebot.Context) error {
	return func(ctx telebot.Context) error {
		user := ctx.Sender()
		log.Printf("%s %s | %d started bot", user.FirstName, user.LastName, user.ID)
		ctx.Send(constants.WelcomeMsg)
		return MainMenuHandler(ctx)
	}
}
