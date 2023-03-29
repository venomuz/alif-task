package models

import (
	"github.com/google/uuid"
	"time"
)

type Transactions struct {
	ID             uuid.UUID `gorm:"type:uuid not null;primaryKey"`
	AccountID      uuid.UUID `gorm:"type:uuid not null;index"`
	Method         string    `gorm:"type:varchar(20) not null"`
	Reason         string    `gorm:"type:varchar(60) not null"`
	Amount         float64   `gorm:"type:decimal(13,2);default:0.00"`
	AccountBalance float64   `gorm:"type:decimal(13,2);default:0.00"`
	Sender         uuid.UUID `gorm:"type:uuid;default:null"`
	Receiver       uuid.UUID `gorm:"type:uuid;default:null"`
	CreatedAt      time.Time `gorm:"type:timestamp not null"`
}

type TransactionOut struct {
	ID             uuid.UUID  `json:"id"`
	AccountID      uuid.UUID  `json:"accountId"`
	Method         string     `json:"method"`
	Reason         string     `json:"reason"`
	Amount         float64    `json:"amount"`
	AccountBalance float64    `json:"accountBalance"`
	Sender         *uuid.UUID `json:"sender,omitempty"`
	Receiver       *uuid.UUID `json:"receiver,omitempty"`
	CreatedAt      time.Time  `json:"createdAt"`
}

type TopUpInput struct {
	AccountID      uuid.UUID `json:"-"`
	AccountPinCode uint16    `json:"-"`
	Amount         float64   `json:"amount"  binding:"required,min=500"`
	PinCode        uint16    `json:"pinCode" binding:"required,min=1000,max=9999"`
}

type TransferByPhoneNumberInput struct {
	AccountID      uuid.UUID `json:"-"`
	AccountPinCode uint16    `json:"-"`
	ReceiverPhone  string    `json:"receiverPhone" binding:"required,min=12,max=12" example:"998903456789"`
	Amount         float64   `json:"amount" binding:"required,min=500"`
	PinCode        uint16    `json:"pinCode" binding:"required,min=1000,max=9999"`
}

type WithdrawalFundsInput struct {
	AccountID      uuid.UUID `json:"-"`
	AccountPinCode uint16    `json:"-"`
	Amount         float64   `json:"amount"  binding:"required,min=500"`
	PinCode        uint16    `json:"pinCode" binding:"required,min=1000,max=9999"`
}
