package bot

import (
	"fmt"
	"log"
	"magical-crwler/services/notification"
	"strconv"
	"time"

	"gopkg.in/telebot.v4"
)

type Bot struct {
	bot    *telebot.Bot
	config BotConfig
}
type BotConfig struct {
	Token  string
	Poller time.Duration
}

func NewBot(config BotConfig) (*Bot, error) {
	bot, err := telebot.NewBot(telebot.Settings{
		Token:  config.Token,
		Poller: &telebot.LongPoller{Timeout: config.Poller},
	})
	if err != nil {
		log.Fatalln(err.Error())
	}
	return &Bot{
		bot:    bot,
		config: config,
	}, nil
}

func (b *Bot) Notify(recipientIdentifier string, m *notification.Message) error {

	teleUserId, err := strconv.Atoi(recipientIdentifier)
	if err != nil {
		return err
	}
	recipient := telebot.ChatID(teleUserId)
	message := fmt.Sprintf("%s \n\n %s", m.Title, m.Content)
	_, err = b.bot.Send(recipient, message)
	if err != nil {
		return err
	}

	return nil

}

func (b *Bot) StartBot() {
	log.Print("Bot is running !")
	RegisterHanlders(b)

	// b.bot.Handle("/start", StartHandler(b))
	// b.bot.Handle(config.FiltersButton, FilterHandlers(b))
	b.bot.Start()
}
