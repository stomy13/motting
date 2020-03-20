package model

import (
	"errors"

	"github.com/jinzhu/gorm"
)

const errMsgPasswordEmpty = "password is empty"
const errMsgEmailEmpty = "email is empty"

type User struct {
	gorm.Model
	Email    string `gorm:"size:256;unique_index;not null"`
	Password string `gorm:"size:1024;not null"`
}

func (u User) Validate() error {
	if u.Email == "" {
		return errors.New(errMsgEmailEmpty)
	}
	if u.Password == "" {
		return errors.New(errMsgPasswordEmpty)
	}
	// validate email
	return nil
}
