package mysql

import (
	"testing"

	"github.com/MasatoTokuse/motting/motting/domain/user"
	"github.com/MasatoTokuse/motting/motting/domain/witticism"
	"github.com/MasatoTokuse/motting/motting/test_helper"
	"github.com/stretchr/testify/assert"
)

func Test_NewWitticismRepository(t *testing.T) {
	db := test_helper.GetConnection()

	repository := NewWitticismRepository(db)

	ownerId := user.UserId("ownerId")
	witticism, _ := witticism.NewWitticism("tellerName", "sentence", &ownerId)

	err := repository.Save(witticism)
	assert.Nil(t, err)
}
