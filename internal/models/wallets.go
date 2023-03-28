package models

import (
	"github.com/google/uuid"
	"time"
)

type Wallets struct {
	ID        uuid.UUID  `gorm:"type:uuid not null;primaryKey"`
	AccountID uuid.UUID  `gorm:"type:uuid not null;index"`
	Balance   float64    `gorm:"type:decimal(13,2);default:0.00"`
	CreatedAt *time.Time `gorm:"type:timestamp;default:null"`
	UpdatedAt *time.Time `gorm:"type:timestamp;default:null"`
}

type WalletOut struct {
	ID        uuid.UUID  `json:"id"`
	AccountID uuid.UUID  `json:"accountId"`
	Balance   float64    `json:"balance"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}
