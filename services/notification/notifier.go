package notification

type Message struct {
	Title   string
	Content string
}

type Notifier interface {
	Notify(userId string, m *Message) error
}
