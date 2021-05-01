package witticism

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// インスタンスが生成できることの確認
func Test_NewWitticism_IsCreatableInstance(t *testing.T) {
	witticism := NewWitticism("tellerName", "sentence", "owner")
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
