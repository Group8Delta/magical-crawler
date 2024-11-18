package notification

type Message struct {
	Title   string
	Content string
	Photo string
}

type Notifier interface {
	Notify(recipientIdentifier string, m *Message) error
}
