package domain

import "time"

// gorm.Modelの定義
type GormModel struct {
	ID        uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"createAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}
