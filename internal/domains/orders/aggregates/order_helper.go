package domains_orders_aggregates

// func NewOrder(id uuid.UUID) *OrderAggregate {
// 	return &OrderAggregate{
// 		ID: id,
// 	}
// }

// func OrderAggregate(command *eventsourcing.Command) *OrderAggregate {
// 	aggregate := command.Aggregate.(*OrderAggregate)

// 	return aggregate
// }

// func (*OrderAggregate) DomainConfig() *eventsourcing.DomainConfig {
// 	return &domains_orders.DomainsOrdersConfig
// }

// func (a *OrderAggregate) ViewID() uuid.UUID {
// 	return a.ID
// }

// func (a *OrderAggregate) SetValue(field string, value interface{}) {
// 	fields := reflect.ValueOf(a).Elem()
// 	fieldValue := fields.FieldByName(field)
// 	fieldType := fieldValue.Type()
// 	v := reflect.ValueOf(value)
// 	v = reflect.Indirect(v)
// 	v = v.Convert(fieldType)
// 	fieldValue.Set(v)
// }

// func (a *OrderAggregate) SetStatus(eventName string) {
// 	for _, value := range orderStatusTransition {
// 		if value.EventName == eventName {
// 			a.Status = value.To
// 		}
// 	}
// }

// func (a *OrderAggregate) EventFunc() map[string][]eventsourcing.EventFunc {
// 	return domains_orders_events.OrderEvent
// }
