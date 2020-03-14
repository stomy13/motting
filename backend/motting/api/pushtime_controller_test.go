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
)

const urlPushTime = "http://loclahost:3000/pushtime/"

func TestPushTimeGET(t *testing.T) {

	expected := model.PushTime{
		UserID: "whitebox",
		PushAt: "10:00",
	}

	setup()

	// テスト用のリクエスト作成
	req := httptest.NewRequest("GET", urlPushTime, nil)
	// テスト用のレスポンス作成
	res := httptest.NewRecorder()

	// ハンドラーの実行
	PushTimeGET(res, req)

	// レスポンスのステータスコードのテスト
	if res.Code != http.StatusOK {
		t.Errorf(errMsgResCode, res.Code)
	}

	// レスポンスのテスト
	var actual model.PushTime
	json.Unmarshal(res.Body.Bytes(), &actual)
	if actual.UserID != expected.UserID {
		t.Errorf(errMsgNotMatchS, expected.UserID, actual.UserID)
	}
	if actual.PushAt != expected.PushAt {
		t.Errorf(errMsgNotMatchS, expected.PushAt, actual.PushAt)
	}

	t.Logf("%#v", res)
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
		t.Errorf(errMsgResCode, res.Code)
	}

	// レスポンスのテスト
	resPushTime := model.PushTime{}
	json.Unmarshal(res.Body.Bytes(), &resPushTime)
	if resPushTime.UserID != expected.UserID {
		t.Errorf(errMsgNotMatchS, expected.UserID, resPushTime.UserID)
	}
	if resPushTime.PushAt != expected.PushAt {
		t.Errorf(errMsgNotMatchS, expected.PushAt, resPushTime.PushAt)
	}

	// DBが更新されていることの確認
	actual := &model.PushTime{}
	db.Where("user_id = ?", expected.UserID).Find(actual)
	if actual.UserID != expected.UserID {
		t.Errorf(errMsgNotMatchS, expected.UserID, actual.UserID)
	}
	if actual.PushAt != expected.PushAt {
		t.Errorf(errMsgNotMatchS, expected.PushAt, actual.PushAt)
	}

}
