package domains_orders_commands

import (
	"eventsourcing/internal/eventsourcing"

	"github.com/gofrs/uuid"
)

type OrderCommand struct{}

func OrderCommands() []eventsourcing.Command {
	return []eventsourcing.Command{
		CreateOrderCommand{},
		PayOrderCommand{},
		CompleteOrderCommand{},
		FailOrderCommand{},
	}
}

type CreateOrderCommand struct {
	ID     uuid.UUID
	Amount int64
}

func (*OrderCommand) CreateOrder(id uuid.UUID, amount int64) eventsourcing.Command {
	return &CreateOrderCommand{
		ID:     id,
		Amount: amount,
	}
}

type PayOrderCommand struct {
	ID         uuid.UUID
	PaidAmount int64
}

func (*OrderCommand) PayOrder(id uuid.UUID, paidAmount int64) eventsourcing.Command {
	return &PayOrderCommand{
		ID:         id,
		PaidAmount: paidAmount,
	}
}

type CompleteOrderCommand struct {
	ID     uuid.UUID
	DoneBy string
}

func (*OrderCommand) CompleteOrder(id uuid.UUID, doneBy string) eventsourcing.Command {
	return &CompleteOrderCommand{
		ID:     id,
		DoneBy: doneBy,
	}
}

type FailOrderCommand struct {
	ID uuid.UUID
}

func (*OrderCommand) FailOrder(id uuid.UUID) eventsourcing.Command {
	return &FailOrderCommand{
		ID: id,
	}
}

// func loadCommand(id uuid.UUID) *eventsourcing.Command {
// 	order := domains_orders_aggregates.NewOrder(id)
// 	return eventsourcing.BuildCommand(order)
// }

// func (*OrderCommand) CreateOrderCommand(id uuid.UUID, amount int64) *eventsourcing.Command {
// 	command := loadCommand(id)
// 	domains_orders_command_handlers.CreateOrderCommandHandler(command, amount)
// 	return command
// }

// func (*OrderCommand) PayOrderCommand(id uuid.UUID, paidAmount int64) *eventsourcing.Command {
// 	command := loadCommand(id)
// 	domains_orders_command_handlers.PayOrderCommandHandler(command, paidAmount)
// 	return command
// }

// func (*OrderCommand) CompleteOrderCommand(id uuid.UUID, doneBy string) *eventsourcing.Command {
// 	command := loadCommand(id)
// 	domains_orders_command_handlers.CompleteOrderCommandHandler(command, doneBy)
// 	return command
// }

// func (*OrderCommand) FailOrderCommand(id uuid.UUID) *eventsourcing.Command {
// 	command := loadCommand(id)
// 	domains_orders_command_handlers.FailOrderCommandHandler(command)
// 	return command
// }
