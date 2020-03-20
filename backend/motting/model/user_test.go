package model

import (
	"testing"

	"github.com/MasatoTokuse/motting/motting/test"
	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	user := &User{
		UserID:   "whitebox",
		Password: "password",
		Email:    "masato@sample.com",
	}
	t.Log(user.UserID)
	t.Log(user.Password)
	t.Log(user.Email)
}

func TestUserValidate(t *testing.T) {
	user := &User{
		UserID:   "whitebox",
		Password: "",
	}

	err := user.Validate()
	if err == nil {
		t.Error("expected error, got nil")
	} else {
		assert.Equal(t, errMsgPasswordEmpty, err.Error())
	}

	user = &User{
		UserID:   "",
		Password: "password",
	}

	err = user.Validate()
	if err == nil {
		t.Error("expected error, got nil")
	} else {
		assert.Equal(t, errMsgUserIDEmpty, err.Error())
	}

	user = &User{
		UserID:   "whitebox",
		Password: "password",
	}

	err = user.Validate()
	if err != nil {
		t.Errorf(test.ErrMsgNotMatchS, "nil", err.Error())
	}
}
