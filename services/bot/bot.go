package bot

import (
	"log"
	"magical-crwler/constants"
	"time"

	"gopkg.in/telebot.v4"
)

type Bot struct {
	Bot    *telebot.Bot
	Config BotConfig
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
		Bot:    bot,
		Config: config,
	}, nil
}

func (b *Bot) RegisterHandlers() {
	b.Bot.Handle("/start", StartHandler(b))
	b.Bot.Handle("/menu", MainMenuHandler)
	b.Bot.Handle(constants.SearchButton, FilterHandlers(b))
	b.Bot.Handle("/exportFile", ExportHandler(b))
}

func (b *Bot) StartBot() {
	log.Print("Bot is running !")
	b.RegisterHandlers()
	b.Bot.Start()
	b.Bot.Start()
}
