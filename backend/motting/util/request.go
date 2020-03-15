package util

import (
	"io"
	"io/ioutil"
	"net/url"
)

// ParseBody is parsing request body to url.values
func ParseBody(rc *io.ReadCloser) (*url.Values, error) {
	bytes, err := ioutil.ReadAll(*rc)
	if err != nil {
		return nil, err
	}
	values, err := url.ParseQuery(string(bytes))
	if err != nil {
		return nil, err
	}
	return &values, nil
}
