package domains_orders_command_handlers

import (
	domains_orders_aggregates "eventsourcing/internal/domains/orders/aggregates"
	domains_orders_commands "eventsourcing/internal/domains/orders/commands"
	"eventsourcing/internal/eventsourcing"
)

func OrderCommandHandlers() []map[eventsourcing.Command]func(eventsourcing.Command) {
	return []map[eventsourcing.Command]func(eventsourcing.Command){
		{domains_orders_commands.CreateOrderCommand{}: ch.CreateOrderCommandHandler},
		{domains_orders_commands.PayOrderCommand{}: ch.PayOrderCommandHandler},
		{domains_orders_commands.CompleteOrderCommand{}: ch.CompleteOrderCommandHandler},
		{domains_orders_commands.FailOrderCommand{}: ch.FailOrderCommandHandler},
	}
}

type orderCommandHandler struct{}

var ch = &orderCommandHandler{}

func (*orderCommandHandler) CreateOrderCommandHandler(c eventsourcing.Command) {
	command := c.(*domains_orders_commands.CreateOrderCommand)
	order := domains_orders_aggregates.ReplayOrder(command.ID)
	order.Create(command.Amount)
	order.Root.Commit()
}

func (*orderCommandHandler) PayOrderCommandHandler(c eventsourcing.Command) {
	command := c.(*domains_orders_commands.PayOrderCommand)
	order := domains_orders_aggregates.ReplayOrder(command.ID)
	order.Pay(command.PaidAmount)
	order.Root.Commit()
}

func (*orderCommandHandler) CompleteOrderCommandHandler(c eventsourcing.Command) {
	command := c.(*domains_orders_commands.CompleteOrderCommand)
	order := domains_orders_aggregates.ReplayOrder(command.ID)
	order.Complete(command.DoneBy)
	order.Root.Commit()
}

func (*orderCommandHandler) FailOrderCommandHandler(c eventsourcing.Command) {
	command := c.(*domains_orders_commands.FailOrderCommand)
	order := domains_orders_aggregates.ReplayOrder(command.ID)
	order.Fail()
	order.Root.Commit()
}

// type OrderCommandHandler struct{}

// func (*OrderCommandHandler) InitCommandHandler() {
// 	OnCommand(CreateOrderCommand)
// }

// func CreateOrderCommandHandler(command *eventsourcing.Command, amount int64) {
// 	order := domains_orders_aggregates.OrderAggregate(command)
// 	event, eventName := order.Create(amount)
// 	command.HandleCurrentEvent(order, event, eventName)
// }

// func PayOrderCommandHandler(command *eventsourcing.Command, paidAmount int64) {
// 	order := domains_orders_aggregates.OrderAggregate(command)
// 	event, eventName := order.Pay(paidAmount)
// 	command.HandleCurrentEvent(order, event, eventName)
// }

// func CompleteOrderCommandHandler(command *eventsourcing.Command, doneBy string) {
// 	order := domains_orders_aggregates.OrderAggregate(command)
// 	event, eventName := order.Complete(doneBy)
// 	command.HandleCurrentEvent(order, event, eventName)
// }

// func FailOrderCommandHandler(command *eventsourcing.Command) {
// 	order := domains_orders_aggregates.OrderAggregate(command)
// 	event, eventName := order.Fail()
// 	command.HandleCurrentEvent(order, event, eventName)
// }
