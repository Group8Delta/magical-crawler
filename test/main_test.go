package test

import (
	// "magical-crwler/config"
	"magical-crwler/database"
	"os"
	"testing"
)

var testDbService database.DbService

func TestMain(m *testing.M) {
	testDbService = database.New()
	// config := config.GetConfig()
	// err := testDbService.Init(config)
	// if err != nil {
	// 	panic("Failed to connect to database: " + err.Error())
	// }

	// defer testDbService.Close()
	code := m.Run()

	os.Exit(code)

}
