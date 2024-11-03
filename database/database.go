package database

import (
	"magical-crwler/config"
	"sync"

	"gorm.io/gorm"
)

var once sync.Once
var dbService DbService

type DbService interface {
	Init(cfg *config.Config) error
	Close()
	GetDb() *gorm.DB
}

func New() DbService {
	once.Do(func() {
		dbService = newPostgresDb()

	})
	return dbService
}
