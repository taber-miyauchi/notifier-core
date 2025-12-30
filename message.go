package core

// Priority indicates the urgency of a notification.
type Priority int

const (
	Low Priority = iota
	Normal
	High
)

// Message represents a notification to be sent.
type Message struct {
	Recipient string
	Subject   string
	Body      string
	Priority  Priority
}

// NewMessage creates a message with Normal priority.
func NewMessage(recipient, subject, body string) Message {
	return Message{
		Recipient: recipient,
		Subject:   subject,
		Body:      body,
		Priority:  Normal,
	}
}
