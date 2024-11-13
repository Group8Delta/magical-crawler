package bot

import (
	"log"
	"magical-crwler/services/notification"
	"time"

	"gopkg.in/telebot.v4"
	"gorm.io/gorm"
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
	return nil
}

func (b *Bot) StartBot(db *gorm.DB) {
	log.Print("Bot is running !")
	RegisterHanlders(b, db)

	// b.bot.Handle("/start", StartHandler(b))
	// b.bot.Handle(config.FiltersButton, FilterHandlers(b))
	b.bot.Start()
}
