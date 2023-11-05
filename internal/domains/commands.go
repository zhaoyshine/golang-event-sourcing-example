package domains

// import (
// 	domains_orders "eventsourcing/internal/domains/orders"
// 	domains_orders_commands "eventsourcing/internal/domains/orders/commands"
// 	"eventsourcing/internal/eventsourcing"
// )

// type config struct {
// 	Order struct {
// 		Command        *domains_orders_commands.OrderCommand
// 		CommandService *eventsourcing.Configuration
// 	}
// }

// var Domain *config

// func InitDomains() {
// 	Domain = &config{
// 		Order: struct {
// 			Command        *domains_orders_commands.OrderCommand
// 			CommandService *eventsourcing.Configuration
// 		}{&domains_orders_commands.OrderCommand{}, domains_orders.Configuration},
// 	}
// }
