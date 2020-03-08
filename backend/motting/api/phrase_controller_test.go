package api

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/MasatoTokuse/motting/motting/dbaccess"
	"github.com/MasatoTokuse/motting/motting/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const errMsgResCode = "invalid code: %d"
const urlPhrase = "http://loclahost:3000/phrase/"

func setup() {
	conargs := &dbaccess.ConnectArgs{
		Address:  "localhost",
		Port:     "33333",
		DBName:   "motting",
		User:     "motting",
		Password: "motting"}
	conargs.SetDefault()
}

func getCount(db *gorm.DB) int {
	var count int
	db.Model(&model.Phrase{}).Count(&count)
	return count
}

func TestPhraseGET(t *testing.T) {

	setup()

	// テスト用のリクエスト作成
	req := httptest.NewRequest("GET", urlPhrase, nil)
	// テスト用のレスポンス作成
	res := httptest.NewRecorder()

	// ハンドラーの実行
	PhraseGET(res, req)

	// レスポンスのステータスコードのテスト
	if res.Code != http.StatusOK {
		t.Errorf(errMsgResCode, res.Code)
	}

	// レスポンスのボディのテスト
	// if res.Body.String() != "{\"ID\":1,\"CreatedAt\":\"2020-02-26T17:08:09Z\",\"UpdatedAt\":\"2020-02-26T17:08:09Z\",\"DeletedAt\":null,\"UserID\":\"whitebox\",\"Text\":\"諸行無常\",\"Author\":\"釈迦\"}" {
	// 	t.Errorf("invalid response: %#v", res.Body.String())
	// }

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
	req := httptest.NewRequest("POST", urlPhrase, strings.NewReader(values.Encode()))
	// Content-Type 設定
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// テスト用のレスポンス作成
	res := httptest.NewRecorder()

	// ハンドラーの実行
	PhrasePOST(res, req)

	// レスポンスのステータスコードのテスト
	if res.Code != http.StatusOK {
		t.Errorf(errMsgResCode, res.Code)
	}

	// レスポンスのボディのテスト
	if res.Body.String() != "ok" {
		t.Errorf("invalid response: %#v", res.Body.String())
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

func TestPhraseDELETE(t *testing.T) {

	const id = "1"

	setup()
	db := dbaccess.ConnectGorm()
	defer db.Close()

	// 実行前テーブル件数取得
	before := getCount(db)

	// テスト用のリクエスト作成
	values := url.Values{}
	values.Set("userid", "whitebox")
	values.Add("id", id)
	req := httptest.NewRequest("DELETE", urlPhrase, strings.NewReader(values.Encode()))
	// テスト用のレスポンス作成
	res := httptest.NewRecorder()

	// ハンドラーの実行
	PhraseDELETE(res, req)

	// レスポンスのステータスコードのテスト
	if res.Code != http.StatusOK {
		t.Errorf(errMsgResCode, res.Code)
	}

	// レスポンスのボディのテスト
	if res.Body.String() != "ok" {
		t.Errorf("invalid response: %#v", res.Body.String())
	}

	// 実行前テーブル件数取得
	after := getCount(db)
	diff := after - before

	// 削除されていることの確認
	if diff != -1 {
		t.Errorf("expected %d, got %d", -1, diff)
	}

	// 元に戻す
	phrase := &model.Phrase{}
	db.Unscoped().Model(&model.Phrase{}).Where("id = ?", id).First(phrase)
	phrase.DeletedAt = nil
	db.Unscoped().Save(phrase)

}

func TestPhrasePATCH(t *testing.T) {

	const id = "2"
	const userid = "whitebox"
	const text = "諸行無常2"
	const author = "釈迦2"

	setup()
	db := dbaccess.ConnectGorm()
	defer db.Close()

	// テスト用のリクエスト作成
	values := url.Values{}
	values.Set("id", id)
	values.Add("userid", userid)
	values.Add("text", text)
	values.Add("author", author)
	req := httptest.NewRequest("PATCH", urlPhrase, strings.NewReader(values.Encode()))
	// テスト用のレスポンス作成
	res := httptest.NewRecorder()

	// ハンドラーの実行
	PhrasePATCH(res, req)

	// レスポンスのステータスコードのテスト
	if res.Code != http.StatusOK {
		t.Errorf(errMsgResCode, res.Code)
	}

	// レスポンスのボディのテスト
	if res.Body.String() != "ok" {
		t.Errorf("invalid response: %#v", res.Body.String())
	}

	// 更新されていることの確認
	phrase := &model.Phrase{}
	db.Where("id = ?", id).Find(phrase)
	if phrase.UserID != userid {
		t.Errorf("expected %s, got %s", userid, phrase.UserID)
	}
	if phrase.Text != text {
		t.Errorf("expected %s, got %s", text, phrase.Text)
	}
	if phrase.Author != author {
		t.Errorf("expected %s, got %s", author, phrase.Author)
	}

}
