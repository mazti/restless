package model

import (
	"time"
)

type Sample struct {
	ID        int64 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `gorm:"type:varchar(255)"`
	Provider  string `gorm:"type:varchar(255)"`
}
