package error_response

type ErrorResponse struct {
	statusCode int
	message    string
	errors     []ValidateError
}

func NewError() *ErrorResponse {
	return &ErrorResponse{}
}

func (errorResponse *ErrorResponse) Append(filedName string, err error) {
	if err != nil {
		validateError := ValidateError{filedName: filedName, message: err.Error()}
		errorResponse.errors = append(errorResponse.errors, validateError)
	}
}

func (errorResponse *ErrorResponse) hasError() bool {
	return errorResponse.errors != nil
}

func (errorResponse *ErrorResponse) Error() string {
	return errorResponse.message
}

type ValidateError struct {
	filedName string
	message   string
}

type TestError struct {
	ErrorResponse
}
