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

func (error *ValidateErrors) Append(filedName string, err error) {
	if err != nil {
		validateError := ValidateError{filedName: filedName, message: err.Error()}
		error.errors = append(error.errors, validateError)
	}
}

func (error *ValidateErrors) HasError() bool {
	return error.errors != nil
}

func (error *ValidateErrors) Error() string {
	return error.message
}

// 1項目に対するバリデーションエラー
type ValidateError struct {
	filedName string
	message   string
}
