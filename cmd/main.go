package main

import (
	"fmt"

	"github.com/gofrs/uuid"

	"eventsourcing/internal/config"
	"eventsourcing/internal/db"
	"eventsourcing/internal/domains"
	"eventsourcing/internal/models"
)

func main() {
	cfg := config.LoadConfig()
	db.MustInit(cfg)
	ExecuteOrderCommand()
}

func ExecuteOrderCommand() {
	orderID := uuid.Must(uuid.NewV4())
	fmt.Printf("orderID=%#v\n", orderID.String())
	command := domains.Domain.Order.Command.CreateOrder(orderID, 100)
	domains.Domain.Order.CommandService.Execute(command)
	if err != nil {
		fmt.Println("Failed to execute CreateOrderCommand:", err)
		return
	}
	PrintOrder(orderID)

	command = domains.Order.PayOrderCommand(orderID, 100)
	err = command.Execute()
	if err != nil {
		fmt.Println("Failed to execute PayOrderCommand:", err)
		return
	}
	PrintOrder(orderID)

	command = domains.Order.FailOrderCommand(orderID)
	err = command.Execute()
	if err != nil {
		fmt.Println("Failed to execute FailOrderCommand:", err)
		return
	}
	PrintOrder(orderID)

	command = domains.Order.CompleteOrderCommand(orderID, "admin")
	err = command.Execute()
	if err != nil {
		fmt.Println("Failed to execute CompleteOrderCommand:", err)
		return
	}
	PrintOrder(orderID)

	command = domains.Order.FailOrderCommand(orderID)
	err = command.Execute()
	if err != nil {
		fmt.Println("Failed to execute FailOrderCommand:", err)
		return
	}
	PrintOrder(orderID)
}

func PrintOrder(id uuid.UUID) {
	order := models.OrderView{}
	dbInfo := fmt.Sprintf("SELECT * FROM %s WHERE id=$1",
		order.Table())
	err := db.DB.Get(&order, dbInfo, id)
	if err != nil {
		fmt.Println("Failed to get order:", err)
		return
	}
	fmt.Printf("%+v\n", order)
}
