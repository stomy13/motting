package mysql

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewDBErrorIfNotNil_ErrorIsNil(t *testing.T) {
	dbError := NewDBErrorIfNotNil(nil)
	assert.Nil(t, dbError)
}

func Test_NewDBErrorIfNotNil_ErrorIsNotNil(t *testing.T) {
	dbError := NewDBErrorIfNotNil(errors.New("TestDBError"))
	assert.NotNil(t, dbError)
	_, ok := dbError.(*DBError)
	assert.True(t, ok)
}
