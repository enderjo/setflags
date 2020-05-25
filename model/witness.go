package model

import (
	"github.com/gofrs/uuid"
)

//Witness witness
type Witness struct {
	FlagID  uuid.UUID `json:"flagId" gorm:"column:flagId;type:varchar(64);primary_key"`
	PayeeID uuid.UUID `json:"payeeId" gorm:"column:payeeId;type:varchar(64);not null"`
}
