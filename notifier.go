package core

import "context"

// Notifier defines the contract for sending notifications.
// Implementations include EmailNotifier, SMSNotifier, etc.
type Notifier interface {
	Send(ctx context.Context, msg Message) error
}
