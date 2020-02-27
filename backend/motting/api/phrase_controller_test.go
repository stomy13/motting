package api

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/MasatoTokuse/motting/motting/dbaccess"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func setup() {
	conargs := &dbaccess.ConnectArgs{
		Address:  "localhost",
		Port:     "33306",
		DBName:   "motting",
		User:     "motting",
		Password: "motting"}
	conargs.SetDefault()
}

func getCount(db *gorm.DB) int {
	var count int
	db.Model(&Phrase{}).Count(&count)
	return count
}

func TestPhraseGET(t *testing.T) {

	setup()

	// テスト用のリクエスト作成
	req := httptest.NewRequest("GET", "http://loclahost:3000/phrase/", nil)
	// テスト用のレスポンス作成
	res := httptest.NewRecorder()

	// ハンドラーの実行
	PhraseGET(res, req)

	// レスポンスのステータスコードのテスト
	if res.Code != http.StatusOK {
		t.Errorf("invalid code: %d", res.Code)
	}

	// レスポンスのボディのテスト
	if res.Body.String() != "{\"ID\":1,\"CreatedAt\":\"2020-02-26T17:08:09Z\",\"UpdatedAt\":\"2020-02-26T17:08:09Z\",\"DeletedAt\":null,\"UserID\":\"whitebox\",\"Text\":\"諸行無常\",\"Author\":\"釈迦\"}" {
		t.Errorf("invalid response: %#v", res.Body.String())
	}

	t.Logf("%#v", res)
}

func TestPhrasePOST(t *testing.T) {

	setup()
	db := dbaccess.ConnectGorm()
	defer db.Close()

	// 実行前テーブル件数取得
	before := getCount(db)

	// テスト用のリクエスト作成
	values := url.Values{}
	values.Set("userid", "whitebox")
	values.Add("text", "諸行無常")
	values.Add("author", "釈迦")
	req := httptest.NewRequest("POST", "http://loclahost:3000/phrase/", strings.NewReader(values.Encode()))
	// Content-Type 設定
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// テスト用のレスポンス作成
	res := httptest.NewRecorder()

	// ハンドラーの実行
	PhrasePOST(res, req)

	// レスポンスのステータスコードのテスト
	if res.Code != http.StatusOK {
		t.Errorf("invalid code: %d", res.Code)
	}

	// レスポンスのボディのテスト
	if res.Body.String() != "ok" {
		t.Errorf("invalid response: %#v", res)
	}

	// 実行後テーブル件数取得
	after := getCount(db)
	diff := after - before

	// 1レコード追加されていることの確認
	if diff != 1 {
		t.Errorf("expected %d, got %d", 1, diff)
	}

	t.Logf("%#v", res)
}
