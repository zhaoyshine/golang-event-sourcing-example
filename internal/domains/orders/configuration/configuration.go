package domains_orders_configuration

import (
	"sync"

	domains_orders_command_handlers "eventsourcing/internal/domains/orders/command_handlers"
	domains_orders_commands "eventsourcing/internal/domains/orders/commands"
	domains_orders_events "eventsourcing/internal/domains/orders/events"
	domains_orders_repository "eventsourcing/internal/domains/orders/repository"
	domains_orders_subscribers "eventsourcing/internal/domains/orders/subscribers"
	"eventsourcing/internal/eventsourcing"
)

var DomainsOrdersConfig = eventsourcing.DomainConfig{
	ViewTable:  "order_views",
	EventTable: "order_events",
}

var (
	OrderConfiguration *eventsourcing.Configuration
	once               sync.Once
)

func MustInit() {
	once.Do(func() {
		OrderConfiguration = eventsourcing.NewConfiguration("order_views", "order_events")
		OrderConfiguration.AddCommand(domains_orders_commands.OrderCommands())
		OrderConfiguration.AddEvent(domains_orders_events.OrderEvents())

		OrderConfiguration.OnCommand(domains_orders_command_handlers.OrderCommandHandlers())
		OrderConfiguration.OnEventApplied(domains_orders_subscribers.OrderEventHandlers())
		OrderConfiguration.OnEventCommitting(domains_orders_repository.OrderRepositoryHandlers())
		OrderConfiguration.OnEventCommitted(domains_orders_subscribers.OrderMessagers())
	})
}
