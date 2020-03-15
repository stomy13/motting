package model

import (
	"github.com/jinzhu/gorm"
)

type Subscription struct {
	gorm.Model
	UserID   string `gorm:"size:24"`
	Endpoint string `gorm:"size:2048"`
	P256dh   string `gorm:"size:255"`
	Auth     string `gorm:"size:255"`
}
