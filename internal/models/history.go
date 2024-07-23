package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type History struct {
	ID          string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Date        time.Time `json:"date"`
	User        string    `json:"user"`
	Action      string    `json:"action"`
	Description string    `gorm:"type:text" json:"description"`
	TableID     string    `gorm:"type:uuid" json:"table_id"`
}

func (h *History) BeforeCreate(tx *gorm.DB) (err error) {
	h.ID = uuid.NewString()

	return
}
