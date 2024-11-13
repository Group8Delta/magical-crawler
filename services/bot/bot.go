package bot

import (
	"log"
	"magical-crwler/config"
	"magical-crwler/constants"
	"time"

	"gopkg.in/telebot.v4"
	"gorm.io/gorm"
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

func (b *Bot) RegisterHandlers(db *gorm.DB) {
	b.Bot.Handle("/menu", MainMenuHandler)
	b.Bot.Handle("/start", StartHandler(b))
	b.Bot.Handle(&telebot.Btn{Unique: "export"}, ExportHandler(b))
	b.Bot.Handle(constants.FiltersButton, FilterHandlers(b))
	b.Bot.Handle("/exportFile", ExportHandler(b))
	b.Bot.Handle(&telebot.Btn{Unique: "export"}, ExportHandler(b))
	b.Bot.Handle(&telebot.Btn{Unique: "export_csv"}, export_csv_Handler(b))
	b.Bot.Handle(&telebot.Btn{Unique: "export_xlsx"}, export_xlsx_Handler(b))
	b.bot.Handle(config.AdminPanelButton, AdminHandler(b))
	b.bot.Handle(config.AddAdminButton, AddAdminHandler(b, db))
	b.bot.Handle(config.RemoveAdminButton, RemoveAdminHandler(b, db))
}

func (b *Bot) StartBot(db *gorm.DB) {
	log.Print("Bot is running !")
	b.RegisterHandlers(db)
	b.Bot.Start()
	b.Bot.Start()
}
