package test

import (
	"magical-crwler/database"
	"magical-crwler/services/FilterServices"
	"testing"
)

func TestApplyFilterService(t *testing.T) {
	repository := database.NewRepository(database.New())
	filterService := FilterServices.NewFilterServices(repository, nil)
	filterService.ApplyFilters()
}
