package eventsourcing

import (
	"fmt"
	"time"
)

type AggregateEvent struct {
	AggregateID   string
	Type          string
	Version       int64
	Payload       []byte
	OccurredAt    time.Time
	OriginalEvent Event
	commited      bool
}

func (e *AggregateEvent) Store(r *AggregateRoot) bool {
	err := r.Repository.SaveEvent(e, r.Aggregate)
	if err != nil {
		fmt.Println("Failed to save event:", err)
		return false
	}
	return true
}
