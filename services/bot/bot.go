package bot

import (
	"log"
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

func (b *Bot) StartBot() {
	log.Print("Bot is running !")
	// Start(b)
	b.bot.Start()
}
