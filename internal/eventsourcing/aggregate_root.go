package eventsourcing

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

type AggregateRoot struct {
	Aggregate         Aggregate
	Version           int64
	AppliedEvents     []*AggregateEvent
	UncommittedEvents []*AggregateEvent
	Configuration     *Configuration
	Repository        Repository
}

func (a *AggregateRoot) Apply(event Event) {
	b, err := json.Marshal(event)
	if err != nil {
		fmt.Println(err)
	}
	name := reflect.TypeOf(event).Name()
	e := &AggregateEvent{
		AggregateID:   a.Aggregate.SID(),
		Type:          name,
		Version:       a.Version + 1,
		Payload:       b,
		OccurredAt:    time.Now().Local().UTC(),
		OriginalEvent: event,
	}
	a.publishEventWhenApply(e)
	a.AppliedEvents = append(a.AppliedEvents, e)
	a.publishEventWhenCommitting(e)
	a.UncommittedEvents = append(a.UncommittedEvents, e)
	a.Version = e.Version
}

func (a *AggregateRoot) Commit() {
	for _, e := range a.UncommittedEvents {
		if e.commited {
			continue
		}

		saved := e.Store(a)
		if !saved {
			a.rollbackTransaction()
			return
		}
		a.publishEventWhenCommitting(e)

		err := a.commitTransaction()
		if err != nil {
			a.rollbackTransaction()
			return
		}
		a.publishEventAfterCommited(e)
		e.commited = true
	}
	a.UncommittedEvents = make([]*AggregateEvent, 0)
}

func (a *AggregateRoot) rollbackTransaction() {
}

func (a *AggregateRoot) commitTransaction() error {
	return nil
}

func (a *AggregateRoot) publishEventWhenApply(e *AggregateEvent) {
	hs, ok := a.Configuration.eventAppliedHandlers[e.Type]
	if ok {
		e.Publish(a.Aggregate, hs)
	}
}

func (a *AggregateRoot) publishEventWhenCommitting(e *AggregateEvent) {
	hs, ok := a.Configuration.eventCommittingHandlers[e.Type]
	if ok {
		e.Publish(a.Aggregate, hs)
	}
}

func (a *AggregateRoot) publishEventAfterCommited(e *AggregateEvent) {
	hs, ok := a.Configuration.eventCommittedHandlers[e.Type]
	if ok {
		e.Publish(a.Aggregate, hs)
	}
}
