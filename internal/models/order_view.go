package models

import (
	"github.com/gofrs/uuid"
)

type OrderView struct {
	ID         uuid.UUID `db:"id"`
	Amount     int64     `db:"amount"`
	PaidAmount int64     `db:"paid_amount"`
	Status     string    `db:"status"`
}

// func Table() string {
// 	return domains_orders.DomainsOrdersConfig.ViewTable
// }
