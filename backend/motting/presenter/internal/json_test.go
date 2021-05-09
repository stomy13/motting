package internal

import (
	"bytes"
	"testing"

	"github.com/MasatoTokuse/motting/motting/application/witticism"
	"github.com/stretchr/testify/assert"
)

func Test_UnmarshalJson_CanUnmarshalToStruct(t *testing.T) {
	buffer := bytes.NewBufferString(`{"tellerName":"tellerName","sentence":"sentence","ownerId":"ownerId"}`)
	var addWitticismCommand witticism.AddWitticismCommand
	err := UnmarshalJson(buffer, &addWitticismCommand)
	assert.Nil(t, err)
	assert.Equal(t, "tellerName", addWitticismCommand.TellerName)
	assert.Equal(t, "sentence", addWitticismCommand.Sentence)
	assert.Equal(t, "ownerId", addWitticismCommand.OwnerId)
}
