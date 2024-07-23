package models

type Srids struct {
	Srid int32 `gorm:"primaryKey;type:text:integer;unique"`
}
