package database

import (
	"magical-crwler/config"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	New()
	config := config.GetConfig()
	err := dbService.Init(config)
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	defer dbService.Close()
	code := m.Run()

	os.Exit(code)

}

func TestDatabaseConnection(t *testing.T) {

	db, err := dbService.GetDb().DB()

	if err != nil {
		t.Fatalf("Failed to Get Database Connections: %v", err)

	}
	if err := db.Ping(); err != nil {
		t.Fatalf("Failed to ping database: %v", err)
	}
	t.Log("Successfully connected to the database")
}

func TestSelectQuery(t *testing.T) {
	db, _ := dbService.GetDb().DB()
	var result int
	err := db.QueryRow("SELECT 1").Scan(&result)
	if err != nil {
		t.Fatalf("Query failed: %v", err)
	}

	if result != 1 {
		t.Fatalf("Expected 1 but got %d", result)
	}
	t.Log("Select Query executed successfully")
}
