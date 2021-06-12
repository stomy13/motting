package witticism

import (
	"testing"

	"github.com/MasatoTokuse/motting/motting/infrastracture/persistence/mysql"
	"github.com/MasatoTokuse/motting/motting/test_helper"
	"github.com/stretchr/testify/assert"
)

func Test_NewWitticismUsecaseInteractor(t *testing.T) {
	db := test_helper.GetConnection()
	repository := mysql.NewWitticismRepository(db)
	interactor := NewWitticismUsecaseInteractor(repository)

	command := AddWitticismCommand{TellerName: "tellerName", Sentence: "sentence", OwnerId: "ownerId"}
	err := interactor.AddWitticism(&command)
	assert.Nil(t, err)
}
