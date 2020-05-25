package model

import (
	"github.com/gofrs/uuid"
)

//User user
type User struct {
	ID       uuid.UUID `json:"id"`
	FullName string
	avatar   string
	Paid     bool
	Verified bool
}
