package models

import (
	"github.com/google/uuid"
	"time"
)

type Users struct {
	ID        uuid.UUID `gorm:"type:uuid not null;primaryKey"`
	Role      string    `gorm:"type:varchar(255);not null"`
	Name      string    `gorm:"type:varchar(255);not null"`
	Username  string    `gorm:"type:varchar(30);not null"`
	Password  string    `gorm:"type:varchar(150);not null"`
	Phone     string    `gorm:"type:varchar(20);not null"`
	Email     string    `gorm:"type:varchar(50);not null"`
	LastVisit time.Time `gorm:"type:timestamp;default:null"`
	CreatedAt time.Time `gorm:"type:timestamp;default:null"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:null"`
	DeletedAt time.Time `gorm:"type:timestamp;default:null"`
}
