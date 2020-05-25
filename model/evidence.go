package model

import (
	"time"

	"github.com/gofrs/uuid"
)

//Evidence evidence
type Evidence struct {
	AttachmentID uuid.UUID  `json:"attachmentId" gorm:"column:attachmentId;type:varchar(64);primary_key"`
	FlagID       uuid.UUID  `json:"flagId" gorm:"column:flagId;type:varchar(64)"`
	File         string     `json:"file" gorm:"column:file;type:varchar(256)"`
	Type         int        `json:"type" gorm:"column:type;type:varchar(8)"`
	CreatedAt    *time.Time `json:"createdAt" gorm:"column:createdAt;type:datetime"`
}
