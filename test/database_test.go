package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"magical-crwler/database"
	"magical-crwler/models"
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

func TestGetAllFilters(t *testing.T) {
	var repository = database.NewRepository(testDbService)
	filters, err := repository.GetAllFilters()
	fmt.Printf("%+v\n", filters)
	assert.NoError(t, err)
}

func TestGetExistingFiltersAds(t *testing.T) {
	var repository = database.NewRepository(testDbService)
	// Create or retrieve a sample filter with gorm.Model's ID
	filter := models.Filter{Model: gorm.Model{ID: 1}}
	adIDs, err := repository.GetExistingFiltersAds(filter)
	fmt.Printf("%+v\n", adIDs)
	assert.NoError(t, err)
}

func TestGetAdsByIDs(t *testing.T) {
	var repository = database.NewRepository(testDbService)
	// Create or retrieve a sample filter with gorm.Model's ID
	ids := []int{1}
	adIDs, err := repository.GetAdsByIDs(ids)
	fmt.Printf("%+v\n", adIDs)
	assert.NoError(t, err)
}

func TestGetAFilterOwner(t *testing.T) {
	var repository = database.NewRepository(testDbService)
	// Create or retrieve a sample filter with gorm.Model's ID
	filter := models.Filter{Model: gorm.Model{ID: 1}}
	user, err := repository.GetAFilterOwner(filter)
	fmt.Printf("%+v\n", user)
	assert.NoError(t, err)
}

func TestSearchAdIDs(t *testing.T) {
	var repository = database.NewRepository(testDbService)
	// Create a sample filter with relevant fields populated
	filter := models.Filter{
		PriceRange: &models.Range{Min: 1000, Max: 5000},
	}
	adIDs, err := repository.SearchAdIDs(filter)
	fmt.Printf("%+v\n", adIDs)
	assert.NoError(t, err)
}
