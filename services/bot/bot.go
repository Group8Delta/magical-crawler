package bot

import (
	"fmt"
	"log"
	"magical-crwler/constants"
	"magical-crwler/services/notification"
	"strconv"
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

func (b *Bot) Notify(recipientIdentifier string, m *notification.Message) error {

	teleUserId, err := strconv.Atoi(recipientIdentifier)
	if err != nil {
		return err
	}
	recipient := telebot.ChatID(teleUserId)
	message := fmt.Sprintf("%s \n\n %s", m.Title, m.Content)
	_, err = b.Bot.Send(recipient, message)
	if err != nil {
		return err
	}

	return nil
}
func (b *Bot) RegisterHandlers(db *gorm.DB) {
	b.Bot.Handle("/menu", MainMenuHandler)
	b.Bot.Handle("/start", StartHandler(b, db))
	b.Bot.Handle(&telebot.Btn{Unique: "export"}, ExportHandler(b))
	b.Bot.Handle(constants.FiltersButton, FilterHandlers(b))
	b.Bot.Handle("/exportFile", ExportHandler(b))
	b.Bot.Handle(&telebot.Btn{Unique: "export"}, ExportHandler(b))
	b.Bot.Handle(&telebot.Btn{Unique: "export_csv"}, export_csv_Handler(b))
	b.Bot.Handle(&telebot.Btn{Unique: "export_xlsx"}, export_xlsx_Handler(b))
	b.Bot.Handle(constants.AdminPanelButton, AdminHandler(b))
	b.Bot.Handle(constants.AddAdminButton, AddAdminHandler(b, db))
	b.Bot.Handle(constants.RemoveAdminButton, RemoveAdminHandler(b, db))
}

func (b *Bot) StartBot(db *gorm.DB) {
	log.Print("Bot is running !")
	b.RegisterHandlers(db)
	b.Bot.Start()
	b.Bot.Start()
}
