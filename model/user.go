package model

import (
	"github.com/gofrs/uuid"
)

//User user
type User struct {
	ID       uuid.UUID `json:"id" gorm:"column:id;type:varchar(64);primary_key"`
	FullName string    `json:"fullName" gorm:"column:fullName;type:varchar(64);not null"`
	Avatar   string    `json:"avatar" gorm:"column:avatar;type:varchar(256)"`
	Paid     bool      `json:"paid" gorm:"column:paid;type:tinyint(1)"`
	Verified bool      `json:"verified" gorm:"column:verified;type:tinyint(1)"`
}
