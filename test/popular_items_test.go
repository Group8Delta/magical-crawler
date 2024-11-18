package test

import (
	"magical-crwler/database"
	"magical-crwler/models"
	"magical-crwler/models/Dtos"
	"testing"

	"github.com/stretchr/testify/assert"
)

func seedTestAds() {
	baseAds := []models.Ad{
		{Link: "link1", VisitCount: 100},
		{Link: "link2", VisitCount: 50},
		{Link: "link3", VisitCount: 200},
	}

	for _, baseAd := range baseAds {
		testRepo.CreateAd(Dtos.AdDto{
			Link:       baseAd.Link,
			VisitCount: baseAd.VisitCount,
		})
	}
}

func seedTestFilters() {
	cities := []string{"Tehran", "Esfahan", "Mashhad"}

	filters := []models.Filter{
		{City: &cities[0], SearchedCount: 50},
		{City: &cities[1], SearchedCount: 75},
		{City: &cities[2], SearchedCount: 20},
	}
	for _, filter := range filters {
		testRepo.CreateFilter(Dtos.FilterDto{
			City:          filter.City,
			SearchedCount: filter.SearchedCount,
		})
	}
}

func TestGetMostVisitedAds(t *testing.T) {
	repo := database.NewRepository(database.New())
	defer testDbService.Close()

	seedTestAds()

	resAds, err := repo.GetMostVisitedAds(2)

	assert.NoError(t, err)
	assert.Len(t, resAds, 2)
	assert.Equal(t, "link3", resAds[0].Link)
	assert.Equal(t, "link1", resAds[1].Link)
}

func TestGetMostSearchedFilters(t *testing.T) {
	defer testDbService.Close()

	seedTestFilters()

	filters, err := testRepo.GetMostSearchedFilters(2)

	assert.NoError(t, err)
	assert.Len(t, filters, 2)
	assert.Equal(t, "Esfahan", *filters[0].City)
	assert.Equal(t, "Tehran", *filters[1].City)
}

func TestGetMostSearchedSingleFilters(t *testing.T) {
	defer testDbService.Close()

	testDbService.GetDb().Exec(`
		INSERT INTO filters (price_min, price_max, city, searched_count)
		VALUES 
			(100, 200, NULL, 5000), 
			(200, 300, NULL, 6500), 
			(100, 200, NULL, 1000),
			(NULL, NULL, 'Tehran', 1000)
	`)
	results, err := testRepo.GetMostSearchedSingleFilters(3)

	assert.NoError(t, err)
	assert.Len(t, results, 3)
	assert.Equal(t, "200-300", results[0].Value)
	assert.Equal(t, "100-200", results[1].Value)
	assert.Equal(t, "Tehran", results[2].Value)
}
