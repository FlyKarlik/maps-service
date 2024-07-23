package models

type Column struct {
	ID           string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	TableID      string `gorm:"type:uuid;default:uuid_generate_v4();not null" json:"table_id"`
	DomainID     string `gorm:"type:uuid" json:"domain_id"`
	Name         string `gorm:"not null" json:"name"`
	Alias        string `gorm:"not null" json:"alias"`
	DataType     string `gorm:"not null" json:"data_type"`
	Nullable     bool   `gorm:"not null" json:"nullable"`
	Length       int    `json:"length"`
	DefaultValue string `json:"default_value"`
}
