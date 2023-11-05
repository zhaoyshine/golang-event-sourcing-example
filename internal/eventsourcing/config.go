package eventsourcing

type DomainConfig struct {
	ViewTable  string
	EventTable string
}

// type Command struct {
// 	events       []Event
// 	currentEvent *Event
// 	Aggregate    interface{}
// 	DomainConfig *DomainConfig
// 	eventHandle  map[string][]EventFunc
// }

type EventFunc func(c *Command)
