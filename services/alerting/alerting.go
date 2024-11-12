package alerting

import (
	"magical-crwler/config"
	"magical-crwler/services/notification"
	"strconv"
)

type Alerter struct {
	config     *config.Config
	notifier   notification.Notifier
	AlertQueue IAlertQueue
}

var adminSelector = 0

func NewAlerter(config *config.Config, notifier notification.Notifier) *Alerter {
	return &Alerter{config: config, notifier: notifier, AlertQueue: NewAlertQueue()}
}

func (n *Alerter) SendAlert(m *Alert) {
	n.AlertQueue.Push(m)
}

func (n *Alerter) RunAdminNotifier() {
	go func() {
		for {
			if adminSelector >= 0 && adminSelector < len(config.AdminUserIds) {
				adminSelector = 0
				continue
			}
			if n.AlertQueue.Len() < 1 {
				continue
			}
			m := n.AlertQueue.Pop().(*Alert)
			if m != nil {
				user_id := strconv.Itoa(config.AdminUserIds[adminSelector])
				n.notifier.Notify(user_id, &notification.Message{Title: m.Title, Content: m.Content})
				adminSelector++
			}

		}
	}()

}
