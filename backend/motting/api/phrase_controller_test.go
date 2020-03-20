package api

import (
	"encoding/json"
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

const urlPhrase = "http://loclahost:3001/api/v1/phrase/"

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

	query := "?userid=whitebox"
	testPhraseGET(t, query, (*testdata)[0:5])

	query = "?userid=whitebox"
	query += "&text=3"
	testPhraseGET(t, query, (*testdata)[2:3])

	query = "?userid=whitebox"
	query += "&author=4"
	testPhraseGET(t, query, (*testdata)[3:4])

	query = "?userid=box"
	testPhraseGET(t, query, []model.Phrase{})

}

func testPhraseGET(t *testing.T, query string, expected []model.Phrase) {

	// テスト用のリクエストとレスポンスを作成
	req := httptest.NewRequest("GET", urlPhrase+query, nil)
	res := httptest.NewRecorder()

	// ハンドラーの実行
	PhraseGET(res, req)

	// レスポンスのステータスコードのテスト
	if res.Code != http.StatusOK {
		t.Errorf(test.ErrMsgResCode, res.Code)
	}

	// レスポンスのボディのテスト
	var actual []model.Phrase
	json.Unmarshal(res.Body.Bytes(), &actual)

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
	expected := (*prepareTestDataPhrase(db))[:5]
	phrase := model.Phrase{
		Model:  gorm.Model{ID: 11},
		UserID: "whitebox",
		Text:   "諸行無常",
		Author: "釈迦",
	}
	expected = append(expected, phrase)

	// 実行前テーブル件数取得
	before := getCount(db)

	// テスト用のリクエストとレスポンスを作成
	values := url.Values{}
	values.Set("userid", phrase.UserID)
	values.Add("text", phrase.Text)
	values.Add("author", phrase.Author)
	req := httptest.NewRequest("POST", urlPhrase, strings.NewReader(values.Encode()))
	res := httptest.NewRecorder()

	// ハンドラーの実行
	PhrasePOST(res, req)

	// レスポンスのステータスコードのテスト
	if res.Code != http.StatusOK {
		t.Errorf(test.ErrMsgResCode, res.Code)
	}

	// レスポンスのボディのテスト
	var actual []model.Phrase
	json.Unmarshal(res.Body.Bytes(), &actual)

	// 件数が一致すること
	assert.Equal(t, len(expected), len(actual), test.ErrMsgNotMatchD, len(expected), len(actual))

	// 各フィールドが一致すること
	for i, act := range actual {
		assert.Equal(t, expected[i].ID, act.ID, test.ErrMsgNotMatchD, expected[i].ID, act.ID)
		assert.Equal(t, expected[i].UserID, act.UserID, test.ErrMsgNotMatchS, expected[i].UserID, act.UserID)
		assert.Equal(t, expected[i].Text, act.Text, test.ErrMsgNotMatchS, expected[i].Text, act.Text)
		assert.Equal(t, expected[i].Author, act.Author, test.ErrMsgNotMatchS, expected[i].Author, act.Author)
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

	// テストデータ準備
	db := dbaccess.ConnectGormInTest()
	defer db.Close()
	expected := (*prepareTestDataPhrase(db))[1:5]

	// 実行前テーブル件数取得
	before := getCount(db)

	// テスト用のリクエストとレスポンスを作成
	values := url.Values{}
	values.Set("userid", "whitebox")
	values.Add("id", "1")
	req := httptest.NewRequest("DELETE", urlPhrase, strings.NewReader(values.Encode()))
	res := httptest.NewRecorder()

	// ハンドラーの実行
	PhraseDELETE(res, req)

	// レスポンスのステータスコードのテスト
	if res.Code != http.StatusOK {
		t.Errorf(test.ErrMsgResCode, res.Code)
	}

	// レスポンスのボディのテスト
	var actual []model.Phrase
	json.Unmarshal(res.Body.Bytes(), &actual)

	// 件数が一致すること
	assert.Equal(t, len(expected), len(actual), test.ErrMsgNotMatchD, len(expected), len(actual))

	// 各フィールドが一致すること
	for i, act := range actual {
		assert.Equal(t, expected[i].ID, act.ID, test.ErrMsgNotMatchD, expected[i].ID, act.ID)
		assert.Equal(t, expected[i].UserID, act.UserID, test.ErrMsgNotMatchS, expected[i].UserID, act.UserID)
		assert.Equal(t, expected[i].Text, act.Text, test.ErrMsgNotMatchS, expected[i].Text, act.Text)
		assert.Equal(t, expected[i].Author, act.Author, test.ErrMsgNotMatchS, expected[i].Author, act.Author)
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
	expected := (*prepareTestDataPhrase(db))[:5]
	expected[1].Text = text
	expected[1].Author = author

	// テスト用のリクエストとレスポンスを作成
	values := url.Values{}
	values.Set("id", id)
	values.Add("userid", userid)
	values.Add("text", text)
	values.Add("author", author)
	req := httptest.NewRequest("PATCH", urlPhrase, strings.NewReader(values.Encode()))
	res := httptest.NewRecorder()

	// ハンドラーの実行
	PhrasePATCH(res, req)

	// レスポンスのステータスコードのテスト
	if res.Code != http.StatusOK {
		t.Errorf(test.ErrMsgResCode, res.Code)
	}

	// レスポンスのボディのテスト
	var actual []model.Phrase
	json.Unmarshal(res.Body.Bytes(), &actual)

	// 件数が一致すること
	assert.Equal(t, len(expected), len(actual), test.ErrMsgNotMatchD, len(expected), len(actual))

	// 各フィールドが一致すること
	for i, act := range actual {
		assert.Equal(t, expected[i].ID, act.ID, test.ErrMsgNotMatchD, expected[i].ID, act.ID)
		assert.Equal(t, expected[i].UserID, act.UserID, test.ErrMsgNotMatchS, expected[i].UserID, act.UserID)
		assert.Equal(t, expected[i].Text, act.Text, test.ErrMsgNotMatchS, expected[i].Text, act.Text)
		assert.Equal(t, expected[i].Author, act.Author, test.ErrMsgNotMatchS, expected[i].Author, act.Author)
	}

	// 更新されていることの確認
	phrase := &model.Phrase{}
	db.Where("id = ?", id).Find(phrase)
	assert.Equal(t, userid, phrase.UserID, test.ErrMsgNotMatchS, userid, phrase.UserID)
	assert.Equal(t, text, phrase.Text, test.ErrMsgNotMatchS, text, phrase.Text)
	assert.Equal(t, author, phrase.Author, test.ErrMsgNotMatchS, author, phrase.Author)

}
