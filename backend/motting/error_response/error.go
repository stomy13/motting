package error_response

type ValidateErrors struct {
	// トップレベルのエラーメッセージ
	Message string `json:"message"`
	// 複数の項目のエラーを返却するために使用する
	Errors []ValidateError `json:"errors"`
}

func NewValidateErrors() *ValidateErrors {
	return &ValidateErrors{Message: "入力項目にエラーがあります。"}
}

func (validateErrors *ValidateErrors) Append(filedName string, err error) {
	if err != nil {
		validateError := ValidateError{FiledName: filedName, Message: err.Error()}
		validateErrors.Errors = append(validateErrors.Errors, validateError)
	}
}

func (validateErrors *ValidateErrors) HasError() bool {
	return validateErrors.Errors != nil
}

func (validateErrors *ValidateErrors) Error() string {
	return validateErrors.Message
}

// 1項目に対するバリデーションエラー
type ValidateError struct {
	FiledName string `json:"filedName"`
	Message   string `json:"message"`
}
