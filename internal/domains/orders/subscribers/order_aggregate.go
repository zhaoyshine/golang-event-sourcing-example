package domains_orders_subscribers

import (
	domains_orders_aggregates "eventsourcing/internal/domains/orders/aggregates"
	domains_orders_events "eventsourcing/internal/domains/orders/events"
	"eventsourcing/internal/eventsourcing"
)

func OrderEventHandlers() []map[eventsourcing.Event]func(eventsourcing.Event, eventsourcing.Aggregate) {
	return []map[eventsourcing.Event]func(eventsourcing.Event, eventsourcing.Aggregate){
		{domains_orders_events.OrderCreatedEvent{}: oa.CreatedEventHandler},
		{domains_orders_events.OrderPaidEvent{}: oa.PaidEventHandler},
		{domains_orders_events.OrderCompletedEvent{}: oa.CompletedEventHandler},
		{domains_orders_events.OrderFailedEvent{}: oa.FailedEventHandler},
	}
}

type orderEventHandler struct{}

var oa = &orderEventHandler{}

func (*orderEventHandler) CreatedEventHandler(e eventsourcing.Event, a eventsourcing.Aggregate) {
	aggregate := a.(*domains_orders_aggregates.OrderAggregate)
	event := e.(*domains_orders_events.OrderCreatedEvent)
	aggregate.Status = "Created"
	aggregate.Amount = event.Amount
}

func (*orderEventHandler) PaidEventHandler(e eventsourcing.Event, a eventsourcing.Aggregate) {
	aggregate := a.(*domains_orders_aggregates.OrderAggregate)
	event := e.(*domains_orders_events.OrderPaidEvent)
	aggregate.Status = "Paid"
	aggregate.PaidAmount = event.PaidAmount
}

func (*orderEventHandler) CompletedEventHandler(e eventsourcing.Event, a eventsourcing.Aggregate) {
	aggregate := a.(*domains_orders_aggregates.OrderAggregate)
	event := e.(*domains_orders_events.OrderCompletedEvent)
	aggregate.Status = "Completed"
	aggregate.DoneBy = event.DoneBy
}

func (*orderEventHandler) FailedEventHandler(e eventsourcing.Event, a eventsourcing.Aggregate) {
	aggregate := a.(*domains_orders_aggregates.OrderAggregate)
	_ = e.(*domains_orders_events.OrderFailedEvent)
	aggregate.Status = "Failed"
}
