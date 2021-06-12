package internal

import (
	"bytes"
	"testing"

	"github.com/MasatoTokuse/motting/motting/application/witticism"
	"github.com/stretchr/testify/assert"
)

// JSON文字列から構造体を返すこと
func Test_UnmarshalJson_CanUnmarshalToStruct(t *testing.T) {
	buffer := bytes.NewBufferString(`{"tellerName":"tellerName","sentence":"sentence","ownerId":"ownerId"}`)
	var addWitticismCommand witticism.AddWitticismCommand
	err := UnmarshalJson(buffer, &addWitticismCommand)
	assert.Nil(t, err)
	assert.Equal(t, "tellerName", addWitticismCommand.TellerName)
	assert.Equal(t, "sentence", addWitticismCommand.Sentence)
	assert.Equal(t, "ownerId", addWitticismCommand.OwnerId)
}

// JSON文字列でない場合はエラーを返すこと
func Test_UnmarshalJson_CannotUnmarshalToStruct(t *testing.T) {
	buffer := bytes.NewBufferString(`"tellerName":"tellerName","sentence":"sentence","ownerId":"ownerId"`)
	var addWitticismCommand witticism.AddWitticismCommand
	err := UnmarshalJson(buffer, &addWitticismCommand)
	assert.NotNil(t, err)
	_, ok := err.(*UnmarshalJsonError)
	assert.True(t, ok)
}
