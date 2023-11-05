package domains_orders_aggregates

import (
	domains_orders_events "eventsourcing/internal/domains/orders/events"
	"eventsourcing/internal/eventsourcing"

	"github.com/gofrs/uuid"
)

func ReplayOrder(ID uuid.UUID) *OrderAggregate {
	return &OrderAggregate{}
}

var _ eventsourcing.Aggregate = (*OrderAggregate)(nil)

type OrderAggregate struct {
	ID         uuid.UUID `json:"-"`
	Amount     int64     `json:"Amount,omitempty"`
	PaidAmount int64     `json:"PaidAmount,omitempty"`
	Status     string    `json:"Status,omitempty"`
	DoneBy     string    `json:"DoneBy,omitempty"`
	Root       *eventsourcing.AggregateRoot
}

var orderStatusTransition = []eventsourcing.Status{
	{From: "", To: "created", EventName: "CreateOrder"},
	{From: "created", To: "paid", EventName: "PayOrder"},
	{From: "paid", To: "completed", EventName: "CompleteOrder"},
	{From: "completed", To: "failed", EventName: "FailOrder"},
	{From: "paid", To: "failed", EventName: "FailOrder"},
	{From: "created", To: "failed", EventName: "FailOrder"},
}

func (*OrderAggregate) StatusTransition() []eventsourcing.Status {
	return orderStatusTransition
}

func (a *OrderAggregate) SID() string {
	return a.ID.String()
}

func (a *OrderAggregate) GetStatus() string {
	return a.Status
}

// event
func (a *OrderAggregate) Create(amount int64) {
	event := &domains_orders_events.OrderCreatedEvent{
		Amount: amount,
	}
	a.Root.Apply(event)
}

func (a *OrderAggregate) Pay(paidAmount int64) {
	event := &domains_orders_events.OrderPaidEvent{
		PaidAmount: paidAmount,
	}
	a.Root.Apply(event)
}

func (a *OrderAggregate) Complete(doneBy string) {
	event := &domains_orders_events.OrderCompletedEvent{
		DoneBy: doneBy,
	}
	a.Root.Apply(event)
}

func (a *OrderAggregate) Fail() {
	event := &domains_orders_events.OrderFailedEvent{}
	a.Root.Apply(event)
}
