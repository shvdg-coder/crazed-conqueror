package events

// EventBus defines the contract for publishing and subscribing to events
type EventBus interface {
	Publish(event Event) error
	Subscribe(eventType string, handler EventHandler) error
}
