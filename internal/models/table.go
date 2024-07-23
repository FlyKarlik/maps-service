package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Table struct for table
type Table struct {
	ID                 string `gorm:"primaryKey;type:text;default:uuid_generate_v4()" json:"id"`
	Name               string `gorm:"not null;unique" json:"name"`
	Alias              string `gorm:"not null" json:"alias"`
	IsRelated          bool   `gorm:"not null;default:false" json:"is_related"`
	IsVersioned        bool   `gorm:"not null;default:false" json:"is_versioned"`
	IsArchived         bool   `gorm:"not null;default:false" json:"is_archived"`
	IsGeometryNullable bool   `gorm:"not null;default:false" json:"is_geometry_nullable"`
	GeometryType       string `gorm:"not null" json:"geometry_type"`
	SRID               int32  `gorm:"not null" json:"srid"`
	TableType          string `gorm:"not null" json:"table_type"`
}

// BeforeCreate function to get uuid
func (t *Table) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.NewString()

	return
}

type UniqueValue struct {
	Value string `db:"value"`
}

type TableFeature struct {
	Feature      string            `json:"attributes" db:"attributes"`
	Geometry     string            `json:"geometry" db:"geometry"`
	Layer        string            `json:"layer" db:"name"`
	Alias        map[string]string `json:"alias"`
	LayerFeature string            `json:"layer_feature"`
}
