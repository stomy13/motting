package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/MasatoTokuse/motting/motting/dbaccess"
	"github.com/MasatoTokuse/motting/motting/model"
	"github.com/MasatoTokuse/motting/motting/test"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

const urlPushTime = "http://loclahost:3001/api/v1/pushtime/"

func setup() {
	conargs := &dbaccess.ConnectArgs{
		Address:  "localhost",
		Port:     "33333",
		DBName:   "motting",
		User:     "motting",
		Password: "motting"}
	conargs.SetDefault()
}

func cleanupTestDataPushTime(db *gorm.DB) {
	db.Exec("TRUNCATE TABLE push_times;")
}

func prepareTestDataPushTime(db *gorm.DB) *[]model.PushTime {

	cleanupTestDataPushTime(db)

	pushtimes := []model.PushTime{}
	for i := 1; i <= 2; i++ {
		userid := ""
		if i == 1 {
			userid = "whitebox"
		} else {
			userid = "blackbox"
		}

		phrase := model.PushTime{
			UserID: userid,
			PushAt: "12:00",
		}
		pushtimes = append(pushtimes, phrase)
	}

	for i := range pushtimes {
		db.Create(&pushtimes[i])
	}

	return &pushtimes
}

func TestPushTimeGET(t *testing.T) {

	// テストデータ準備
	db := dbaccess.ConnectGormInTest()
	defer db.Close()
	testdata := prepareTestDataPushTime(db)

	// テスト用のリクエスト作成
	query := "?userid=whitebox"
	testPushTimeGET(t, query, (*testdata)[0])

	// 初回アクセス時のテスト
	query = "?userid=bluebox"
	expected := model.PushTime{
		UserID: "bluebox",
		PushAt: "10:00",
	}
	expected.ID = 3
	testPushTimeGET(t, query, expected)

}

func testPushTimeGET(t *testing.T, query string, expected model.PushTime) {

	// テスト用のリクエストとレスポンスを作成
	req := httptest.NewRequest("GET", urlPushTime+query, nil)
	res := httptest.NewRecorder()

	// ハンドラーの実行
	PushTimeGET(res, req)

	// レスポンスのステータスコードのテスト
	if res.Code != http.StatusOK {
		t.Errorf(test.ErrMsgResCode, res.Code)
	}

	// レスポンスのボディのテスト
	var actual model.PushTime
	err := json.Unmarshal(res.Body.Bytes(), &actual)
	if err != nil {
		t.Error(err)
	}

	// 各フィールドが一致すること
	assert.Equal(t, expected.ID, actual.ID, test.ErrMsgNotMatchD, expected.ID, actual.ID)
	assert.Equal(t, expected.UserID, actual.UserID, test.ErrMsgNotMatchS, expected.UserID, actual.UserID)
	assert.Equal(t, expected.PushAt, actual.PushAt, test.ErrMsgNotMatchS, expected.PushAt, actual.PushAt)

}

func TestPushTimePATCH(t *testing.T) {

	// テストデータ準備
	db := dbaccess.ConnectGormInTest()
	defer db.Close()
	prepareTestDataPushTime(db)

	expected := model.PushTime{
		UserID: "blackbox",
		PushAt: "18:00",
	}

	// テスト用のリクエストとレスポンスを作成
	values := url.Values{}
	values.Add("userid", expected.UserID)
	values.Add("pushAt", expected.PushAt)
	req := httptest.NewRequest("PATCH", urlPhrase, strings.NewReader(values.Encode()))
	res := httptest.NewRecorder()

	// ハンドラーの実行
	PushTimePATCH(res, req)

	// レスポンスのステータスコードのテスト
	if res.Code != http.StatusOK {
		t.Errorf(test.ErrMsgResCode, res.Code)
	}

	// レスポンスのテスト
	resPushTime := &model.PushTime{}
	json.Unmarshal(res.Body.Bytes(), resPushTime)
	assert.Equal(t, expected.UserID, resPushTime.UserID, test.ErrMsgNotMatchS, expected.UserID, resPushTime.UserID)
	assert.Equal(t, expected.PushAt, resPushTime.PushAt, test.ErrMsgNotMatchS, expected.PushAt, resPushTime.PushAt)

	// DBが更新されていることの確認
	actual := &model.PushTime{}
	db.Where("user_id = ?", expected.UserID).Find(actual)
	assert.Equal(t, expected.UserID, actual.UserID, test.ErrMsgNotMatchS, expected.UserID, actual.UserID)
	assert.Equal(t, expected.PushAt, actual.PushAt, test.ErrMsgNotMatchS, expected.PushAt, actual.PushAt)

}
