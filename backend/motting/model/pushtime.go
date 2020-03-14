package model

import (
	"github.com/jinzhu/gorm"
)

type PushTime struct {
	gorm.Model
	UserID string `gorm:"size:24"`
	PushAt string `gorm:"size:5"`
}

// TODO:Valid
