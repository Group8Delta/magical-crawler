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

func (w *WatchList) RunWatcher() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			wls, err := w.repo.GetCurrentMinuteWatchLists()
			if err != nil {
				fmt.Println(err)
				continue
			}
			for _, v := range wls {
				ads, err := w.repo.GetAdsByFilterId(int(v.FilterID))
				if err != nil {
					fmt.Println(err)
					continue
				}
				c := ""
				for _, v := range ads {
					c += v.Link + "\n"
				}
				w.notifier.Notify(strconv.Itoa(int(v.UserID)), &notification.Message{Title: "your watch list ads:", Content: c})
			}
		}
	}
}

func (w *WatchList) StopWatcher(userId int, filterId int, duration time.Duration) {

}

func (w *WatchList) runWatchListScheduler(userId int, filterId int, duration time.Duration, stopStatus chan bool) {

}
