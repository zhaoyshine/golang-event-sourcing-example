package domains_orders_events

import "eventsourcing/internal/eventsourcing"

// var (
// 	OrderEvent = make(map[string][]eventsourcing.EventFunc)
// 	mu         sync.Mutex
// )

// func RegisterEventFunc(eventName string, f eventsourcing.EventFunc) {
// 	mu.Lock()
// 	defer mu.Unlock()

// 	OrderEvent[eventName] = append(OrderEvent[eventName], f)
// }

// type CreateOrderEventParam struct {
// 	Amount int64 `json:"Amount,omitempty"`
// }

// func CreateOrderEvent(amount int64) (*CreateOrderEventParam, string) {
// 	return &CreateOrderEventParam{
// 		Amount: amount,
// 	}, "CreateOrder"
// }

// type PayOrderEventParam struct {
// 	PaidAmount int64 `json:"PaidAmount,omitempty"`
// }

// func PayOrderEvent(paidAmount int64) (*PayOrderEventParam, string) {
// 	return &PayOrderEventParam{
// 		PaidAmount: paidAmount,
// 	}, "PayOrder"
// }

// type CompleteOrderEventParam struct {
// 	DoneBy string `json:"DoneBy,omitempty"`
// }

// func CompleteOrderEvent(doneBy string) (*CompleteOrderEventParam, string) {
// 	return &CompleteOrderEventParam{
// 		DoneBy: doneBy,
// 	}, "CompleteOrder"
// }

// type EventBlankParam struct{}

// func FailOrderEvent() (*EventBlankParam, string) {
// 	return &EventBlankParam{}, "FailOrder"
// }

func OrderEvents() []eventsourcing.Event {
	return []eventsourcing.Event{
		OrderCreatedEvent{},
		OrderPaidEvent{},
		OrderCompletedEvent{},
		OrderFailedEvent{},
	}
}

type OrderCreatedEvent struct {
	Amount int64 `json:"amount,omitempty"`
}
type OrderPaidEvent struct {
	PaidAmount int64 `json:"paid_amount,omitempty"`
}
type OrderCompletedEvent struct {
	DoneBy string `json:"done_by,omitempty"`
}
type OrderFailedEvent struct{}
