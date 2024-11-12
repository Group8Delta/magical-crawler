package alerting

import (
	"magical-crwler/config"
	"magical-crwler/services/notification"
	"strconv"
)

type Alerting struct {
	config            *config.Config
	notifier          notification.Notifier
	adminMessageQueue IAdminMessageQueue
}

var adminSelector = 0

func New(config *config.Config, notifier notification.Notifier) *Alerting {
	return &Alerting{config: config, notifier: notifier, adminMessageQueue: NewAdminMessageQueue()}
}

func (n *Alerting) NotifyAdmins(m *AdminMessage) {
	n.adminMessageQueue.Push(m)
}

func (n *Alerting) RunAdminNotifier() {
	go func() {
		for {
			if adminSelector >= 0 && adminSelector < len(config.AdminUserIds) {
				adminSelector = 0
				continue
			}
			if n.adminMessageQueue.Len() < 1 {
				continue
			}
			m := n.adminMessageQueue.Pop().(*AdminMessage)
			if m != nil {
				user_id := strconv.Itoa(config.AdminUserIds[adminSelector])
				n.notifier.Notify(user_id, &notification.Message{Title: m.Title, Content: m.Content})
				adminSelector++
			}

		}
	}()

}
