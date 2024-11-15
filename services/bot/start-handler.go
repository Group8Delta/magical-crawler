package bot

import (
	"log"
	"magical-crwler/constants"
	"magical-crwler/models"

	"gopkg.in/telebot.v4"
	"gorm.io/gorm"
)

func StartHandler(b *Bot, db *gorm.DB) func(ctx telebot.Context) error {
	return func(ctx telebot.Context) error {
		telUser := ctx.Sender()
		log.Printf("%s %s | %d started bot", telUser.FirstName, telUser.LastName, telUser.ID)

		user, err := models.FindOrCreateUser(db, uint(telUser.ID), telUser.FirstName, telUser.LastName)
		if err != nil {
			return ctx.Reply("An error occurred while accessing the database.")
		}

		ctx.Send(constants.WelcomeMsg)
		return MainMenuHandler(ctx, db, user)
	}
}
