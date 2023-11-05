package domains_orders_subscribers

import (
	domains_orders_events "eventsourcing/internal/domains/orders/events"
	"eventsourcing/internal/eventsourcing"
)

func OrderMessagers() []map[eventsourcing.Event]func(eventsourcing.Event, eventsourcing.Aggregate) {
	return []map[eventsourcing.Event]func(eventsourcing.Event, eventsourcing.Aggregate){
		{domains_orders_events.OrderCreatedEvent{}: om.AfterOrderCreated},
		{domains_orders_events.OrderPaidEvent{}: om.AfterOrderPaid},
		{domains_orders_events.OrderCompletedEvent{}: om.AfterOrderCompleted},
		{domains_orders_events.OrderFailedEvent{}: om.AfterOrderFailed},
	}
}

type orderMessager struct{}

var om = &orderMessager{}

func (*orderMessager) AfterOrderCreated(e eventsourcing.Event, a eventsourcing.Aggregate) {
}

func (*orderMessager) AfterOrderPaid(e eventsourcing.Event, a eventsourcing.Aggregate) {
}

func (*orderMessager) AfterOrderCompleted(e eventsourcing.Event, a eventsourcing.Aggregate) {
}

func (*orderMessager) AfterOrderFailed(e eventsourcing.Event, a eventsourcing.Aggregate) {
}
