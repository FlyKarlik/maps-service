package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// Layer struct for layer
type Layer struct {
	ID string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`

	Name         string `json:"name"`
	LayerType    string `json:"layer_type"`
	TableID      string `json:"table_id"`
	CreateUserIP string `json:"create_user_ip"`
	UpdateUserIP string `json:"update_user_ip"`
	CreateUserID string `json:"create_user_id"`
	UpdateUserID string `json:"update_user_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

// BeforeCreate function to get uuid for new layer
func (layer *Layer) BeforeCreate(tx *gorm.DB) (err error) {
	layer.ID = uuid.NewString()

	return
}

// LayerGroupRelation struct
type GroupLayerRelation struct {
	ID         string `gorm:"primaryKey;type:uuid" json:"id"`
	GroupID    string `gorm:"type:uuid;not null" json:"group_id"`
	LayerID    string `gorm:"type:uuid;not null" json:"layer_id"`
	LayerOrder int32  `gorm:"not null" json:"layer_order"`
}

// LayerGroupRelationJoinsGroup join for relation and group
type LayerGroupRelationJoinsGroup struct {
	Group
	GroupLayerRelation
}

// BeforeCreate function to get new uuid for lg
func (lg *GroupLayerRelation) BeforeCreate(tx *gorm.DB) (err error) {
	lg.ID = uuid.NewString()

	return
}

// LayerStyleRelation struct
type LayerStyleRelation struct {
	ID      string `gorm:"primaryKey;type:uuid" json:"id"`
	StyleID string `gorm:"type:uuid;index:lsr_idx;unique;not null" json:"style_id"`
	LayerID string `gorm:"type:uuid;index:lsr_idx;unique;not null" json:"layer_id"`
}

// BeforeCreate function to get new uuid for ls
func (ls *LayerStyleRelation) BeforeCreate(tx *gorm.DB) (err error) {
	ls.ID = uuid.NewString()

	return
}
