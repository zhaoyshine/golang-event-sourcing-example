package eventsourcing

func (e *AggregateEvent) Publish(a Aggregate, hs []func(Event, Aggregate)) {
	for _, h := range hs {
		h(e.OriginalEvent, a)
	}
}
