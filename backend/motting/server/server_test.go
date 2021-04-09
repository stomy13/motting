package server

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MasatoTokuse/motting/motting/dbaccess"
)

func TestRunServer(t *testing.T) {

	conargs := &dbaccess.ConnectArgs{
		Address:  "localhost",
		Port:     "3306",
		DBName:   "motting",
		User:     "motting",
		Password: "motting"}

	conargs.SetDefault()

	// テスト用のサーバーを起動
	handler := NewHandler()
	ts := httptest.NewServer(handler)
	defer ts.Close()

	// RootHandler のテスト
	res, err := http.Get(ts.URL + "/")
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("invalid response: %v", res)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	res.Body.Close()
	if string(body) != "Hello Golang" {
		t.Errorf("invalid body: %s", body)
	}
}
