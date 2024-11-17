package test

import (
	"magical-crwler/models"
	"magical-crwler/models/Dtos"
	"sync"
	"testing"
)

var wl *models.WatchList

func TestWatchListFunctionality(t *testing.T) {
	wl, err := testWatchListService.CreateWatchList(Dtos.WatchListDto{UserId: 1, FilterId: 1, UpdateCycle: 2})
	if err != nil {
		t.Fatalf("error in create watch list : %v", err)
	}
	t.Logf("created watch list: %v", wl)

	err = testWatchListService.UpdateWatchList(int(wl.ID), Dtos.WatchListDto{UpdateCycle: 1})
	if err != nil {
		t.Fatalf("error in update watch list : %v", err)
	}
	t.Logf("watch list updated")

	wg := sync.WaitGroup{}
	wg.Add(1)
	err = testWatchListService.SendAdsByWatchList(wl, &wg)
	wg.Wait()
	if err != nil {
		t.Fatalf("error in update watch list : %v", err)
	}

	err = testWatchListService.DeleteWatchList(int(wl.ID))

	if err != nil {
		t.Fatalf("error in delete watch list : %v", err)
	}
}
