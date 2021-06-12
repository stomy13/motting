package mysql

type DBError struct {
	Message string `json:"message"`
}

func NewDBErrorIfNotNil(err error) error {
	if err == nil {
		return nil
	}
	return &DBError{Message: err.Error()}
}

func (DBError *DBError) Error() string {
	return DBError.Message
}
