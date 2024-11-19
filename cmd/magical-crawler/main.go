package main

import (
	"fmt"
	"log"
	"magical-crwler/config"
	"magical-crwler/database"
	"magical-crwler/services/FilterServices"
	"magical-crwler/services/alerting"
	"magical-crwler/services/bot"
	"magical-crwler/services/crawler"
	"magical-crwler/services/watchList"
	"os"
	"time"
)

func main() {
	conf := config.GetConfig()

	dbService := database.New()
	dbService.Init(conf)
	defer dbService.Close()

	db, err := dbService.GetDb().DB()
	if err != nil {
		fmt.Println("database connection error", err)
		os.Exit(1)

	}
	err = db.Ping()
	if err != nil {
		fmt.Println("database connection error", err)
		os.Exit(1)

	}
	repo := database.NewRepository(dbService)

	err = setAdminTelegramIds(repo)
	if err != nil {
		fmt.Println("set admins had error:", err)
		os.Exit(1)
	}
	// I commented on this part because it needs a VPN to run
	bot, err := bot.NewBot(repo, bot.BotConfig{
		Token:  conf.BotToken,
		Poller: 10 * time.Second,
	})
	if err != nil {
		log.Println(err.Error())
	}

	alerter := alerting.NewAlerter(conf, bot)
	alerter.RunAdminNotifier()

	initialCrawlers(conf, repo, alerter)
	filterService := FilterServices.NewFilterServices(repo, bot)
	runFilterRunner(*filterService)

	watchListService := watchList.New(repo, bot)
	go watchListService.RunWatcher()

	bot.StartBot(dbService)
	// http.ListenAndServe(":"+config.Port, nil)
}

func initialCrawlers(config *config.Config, repo database.IRepository, alerter *alerting.Alerter) {
	runIncrementalCrawl(config, repo, alerter)
	if config.EnableFullCrawl {
		timeout, _ := getCrawlerSetting(repo)
		runCrawlers(config, repo, 0, alerter, timeout)
		fmt.Println("full crawl started")
	}
}

func runCrawlers(c *config.Config, repo database.IRepository, maxDeepth int, alerter *alerting.Alerter, timeout time.Duration) {

	for _, v := range crawler.CrawlerTypes {
		crawler, err := crawler.New(v, c, repo, maxDeepth, alerter)
		if err != nil {
			panic("Failed to initial Crawler: " + err.Error())
		}
		go crawler.RunCrawler(timeout)

	}
}
func runIncrementalCrawl(c *config.Config, repo database.IRepository, alerter *alerting.Alerter) {
	go func() {
		ticker := time.NewTicker(2 * time.Hour)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				timeout, pageLimit := getCrawlerSetting(repo)
				err := setAdminTelegramIds(repo)
				if err != nil {
					fmt.Println("set admins had error:", err)
				}
				runCrawlers(c, repo, pageLimit, alerter, timeout)

			}
		}
	}()
}

func runFilterRunner(filterService FilterServices.FilterServices) {
	go func(filterService FilterServices.FilterServices) {
		ticker := time.NewTicker(2 * time.Hour)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				filterService.ApplyFilters()
			}
		}
	}(filterService)
}

func setAdminTelegramIds(repo database.IRepository) error {
	users, err := repo.GetAdminUsers()
	if err != nil {
		return err
	}
	config.AdminTelegramIds = []int{}
	for _, v := range users {
		config.AdminTelegramIds = append(config.AdminTelegramIds, int(v.TelegramID))
	}
	return nil
}

func getCrawlerSetting(repo database.IRepository) (timeout time.Duration, pageLimit int) {
	setting, err := repo.GetCrawlerSetting()

	if err != nil {
		fmt.Printf("error in get Crawler setting:%v\n", err)
		timeout = time.Second * 2000
		pageLimit = 1
	} else {
		timeout = time.Second * time.Duration(setting.CrawlTimeOutPerSearchUrlInSecond)
		pageLimit = setting.PageNumberLimit

	}
	return timeout, pageLimit
}
