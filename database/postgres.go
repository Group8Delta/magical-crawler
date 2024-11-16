package database

import (
	"fmt"
	"magical-crwler/config"
	"magical-crwler/models"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDb struct {
	dbClient *gorm.DB
}

func (p *postgresDb) Init(cfg *config.Config) error {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Tehran",
		cfg.DatabaseHost, cfg.DatabasePort, cfg.DatabaseUser, cfg.DatabasePassword, cfg.DatabaseName, cfg.DatabaseSSLMode)
	var err error

	p.dbClient, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDb, _ := p.dbClient.DB()
	err = sqlDb.Ping()
	if err != nil {
		return err
	}
	p.dbClient.AutoMigrate(&models.User{}, &models.WatchList{}, &models.Ad{}, &models.Bookmark{}, &models.CrawlerError{}, &models.CrawlerFunctionality{}, &models.PriceHistory{}, &models.Role{}, &models.SearchedWord{}, &models.FilteredAd{}, &models.Filter{})
	sqlDb.SetMaxIdleConns(cfg.DatabaseMaxIdleConns)
	sqlDb.SetMaxOpenConns(cfg.DatabaseMaxOpenConns)
	sqlDb.SetConnMaxLifetime(cfg.DatabaseConnMaxLifetime * time.Minute)

	return nil
}
func (p *postgresDb) Close() {
	conn, _ := p.dbClient.DB()
	conn.Close()
}

func (p *postgresDb) GetDb() *gorm.DB {
	return p.dbClient
}

func newPostgresDb() *postgresDb {
	return &postgresDb{}
}
