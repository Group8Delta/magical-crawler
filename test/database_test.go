package test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDatabaseConnection(t *testing.T) {
	db, err := testDbService.GetDb().DB()
	if err != nil {
		t.Fatalf("Failed to Get Database Connections: %v", err)

	}
	if err := db.Ping(); err != nil {
		t.Fatalf("Failed to ping database: %v", err)
	}
	t.Log("Successfully connected to the database")
}

func TestSelectQuery(t *testing.T) {
	db, _ := testDbService.GetDb().DB()
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

func TestGetAdsByFilterId(t *testing.T) {
	ads, err := testRepo.GetAdsByFilterId(1)
	if err != nil {
		t.Fatalf("GetAdsByFilterId failed: %v", err)
	}

	jads, err := json.Marshal(ads)

	if err != nil {
		t.Fatalf("Marshal ads failed: %v", err)
	}
	fmt.Print(string(jads))

}
