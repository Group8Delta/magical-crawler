package bot

import (
	"fmt"
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
type Route struct {
	command string
	handler func(telebot.Context) error
}

var routes []Route

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
	for _, r := range routes {

		b.bot.Handle(fmt.Sprintf("/%s", r.command), r.handler)
	}
	log.Print("Bot is running !")
	b.bot.Start()
}
