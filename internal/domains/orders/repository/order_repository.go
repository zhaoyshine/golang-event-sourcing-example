package domains_orders_repository

import (
	domains_orders_aggregates "eventsourcing/internal/domains/orders/aggregates"
	domains_orders_events "eventsourcing/internal/domains/orders/events"
	"eventsourcing/internal/eventsourcing"
)

func OrderRepositoryHandlers() []map[eventsourcing.Event]func(eventsourcing.Event, eventsourcing.Aggregate) {
	return []map[eventsourcing.Event]func(eventsourcing.Event, eventsourcing.Aggregate){
		{domains_orders_events.OrderCreatedEvent{}: or.CreateEventHandler},
		{domains_orders_events.OrderPaidEvent{}: or.UpdateEventHandler},
		{domains_orders_events.OrderCompletedEvent{}: or.UpdateEventHandler},
		{domains_orders_events.OrderFailedEvent{}: or.UpdateEventHandler},
	}
}

type OrderRepositoryHandler struct{}

var or = &OrderRepositoryHandler{}

func (*OrderRepositoryHandler) CreateEventHandler(e eventsourcing.Event, a eventsourcing.Aggregate) {
	aggregate := a.(*domains_orders_aggregates.OrderAggregate)
	createView(aggregate)
}

func (*OrderRepositoryHandler) UpdateEventHandler(e eventsourcing.Event, a eventsourcing.Aggregate) {
	aggregate := a.(*domains_orders_aggregates.OrderAggregate)
	updateView(aggregate)
}

func (r *OrderRepositoryHandler) SaveEvent(e eventsourcing.AggregateEvent, a eventsourcing.Aggregate) {
}

func createView(order *domains_orders_aggregates.OrderAggregate) {
}

func updateView(order *domains_orders_aggregates.OrderAggregate) {
}

// func Init() {
// 	domains_orders_events.RegisterEventFunc("CreateOrder", CreateView)
// 	domains_orders_events.RegisterEventFunc("PayOrder", UpdateView)
// 	domains_orders_events.RegisterEventFunc("CompleteOrder", UpdateView)
// 	domains_orders_events.RegisterEventFunc("FailOrder", UpdateView)
// }

// func CreateView(c *eventsourcing.Command) {
// 	aggregate := domains_orders_aggregates.OrderAggregate(c)
// 	table := c.DomainConfig.ViewTable

// 	dbInfo := fmt.Sprintf("INSERT INTO %s (id,amount,paid_amount,status) VALUES (:id,:amount,:paid_amount,:status)",
// 		table)

// 	_, err := db.DB.NamedExec(dbInfo,
// 		map[string]interface{}{
// 			"id":          aggregate.ID,
// 			"amount":      aggregate.Amount,
// 			"paid_amount": aggregate.PaidAmount,
// 			"status":      aggregate.Status,
// 		})
// 	if err != nil {
// 		fmt.Println("Failed to CreateView:", err)
// 	}
// }

// func UpdateView(c *eventsourcing.Command) {
// 	aggregate := domains_orders_aggregates.OrderAggregate(c)
// 	table := c.DomainConfig.ViewTable

// 	dbInfo := fmt.Sprintf("UPDATE %s SET amount=:amount,paid_amount=:paid_amount,status=:status WHERE id=:id",
// 		table)

// 	_, err := db.DB.NamedExec(dbInfo,
// 		map[string]interface{}{
// 			"id":          aggregate.ID,
// 			"amount":      aggregate.Amount,
// 			"paid_amount": aggregate.PaidAmount,
// 			"status":      aggregate.Status,
// 		})
// 	if err != nil {
// 		fmt.Println("Failed to CreateView:", err)
// 	}
// }

// func LoadEvents() {
// }

// func SaveEvent() {
// }
