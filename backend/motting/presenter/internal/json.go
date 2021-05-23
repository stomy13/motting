package internal

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

func UnmarshalJson(body io.Reader, v interface{}) error {
	bytes, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, v)
	if err != nil {
		return NewUnmarshalJsonError(err)
	}
	return nil
}

type UnmarshalJsonError struct {
	err error
}

func NewUnmarshalJsonError(err error) error {
	return &UnmarshalJsonError{err: err}
}

func (unmarshalJsonError *UnmarshalJsonError) Error() string {
	return unmarshalJsonError.err.Error()
}
