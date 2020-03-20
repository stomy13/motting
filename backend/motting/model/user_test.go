package model

import (
	"testing"

	"github.com/MasatoTokuse/motting/motting/test"
	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	user := &User{
		Email:    "whitebox@sample.com",
		Password: "password",
	}
	t.Log(user.Email)
	t.Log(user.Password)
}

func TestUserValidate(t *testing.T) {
	user := &User{
		Email:    "whitebox@sample.com",
		Password: "",
	}

	err := user.Validate()
	if err == nil {
		t.Error("expected error, got nil")
	} else {
		assert.Equal(t, errMsgPasswordEmpty, err.Error())
	}

	user = &User{
		Email:    "",
		Password: "password",
	}

	err = user.Validate()
	if err == nil {
		t.Error("expected error, got nil")
	} else {
		assert.Equal(t, errMsgEmailEmpty, err.Error())
	}

	user = &User{
		Email:    "whitebox@sample.com",
		Password: "password",
	}

	err = user.Validate()
	if err != nil {
		t.Errorf(test.ErrMsgNotMatchS, "nil", err.Error())
	}
}
