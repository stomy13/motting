package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

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
