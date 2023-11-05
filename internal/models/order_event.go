package models

import (
	"github.com/gofrs/uuid"
	"github.com/jackc/pgtype"
)

type OrderEvent struct {
	AggregateID  uuid.UUID    `db:"aggregate_id"`
	Version      int64        `db:"version"`
	EventType    string       `db:"event_type"`
	EventPayload pgtype.JSONB `db:"event_payload"`
}
