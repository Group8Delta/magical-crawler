package watchList

import (
	"fmt"
	"magical-crwler/database"
	"magical-crwler/models"
	"magical-crwler/models/Dtos"
	"magical-crwler/services/notification"
	"strconv"
	"sync"
	"time"
)

type WatchList struct {
	repo     database.IRepository
	notifier notification.Notifier
}

const workersCount = 10

func New(repo database.IRepository, notifier notification.Notifier) *WatchList {
	return &WatchList{repo: repo, notifier: notifier}
}
func (w *WatchList) RunWatcher() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()
	wg := sync.WaitGroup{}
	for {
		select {
		case <-ticker.C:
			wls, err := w.repo.GetCurrentMinuteWatchLists()
			if err != nil {
				fmt.Println(err)
				return
			}

			step := 0
			for {
				if step >= len(wls) {
					break
				}

				nextStep := step + workersCount
				if nextStep > len(wls) {
					nextStep = len(wls)
				}

				for i := step; i < nextStep; i++ {
					wg.Add(1)
					go w.SendAdsByWatchList(&wls[i], &wg)
				}
				wg.Wait()

				step = nextStep
			}

		}
	}
}

func (w *WatchList) SendAdsByWatchList(wl *models.WatchList, wg *sync.WaitGroup) error {
	defer wg.Done()
	err := w.UpdateWatchList(int(wl.ID), Dtos.WatchListDto{})
	if err != nil {
		fmt.Println(err)
		return err
	}

	ads, err := w.repo.GetAdsByFilterId(int(wl.FilterID))
	if err != nil {
		fmt.Println(err)
		return err
	}
	c := ""
	for _, v := range ads {
		c += v.Link + "\n"
	}

	user, err := w.repo.GetUserById(int(wl.UserID))
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("sending watch list ads:\n%v\n to user\n", ads)
	if w.notifier == nil {
		return nil
	}
	return w.notifier.Notify(strconv.Itoa(int(user.TelegramID)), &notification.Message{Title: "your watch list ads:", Content: c})
}

func (w *WatchList) DeleteWatchList(id int) error {
	return w.repo.DeleteWatchList(id)
}

func (w *WatchList) CreateWatchList(wl Dtos.WatchListDto) (*models.WatchList, error) {
	return w.repo.CreateWatchList(wl)
}

func (w *WatchList) UpdateWatchList(id int, wl Dtos.WatchListDto) error {
	return w.repo.UpdateWatchList(id, wl)
}
