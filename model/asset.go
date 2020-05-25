package model

import (
	"time"

	"github.com/gofrs/uuid"
)

//Asset asset
type Asset struct {
	ID       uuid.UUID  `json:"id" gorm:"column:id;type:varchar(64);primary_key"`
	Symbol   string     `json:"symbol" gorm:"column:symbol;type:varchar(64)"`
	PriceUsd int        `json:"priceUsd" gorm:"column:priceUsd;type:int(10)"`
	Balance  int        `json:"balance" gorm:"column:balance;type:int(10)"`
	PaidAt   *time.Time `json:"paidAt" gorm:"column:paidAt;type:datetime"`
}
