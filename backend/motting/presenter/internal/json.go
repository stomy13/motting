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
	return err
}
