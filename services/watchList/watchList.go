package watchList

import (
	"fmt"
	"magical-crwler/database"
	"magical-crwler/services/notification"
	"strconv"
	"sync"
	"time"
)

var watchLists map[string]chan bool = map[string]chan bool{}
var mu sync.Mutex = sync.Mutex{}

type WatchList struct {
	repo     database.IRepository
	notifier notification.Notifier
}

func (w *WatchList) RunWatcher(userId int, filterId int, duration time.Duration, stopStatus chan bool) {
	mu.Lock()
	defer mu.Unlock()
	watchLists[fmt.Sprintf("%d_%d_%d", userId, filterId, duration.Seconds())] = stopStatus
	go w.runWatchListScheduler(userId, filterId, duration, stopStatus)
}

func (w *WatchList) StopWatcher(userId int, filterId int, duration time.Duration) {
	mu.Lock()
	defer mu.Unlock()
	stopStatus := watchLists[fmt.Sprintf("%d_%d_%d", userId, filterId, duration.Seconds())]
	stopStatus <- true
}

func (w *WatchList) runWatchListScheduler(userId int, filterId int, duration time.Duration, stopStatus chan bool) {
	ticker := time.NewTicker(duration)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			ads, err := w.repo.GetAdsByFilterId(filterId)
			if err != nil {
				fmt.Println(err)
				continue
			}
			c := ""

			for _, v := range ads {
				c += v.Link + "\n"
			}

			w.notifier.Notify(strconv.Itoa(userId), &notification.Message{Title: "your watch list ads: ", Content: c})

		case <-stopStatus:
			close(stopStatus)
			break
		}
	}
}
