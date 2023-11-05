package eventsourcing

// func (c *Command) HandleCurrentEvent(aggregate Aggregate, event interface{}, eventName string) {
// 	c.currentEvent.EventPayload.Set(event)
// 	c.currentEvent.EventType = eventName
// 	loadAggregate(aggregate, *c.currentEvent)
// 	c.Aggregate = aggregate
// }

// func (c *Command) Execute() error {
// 	saveEvent(c.currentEvent, c.DomainConfig.EventTable)
// 	for k, fs := range c.eventHandle {
// 		if k == c.currentEvent.EventType {
// 			for _, f := range fs {
// 				f(c)
// 			}
// 			break
// 		}
// 	}
// 	return nil
// }

// func toMap(i interface{}) map[string]interface{} {
// 	m := make(map[string]interface{})

// 	v := reflect.ValueOf(i)

// 	if v.Kind() == reflect.Map {
// 		for _, k := range v.MapKeys() {
// 			m[k.String()] = v.MapIndex(k).Interface()
// 		}
// 	}

// 	return m
// }

// func saveEvent(event *Event, table string) {
// 	dbInfo := fmt.Sprintf("INSERT INTO %s (view_id, version, event_type, event_payload) VALUES (:view_id, :version, :event_type, :event_payload)",
// 		table)
// 	_, err := db.NamedExec(dbInfo, map[string]interface{}{
// 		"view_id":       event.ViewID,
// 		"version":       event.Version,
// 		"event_type":    event.EventType,
// 		"event_payload": event.EventPayload,
// 	})
// 	if err != nil {
// 		fmt.Println("Failed to saveEvent:", err)
// 	}
// }
