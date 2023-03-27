package models

import (
	"github.com/google/uuid"
	"time"
)

type Accounts struct {
	ID          uuid.UUID `gorm:"type:uuid not null;primaryKey"`
	Name        string    `gorm:"type:varchar(60) not null"`
	LastName    string    `gorm:"type:varchar(60) not null"`
	PhoneNumber string    `gorm:"type:varchar(20) not null;index;unique"`
	Password    string    `gorm:"type:varchar(60) not null"`
	Birthday    time.Time `gorm:"type:date;default:null"`
	LastVisit   time.Time `gorm:"type:timestamp;default:null"`
	CreatedAt   time.Time `gorm:"type:timestamp;default:null"`
	UpdatedAt   time.Time `gorm:"type:timestamp;default:null"`
}

type AccountOut struct {
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	LastName    string     `json:"lastName"`
	PhoneNumber string     `json:"phoneNumber"`
	Password    string     `json:"password"`
	Birthday    *time.Time `json:"birthday"`
	LastVisit   *time.Time `json:"lastVisit"`
	CreatedAt   *time.Time `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`
}

type UpdateAccountInput struct {
	ID       uuid.UUID `json:"-"`
	Name     string    `json:"name" binding:"required,min=2" example:"Aziz"`
	LastName string    `json:"lastName" binding:"required,min=2" example:"Farkhadov"`
	Password *string   `json:"password" binding:"min=5,max=60"`
}

type SignUpAccountInput struct {
	Name        string     `json:"name" binding:"required,min=2" example:"Aziz"`
	LastName    string     `json:"lastName" binding:"required,min=2" example:"Farkhadov"`
	PhoneNumber string     `json:"phoneNumber" binding:"required,min=12,max=12" example:"998903456789"`
	Password    string     `json:"password" binding:"required,min=5,max=60"`
	Birthday    *time.Time `json:"birthday" example:"2011-01-11T00:00:00Z"`
}

type SingInAccountInput struct {
	PhoneNumber string `json:"phoneNumber" binding:"required,min=12,max=12" example:"998901231313"`
	Password    string `json:"password" binding:"required,min=5,max=60"`
}

type Tokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
