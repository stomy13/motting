package witticism

import (
	"testing"

	"github.com/MasatoTokuse/motting/motting/domain/user"
	"github.com/stretchr/testify/assert"
)

// インスタンスが生成できることの確認
func Test_NewWitticism_IsCreatableInstance(t *testing.T) {
	tellerName, _ := NewTellerName("tellerName")
	sentence, _ := NewSentence("sentence")
	ownerId := user.UserId("ownerId")
	witticism, err := NewWitticism(tellerName, sentence, &ownerId)
	assert.Nil(t, err)
	assert.NotNil(t, witticism)
}

// WitticismIdがUUIDで生成できること
func Test_NewWitticismId_IsUuid(t *testing.T) {
	_, err := NewWitticismId()
	assert.Nil(t, err)
}

// TellerNameが1文字以上で生成できること
func Test_NewTellerName_IsNotEmpty(t *testing.T) {
	_, err := NewTellerName("t")
	assert.Nil(t, err)
}

// TellerNameが0文字で生成できないこと
func Test_NewTellerName_IsNotCreatableInstance(t *testing.T) {
	_, err := NewTellerName("")
	assert.NotNil(t, err)
}

// Sentenceが1文字以上で生成できること
func Test_NewSentence__IsNotEmpty(t *testing.T) {
	_, err := NewSentence("t")
	assert.Nil(t, err)
}

// Sentenceが0文字で生成できないこと
func Test_NewSentence_IsNotCreatableInstance(t *testing.T) {
	_, err := NewSentence("")
	assert.NotNil(t, err)
}
