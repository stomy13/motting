package model

import (
	"github.com/jinzhu/gorm"
)

type Phrase struct {
	gorm.Model
	UserID string `gorm:"size:24"`
	Text   string `gorm:"size:128"`
	Author string `gorm:"size:24"`
}

// TODO:Valid
