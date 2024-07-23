package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// Group struct for group
type Group struct {
	ID string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`

	Name string `gorm:"unique" json:"name"`

	CreateUserIP string `json:"create_user_ip"`
	UpdateUserIP string `json:"update_user_ip"`
	CreateUserID string `json:"create_user_id"`
	UpdateUserID string `json:"update_user_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

// BeforeCreate function to get uuid for new group
func (group *Group) BeforeCreate(tx *gorm.DB) (err error) {
	group.ID = uuid.NewString()

	return
}

// MapRelationGroupsJoinGroup struct
type MapRelationGroupsJoinGroup struct {
	MapGroupRelation
	Group
}

// GroupRelationMapsJoinMap struct
type GroupRelationMapsJoinMap struct {
	MapGroupRelation
	Map
}

// GroupRelationLayersJoinLayer struct for join layer
type GroupRelationLayersJoinLayer struct {
	GroupLayerRelation
	Layer
}
