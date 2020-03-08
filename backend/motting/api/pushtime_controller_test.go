package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

const urlPushTime = "http://loclahost:3000/pushtime/"

func TestPushTimeGET(t *testing.T) {

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

	// レスポンスのボディのテスト
	// if res.Body.String() != "{\"ID\":1,\"CreatedAt\":\"2020-02-26T17:08:09Z\",\"UpdatedAt\":\"2020-02-26T17:08:09Z\",\"DeletedAt\":null,\"UserID\":\"whitebox\",\"Text\":\"諸行無常\",\"Author\":\"釈迦\"}" {
	// 	t.Errorf("invalid response: %#v", res.Body.String())
	// }

	t.Logf("%#v", res)
}
