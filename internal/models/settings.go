package models

import "time"

type Settings struct {
	ID        uint32     `gorm:"type:serial not null;primaryKey"`
	Title     string     `gorm:"type:varchar(255) not null"`
	Key       string     `gorm:"type:varchar(50) not null"`
	Value     string     `gorm:"type:text not null"`
	CreatedAt *time.Time `gorm:"type:timestamp;default:null"`
	UpdatedAt *time.Time `gorm:"type:timestamp;default:null"`
}

type SettingOut struct {
	ID        uint32     `json:"id"`
	Title     string     `json:"title"`
	Key       string     `json:"key"`
	Value     string     `json:"value"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

type CreateSettingInput struct {
	Title string `json:"title" binging:"required"`
	Key   string `json:"key" binging:"required"`
	Value string `json:"value" binging:"required"`
}

type UpdateSettingInput struct {
	ID    uint32 `json:"-"`
	Title string `json:"title" binging:"required"`
	Key   string `json:"key" binging:"required"`
	Value string `json:"value" binging:"required"`
}
