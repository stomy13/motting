package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/MasatoTokuse/motting/motting/dbaccess"
	"github.com/MasatoTokuse/motting/motting/model"
	"github.com/MasatoTokuse/motting/motting/test"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func prepareTestDataUserPushTo(db *gorm.DB) *[]model.PushTime {

	cleanupTestDataPushTime(db)

	pushtimes := []model.PushTime{}

	userids := []string{"whitebox", "blackbox"}

	for _, userid := range userids {
		for i := 0; i <= 2; i++ {

			phrase := model.PushTime{
				UserID: userid,
				PushAt: "10:0" + strconv.Itoa(i),
			}
			pushtimes = append(pushtimes, phrase)
		}
	}
	phrase := model.PushTime{
		UserID: "greenbox",
		PushAt: "12:00",
	}
	pushtimes = append(pushtimes, phrase)

	for i := range pushtimes {
		db.Create(&pushtimes[i])
	}

	return &pushtimes
}

func TestUsersPushToGET(t *testing.T) {

	// テストデータ準備
	db := dbaccess.ConnectGormInTest()
	defer db.Close()
	testdata := prepareTestDataUserPushTo(db)

	// 異なるUserIDが2件
	query := "?pushAt=10:01"
	expected := (*testdata)[1:2]
	expected = append(expected, (*testdata)[4])
	testUsersPushToGET(t, query, expected)

	// UserIDが1件
	query = "?pushAt=12:00"
	testUsersPushToGET(t, query, (*testdata)[6:7])

	// UserIDが0件
	query = "?pushAt=10:03"
	testUsersPushToGET(t, query, []model.PushTime{})
	query = "?pushAt=09:59"
	testUsersPushToGET(t, query, []model.PushTime{})

}

func testUsersPushToGET(t *testing.T, query string, expected []model.PushTime) {

	// テスト用のリクエストとレスポンスを作成
	req := httptest.NewRequest("GET", urlPushTime+query, nil)
	res := httptest.NewRecorder()

	// ハンドラーの実行
	UsersPushToGET(res, req)

	// レスポンスのステータスコードのテスト
	if res.Code != http.StatusOK {
		t.Errorf(test.ErrMsgResCode, res.Code)
	}

	// レスポンスのボディのテスト
	var actual []model.PushTime
	err := json.Unmarshal(res.Body.Bytes(), &actual)
	if err != nil {
		t.Error(err)
	}

	// 各フィールドが一致すること
	for i, act := range actual {
		assert.Equal(t, expected[i].ID, act.ID, test.ErrMsgNotMatchD, expected[i].ID, act.ID)
		assert.Equal(t, expected[i].UserID, act.UserID, test.ErrMsgNotMatchS, expected[i].UserID, act.UserID)
		assert.Equal(t, expected[i].PushAt, act.PushAt, test.ErrMsgNotMatchS, expected[i].PushAt, act.PushAt)
	}
}
