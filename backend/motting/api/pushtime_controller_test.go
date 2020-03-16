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

const urlPushTime = "http://loclahost:3000/pushtime/"

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
	values := url.Values{}
	values.Set("userid", "whitebox")
	testPushTimeGET(t, &values, (*testdata)[0])

	// 初回アクセス時のテスト
	values = url.Values{}
	values.Set("userid", "bluebox")
	expected := model.PushTime{
		UserID: "bluebox",
		PushAt: "10:00",
	}
	expected.ID = 3
	testPushTimeGET(t, &values, expected)

}

func testPushTimeGET(t *testing.T, values *url.Values, expected model.PushTime) {

	// テスト用のリクエストとレスポンスを作成
	req := httptest.NewRequest("GET", urlPushTime, strings.NewReader(values.Encode()))
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

	expected := model.PushTime{
		UserID: "blackbox",
		PushAt: "18:00",
	}

	db := dbaccess.ConnectGormInTest()
	defer db.Close()

	// テスト用のリクエスト作成
	values := url.Values{}
	values.Add("userid", expected.UserID)
	values.Add("pushAt", expected.PushAt)
	req := httptest.NewRequest("PATCH", urlPhrase, strings.NewReader(values.Encode()))
	// テスト用のレスポンス作成
	res := httptest.NewRecorder()

	// ハンドラーの実行
	PushTimePATCH(res, req)

	// レスポンスのステータスコードのテスト
	if res.Code != http.StatusOK {
		t.Errorf(test.ErrMsgResCode, res.Code)
	}

	// レスポンスのテスト
	resPushTime := model.PushTime{}
	json.Unmarshal(res.Body.Bytes(), &resPushTime)
	if resPushTime.UserID != expected.UserID {
		t.Errorf(test.ErrMsgNotMatchS, expected.UserID, resPushTime.UserID)
	}
	if resPushTime.PushAt != expected.PushAt {
		t.Errorf(test.ErrMsgNotMatchS, expected.PushAt, resPushTime.PushAt)
	}

	// DBが更新されていることの確認
	actual := &model.PushTime{}
	db.Where("user_id = ?", expected.UserID).Find(actual)
	if actual.UserID != expected.UserID {
		t.Errorf(test.ErrMsgNotMatchS, expected.UserID, actual.UserID)
	}
	if actual.PushAt != expected.PushAt {
		t.Errorf(test.ErrMsgNotMatchS, expected.PushAt, actual.PushAt)
	}

}
