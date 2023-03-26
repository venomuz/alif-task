package models

import "time"

type Accounts struct {
	ID          uint32    `gorm:"type:serial not null;primaryKey"`
	Name        string    `gorm:"type:varchar(60) not null"`
	LastName    string    `gorm:"type:varchar(60) not null"`
	PhoneNumber string    `gorm:"type:varchar(20) not null;index"`
	Password    string    `gorm:"type:varchar(60) not null"`
	Birthday    time.Time `gorm:"type:date;default:null"`
	LastVisit   time.Time `gorm:"type:timestamp;default:null"`
	CreatedAt   time.Time `gorm:"type:timestamp;default:null"`
	UpdatedAt   time.Time `gorm:"type:timestamp;default:null"`
}

type AccountOut struct {
	ID          uint32     `json:"id"`
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
	Name     string  `json:"name" binding:"required,min=2" example:"Aziz"`
	LastName string  `json:"lastName" binding:"required,min=2" example:"Farkhadov"`
	Password *string `json:"password" binding:"min=5,max=60"`
}

type SignUpAccountInput struct {
	Name        string     `json:"name" binding:"required,min=2" example:"Aziz"`
	LastName    string     `json:"lastName" binding:"required,min=2" example:"Farkhadov"`
	PhoneNumber string     `json:"phoneNumber" binding:"required,min=9,max=9" example:"998903456789"`
	Password    string     `json:"password" binding:"required,min=5,max=60"`
	Birthday    *time.Time `json:"birthday" time_format:"2006-01-02" example:"2011-01-11"`
}

type SingInAccountInput struct {
	PhoneNumber string `json:"phoneNumber" binding:"required,min=9,max=9" example:"998903456789"`
	Password    string `json:"password" binding:"required,min=5,max=60"`
}

type Tokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
