package models

import (
	"time"

	"gorm.io/gorm"
)

type Cliente struct {
	gorm.Model
	ID        uint
	Name      string
	Apellido  string
	Year      uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Cliente) TableName() string {
	return "clientes"
}
