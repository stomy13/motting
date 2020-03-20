package model

import (
	"errors"

	"github.com/jinzhu/gorm"
)

const errMsgPasswordEmpty = "password is empty"
const errMsgUserIDEmpty = "userid is empty"

type User struct {
	gorm.Model
	UserID   string `gorm:"size:24"`
	Password string `gorm:"size:1024"`
	Email    string `gorm:"size:256"`
}

func (u User) Validate() error {
	if u.UserID == "" {
		return errors.New(errMsgUserIDEmpty)
	}
	if u.Password == "" {
		return errors.New(errMsgPasswordEmpty)
	}
	return nil
}
