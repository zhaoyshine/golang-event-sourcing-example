package eventsourcing

type Repository interface {
	SaveEvent(*AggregateEvent, Aggregate) error
}
