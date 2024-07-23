package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Map struct {
	ID string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`

	Name     string `gorm:"not null" json:"name"`
	Picture  string `gorm:"type:text;not null" json:"picture"`
	Describe string `gorm:"type:text;not null" json:"describe"`
	Active   bool   `gorm:"not null;default:true" json:"active"`

	CreateUserIP string `json:"create_user_ip"`
	UpdateUserIP string `json:"update_user_ip"`
	CreateUserID string `json:"create_user_id"`
	UpdateUserID string `json:"update_user_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

// BeforeCreate function to get uuid for new layer
func (map_ *Map) BeforeCreate(tx *gorm.DB) (err error) {
	map_.ID = uuid.NewString()

	return
}

// MapGroupRelation struct
type MapGroupRelation struct {
	ID string `gorm:"primaryKey;type:uuid" json:"id"`

	GroupID    string `gorm:"type:uuid;not null" json:"group_id"`
	MapID      string `gorm:"type:uuid;not null" json:"map_id"`
	GroupOrder int32  `gorm:"not null" json:"group_order"`
}

// BeforeCreate function to get new uuid for gm
func (gm *MapGroupRelation) BeforeCreate(tx *gorm.DB) (err error) {
	gm.ID = uuid.NewString()

	return
}
