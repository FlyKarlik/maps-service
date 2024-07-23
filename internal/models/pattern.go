package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// Pattern struct for pattern
type Pattern struct {
	ID string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`

	Name string `gorm:"not null" json:"name"`
	Img  string `gorm:"type:text;not null" json:"img"`
	X    int32  `gorm:"default:100;not null" json:"x"`
	Y    int32  `gorm:"default:100;not null" json:"y"`

	CreateUserIP string `json:"create_user_ip"`
	UpdateUserIP string `json:"update_user_ip"`
	CreateUserID string `json:"create_user_id"`
	UpdateUserID string `json:"update_user_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

// BeforeCreate function to get uuid for new pattern
func (pattern *Pattern) BeforeCreate(tx *gorm.DB) (err error) {
	pattern.ID = uuid.NewString()

	return
}
