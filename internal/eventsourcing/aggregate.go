package eventsourcing

type Aggregate interface {
	SID() string
	GetStatus() string
	StatusTransition() []Status
}
