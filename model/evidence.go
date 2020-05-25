package model

import (
	"time"

	"github.com/gofrs/uuid"
)

//Evidence evidence
type Evidence struct {
	AttachmentID uuid.UUID `json:"id"`
	FlagID       uuid.UUID
	File         string
	Type         int
	CreatedAt    *time.Time
}
