package bot

import (
	"log"
	"magical-crwler/constants"
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

func (b *Bot) RegisterHandlers() {
	b.bot.Handle("/start", StartHandler(b))
	b.bot.Handle(constants.SearchButton, FilterHandlers(b))
	b.bot.Handle("/exportFile", ExportHandler(b))
}

func (b *Bot) StartBot() {
	log.Print("Bot is running !")
	b.RegisterHandlers()
	b.bot.Start()
}
