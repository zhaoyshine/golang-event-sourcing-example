package eventsourcing

// var db *sqlx.DB

// func InitDB(sqlxdb *sqlx.DB) {
// 	db = sqlxdb
// }

// type Aggregate interface {
// 	ViewID() uuid.UUID
// 	DomainConfig() *DomainConfig
// 	SetValue(string, interface{})
// 	SetStatus(string)
// 	EventFunc() map[string][]EventFunc
// }

// type Event struct {
// 	ViewID       uuid.UUID    `db:"view_id"`
// 	Version      int          `db:"version"`
// 	EventType    string       `db:"event_type"`
// 	EventPayload pgtype.JSONB `db:"event_payload"`
// }

// func BuildCommand(aggregate Aggregate) *Command {
// 	id := aggregate.ViewID()
// 	domainConfig := aggregate.DomainConfig()
// 	eventTable := domainConfig.EventTable
// 	events := loadEvents(id, eventTable)
// 	for _, event := range events {
// 		loadAggregate(aggregate, event)
// 	}

// 	currentEvent := &Event{
// 		ViewID:  id,
// 		Version: len(events),
// 	}

// 	config := &Command{
// 		events:       events,
// 		currentEvent: currentEvent,
// 		Aggregate:    aggregate,
// 		DomainConfig: domainConfig,
// 		eventHandle:  aggregate.EventFunc(),
// 	}
// 	fmt.Printf("%#v\n", config.Aggregate)
// 	return config
// }

// func loadAggregate(aggregate Aggregate, event Event) {
// 	data := toMap(event.EventPayload.Get())
// 	for k, v := range data {
// 		aggregate.SetValue(k, v)
// 	}
// 	aggregate.SetStatus(event.EventType)
// }

// func loadEvents(id uuid.UUID, table string) []Event {
// 	events := []Event{}
// 	dbInfo := fmt.Sprintf("SELECT * FROM %s WHERE view_id=$1 ORDER BY version ASC",
// 		table)
// 	err := db.Select(&events, dbInfo, id.String())
// 	if err != nil {
// 		fmt.Println("Failed to loadEvents:", err)
// 	}
// 	return events
// }
