package main

import (
	"fmt"
	"magical-crwler/config"
	"magical-crwler/database"
)

func main() {
	config := config.GetConfig()

	database := database.New()
	database.Init(config)
	defer database.Close()

	db, err := database.GetDb().DB()
	if err != nil {
		fmt.Println("database connection error", err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("database connection error", err)
	}

	// bot, err := bot.NewBot(bot.BotConfig{
	// 	Token:  config.BotToken,
	// 	Poller: 10 * time.Second,
	// })
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// bot.StartBot()

}
