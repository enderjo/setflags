package model

import (
	"time"

	"github.com/gofrs/uuid"
)

//Flag flag
type Flag struct {
	ID            uuid.UUID  `json:"id" gorm:"column:id;type:varchar(64);not null;primary_key"`
	PayerID       uuid.UUID  `json:"payerId" gorm:"column:payerId;type:varchar(64);not null"`
	Task          string     `json:"task" gorm:"column:task;type:varchar(512);not null"`
	Days          int        `json:"days" gorm:"column:days;type:int(10)"`
	AssetID       uuid.UUID  `json:"assetId" gorm:"column:assetId;type:varchar(64)"`
	Amount        int        `json:"amount" gorm:"column:amount;type:int(10)"`
	TimesAchieved int        `json:"timesAchived" gorm:"column:timesAchived;type:int(10)"`
	CreatedAt     *time.Time `json:"createAt" gorm:"column:createAt;type:datetime;default:CURRENT_TIMESTAMP"`
	UpdatedAt     *time.Time `json:"updateAt" gorm:"column:updateAt;type:datetime"`
}
