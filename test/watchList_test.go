package test

import (
	"magical-crwler/models"
	"magical-crwler/models/Dtos"
	"sync"
	"testing"
)

var wl *models.WatchList

func TestCreateWatchList(t *testing.T) {
	wl, err := testWatchListService.CreateWatchList(Dtos.WatchListDto{UserId: 1, FilterId: 1, UpdateCycle: 1})
	if err != nil {
		t.Fatalf("error in create watch list : %v", err)
	}
	t.Logf("created watch list: %v", wl)
}
func TestUpdateWatchList(t *testing.T) {
	err := testWatchListService.UpdateWatchList(int(wl.ID), Dtos.WatchListDto{UpdateCycle: 2})
	if err != nil {
		t.Fatalf("error in update watch list : %v", err)
	}
	t.Logf("watch list updated")
}

func TestSendNotify(t *testing.T) {
	wg := sync.WaitGroup{}
	err := testWatchListService.SendAdsByWatchList(wl, &wg)
	if err != nil {
		t.Fatalf("error in update watch list : %v", err)
	}
}
