package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/MasatoTokuse/motting/motting/dbaccess"
	"github.com/MasatoTokuse/motting/motting/model"
	"github.com/MasatoTokuse/motting/motting/test"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/stretchr/testify/assert"
)

const urlPhrase = "http://loclahost:3000/phrase/"

func getCount(db *gorm.DB) int {
	var count int
	db.Model(&model.Phrase{}).Count(&count)
	return count
}

func cleanupTestDataPhrase(db *gorm.DB) {
	db.Exec("TRUNCATE TABLE phrases;")
}

func prepareTestDataPhrase(db *gorm.DB) *[]model.Phrase {

	cleanupTestDataPhrase(db)

	phrases := []model.Phrase{}
	for i := 1; i <= 10; i++ {
		userid := ""
		if i <= 5 {
			userid = "whitebox"
		} else {
			userid = "blackbox"
		}

		phrase := model.Phrase{
			UserID: userid,
			Text:   "text" + strconv.Itoa(i),
			Author: strconv.Itoa(i+10) + "author",
		}
		phrases = append(phrases, phrase)
	}

	for i := range phrases {
		db.Create(&phrases[i])
	}

	return &phrases
}

func TestPhraseGET(t *testing.T) {

	// テストデータ準備
	db := dbaccess.ConnectGormInTest()
	defer db.Close()
	testdata := prepareTestDataPhrase(db)

	values := url.Values{}
	values.Set("userid", "whitebox")
	testPhraseGET(t, &values, (*testdata)[0:5])

	values = url.Values{}
	values.Set("userid", "whitebox")
	values.Add("text", "3")
	testPhraseGET(t, &values, (*testdata)[2:3])

	values = url.Values{}
	values.Set("userid", "whitebox")
	values.Add("author", "4")
	testPhraseGET(t, &values, (*testdata)[3:4])

	values = url.Values{}
	values.Set("userid", "box")
	testPhraseGET(t, &values, []model.Phrase{})

}

func testPhraseGET(t *testing.T, values *url.Values, expected []model.Phrase) {

	// テスト用のリクエストとレスポンスを作成
	req := httptest.NewRequest("GET", urlPhrase, strings.NewReader(values.Encode()))
	res := httptest.NewRecorder()

	// ハンドラーの実行
	PhraseGET(res, req)

	// レスポンスのステータスコードのテスト
	if res.Code != http.StatusOK {
		t.Errorf(test.ErrMsgResCode, res.Code)
	}

	// レスポンスのボディのテスト
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	var actual []model.Phrase
	json.Unmarshal(body, &actual)

	// 件数が一致すること
	assert.Equal(t, len(expected), len(actual), test.ErrMsgNotMatchD, len(expected), len(actual))

	// 各フィールドが一致すること
	for i, act := range actual {
		assert.Equal(t, expected[i].ID, act.ID, test.ErrMsgNotMatchD, expected[i].ID, act.ID)
		assert.Equal(t, expected[i].UserID, act.UserID, test.ErrMsgNotMatchS, expected[i].UserID, act.UserID)
		assert.Equal(t, expected[i].Text, act.Text, test.ErrMsgNotMatchS, expected[i].Text, act.Text)
		assert.Equal(t, expected[i].Author, act.Author, test.ErrMsgNotMatchS, expected[i].Author, act.Author)
	}

}

func TestPhrasePOST(t *testing.T) {

	// テストデータ準備
	db := dbaccess.ConnectGormInTest()
	defer db.Close()
	prepareTestDataPhrase(db)

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
		t.Errorf(test.ErrMsgResCode, res.Code)
	}

	// レスポンスのボディのテスト
	if res.Body.String() == "" {
		t.Errorf(test.ErrMsgInvalidResBody, res.Body.String())
	}

	// 実行後テーブル件数取得
	after := getCount(db)
	diff := after - before

	// 1レコード追加されていることの確認
	if diff != 1 {
		t.Errorf(test.ErrMsgNotMatchD, 1, diff)
	}
}

func TestPhraseDELETE(t *testing.T) {

	const id = "1"

	// テストデータ準備
	db := dbaccess.ConnectGormInTest()
	defer db.Close()
	prepareTestDataPhrase(db)

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
		t.Errorf(test.ErrMsgResCode, res.Code)
	}

	// レスポンスのボディのテスト
	if res.Body.String() == "" {
		t.Errorf(test.ErrMsgInvalidResBody, res.Body.String())
	}

	// 実行後テーブル件数取得
	after := getCount(db)
	diff := after - before

	// 削除されていることの確認
	if diff != -1 {
		t.Errorf(test.ErrMsgNotMatchD, -1, diff)
	}

}

func TestPhrasePATCH(t *testing.T) {

	const id = "2"
	const userid = "whitebox"
	const text = "諸行無常2"
	const author = "釈迦2"

	// テストデータ準備
	db := dbaccess.ConnectGormInTest()
	defer db.Close()
	prepareTestDataPhrase(db)

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
		t.Errorf(test.ErrMsgResCode, res.Code)
	}

	// レスポンスのボディのテスト
	if res.Body.String() == "" {
		t.Errorf(test.ErrMsgInvalidResBody, res.Body.String())
	}

	// 更新されていることの確認
	phrase := &model.Phrase{}
	db.Where("id = ?", id).Find(phrase)
	if phrase.UserID != userid {
		t.Errorf(test.ErrMsgNotMatchS, userid, phrase.UserID)
	}
	if phrase.Text != text {
		t.Errorf(test.ErrMsgNotMatchS, text, phrase.Text)
	}
	if phrase.Author != author {
		t.Errorf(test.ErrMsgNotMatchS, author, phrase.Author)
	}

}
