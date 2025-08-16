package events

import "time"

// Event represents a domain event
type Event interface {
	Type() string
	AggregateID() string
	Timestamp() time.Time
	Data() any
}

// EventHandler processes events
type EventHandler func(Event) error
