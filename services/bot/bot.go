package bot

import (
	"log"
	"magical-crwler/config"
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
	b.bot.Handle("/start", StartHandler(b))
	b.bot.Handle(config.FiltersButton, FilterHandlers(b))
	b.bot.Start()
}
