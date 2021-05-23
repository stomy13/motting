package mysql

type DBError struct {
	err error
}

func NewDBErrorIfNotNil(err error) error {
	if err == nil {
		return nil
	}
	return &DBError{err: err}
}

func (DBError *DBError) Error() string {
	return DBError.err.Error()
}
