package error_response

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// バリデーションエラーがない場合は、Falseを返却することを確認するテスト
func Test_ValidateErrorsHasError_ReturnFalseWhenNotHaveError(t *testing.T) {
	validateErrors := NewValidateErrors()
	assert.Equal(t, false, validateErrors.HasError())
}

// バリデーションエラーがある場合は、Trueを返却することを確認するテスト
func Test_ValidateErrorsHasError_ReturnTrueWhenHaveError(t *testing.T) {
	validateErrors := NewValidateErrors()
	validateErrors.Append("testField", errors.New("testField error"))
	assert.Equal(t, true, validateErrors.HasError())
}
