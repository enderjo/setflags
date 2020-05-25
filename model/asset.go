package model

import (
	"time"

	"github.com/gofrs/uuid"
)

//Asset asset
type Asset struct {
	ID       uuid.UUID `json:"id"`
	Symbol   string
	PriceUsd string
	Balance  int
	PaidAt   *time.Time
}
