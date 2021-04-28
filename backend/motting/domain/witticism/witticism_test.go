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
