package error_response

type ValidateErrors struct {
	// トップレベルのエラーメッセージ
	message string
	// 複数の項目のエラーを返却するために使用する
	errors []ValidateError
}

func NewValidateErrors() *ValidateErrors {
	return &ValidateErrors{}
}

func (validateErrors *ValidateErrors) Append(filedName string, err error) {
	if err != nil {
		validateError := ValidateError{filedName: filedName, message: err.Error()}
		validateErrors.errors = append(validateErrors.errors, validateError)
	}
}

func (validateErrors *ValidateErrors) HasError() bool {
	return validateErrors.errors != nil
}

func (validateErrors *ValidateErrors) Error() string {
	return validateErrors.message
}

// 1項目に対するバリデーションエラー
type ValidateError struct {
	filedName string
	message   string
}
