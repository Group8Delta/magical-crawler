package alerting

import (
	"fmt"
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

			if !(adminSelector >= 0 && adminSelector < len(config.AdminTelegramIds)) {
				adminSelector = 0
				continue
			}
			if n.AlertQueue.Len() < 1 {
				continue
			}

			m := n.AlertQueue.Pop().(*Alert)
			fmt.Printf("alert found:%v\n", m)
			if m != nil {
				recipientIdentifier := strconv.Itoa(config.AdminTelegramIds[adminSelector])
				err := n.notifier.Notify(recipientIdentifier, &notification.Message{Title: m.Title, Content: m.Content})
				if err != nil {
					fmt.Println(err.Error())
				}
				adminSelector++
			}

		}
	}()

}
