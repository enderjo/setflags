package model

import (
	"github.com/gofrs/uuid"
)

//Witness witness
type Witness struct {
	FlagID  uuid.UUID `json:"id"`
	PayeeID string
}
