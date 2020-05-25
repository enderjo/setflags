package model

import (
	"time"

	"github.com/gofrs/uuid"
)

//Flag flag
type Flag struct {
	ID            uuid.UUID `json:"id"gorm:"type:varchar(64);unique_index"`
	PayerID       uuid.UUID
	Task          string
	Days          int
	AssetID       string
	Amount        int
	TimesAchieved int
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
}
