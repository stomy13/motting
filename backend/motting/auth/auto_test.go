package auth

import (
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

const urlSignUp = "http://localhost:3001/api/v1/signup"

func cleanupTestDataUser(db *gorm.DB) {
	db.Exec("TRUNCATE TABLE users;")
}

func prepareTestDataUser(db *gorm.DB) *[]model.User {

	cleanupTestDataUser(db)

	users := []model.User{}
	for i := 1; i <= 2; i++ {
		email := ""
		if i == 1 {
			email = "whitebox"
		} else {
			email = "blackbox"
		}

		user := model.User{
			Email:    email + "@sample.com",
			Password: email,
		}
		users = append(users, user)
	}

	for i := range users {
		db.Create(&users[i])
	}

	return &users
}

func getCountUsers(db *gorm.DB) int {
	var count int
	db.Model(&model.User{}).Count(&count)
	return count
}

func TestSignUpHandler_Success(t *testing.T) {

	// テストデータ準備
	db := dbaccess.ConnectGormInTest()
	defer db.Close()
	db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&model.User{})
	prepareTestDataUser(db)

	expected := model.User{
		Model:    gorm.Model{ID: 3},
		Email:    "bluebox@sample.com",
		Password: "bluebox",
	}

	// 実行前テーブル件数取得
	before := getCountUsers(db)

	// テスト用のリクエストとレスポンスを作成
	values := url.Values{}
	values.Set("email", expected.Email)
	values.Add("password", expected.Password)
	req := httptest.NewRequest("POST", urlSignUp, strings.NewReader(values.Encode()))
	res := httptest.NewRecorder()

	// ハンドラーの実行
	SignUpHandler(res, req)

	// レスポンスのステータスコードのテスト
	if res.Code != http.StatusOK {
		t.Errorf(test.ErrMsgResCode, res.Code)
	}

	// レスポンスのボディのテスト
	assert.Equal(t, "ok", res.Body.String())

	// 実行後テーブル件数取得
	after := getCountUsers(db)
	diff := after - before

	// 1レコード追加されていることの確認
	if diff != 1 {
		t.Errorf(test.ErrMsgNotMatchD, 1, diff)
	}

	// DBに登録されているか
	var actual model.User
	db.Where("id = ?", expected.ID).Find(&actual)

	// 各フィールドが一致すること
	assert.Equal(t, expected.ID, actual.ID)
	assert.Equal(t, expected.Email, actual.Email)
	assert.Equal(t, expected.Password, actual.Password)

}

func TestSignUpHandler_Error(t *testing.T) {

	// テストデータ準備
	db := dbaccess.ConnectGormInTest()
	defer db.Close()
	db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&model.User{})
	prepareTestDataUser(db)

	cases := []struct {
		model.User
		ExpectedRespMsg string
	}{
		{
			User: model.User{
				Email:    "",
				Password: "bluebox",
			},
			ExpectedRespMsg: "email is empty",
		},
		{
			User: model.User{
				Email:    "bluebox@sample.com",
				Password: "",
			},
			ExpectedRespMsg: "password is empty",
		},
		{
			User: model.User{
				Email:    "",
				Password: "",
			},
			ExpectedRespMsg: "email is empty",
		},
		{
			User: model.User{
				Email:    "whitebox@sample.com",
				Password: "whitebox",
			},
			ExpectedRespMsg: "this email is already used",
		},
	}

	for _, c := range cases {

		// 実行前テーブル件数取得
		before := getCountUsers(db)

		// テスト用のリクエストとレスポンスを作成
		values := url.Values{}
		values.Set("email", c.Email)
		values.Add("password", c.Password)
		req := httptest.NewRequest("POST", urlSignUp, strings.NewReader(values.Encode()))
		res := httptest.NewRecorder()

		// ハンドラーの実行
		SignUpHandler(res, req)

		// レスポンスのステータスコードのテスト
		if res.Code != http.StatusOK {
			t.Errorf(test.ErrMsgResCode, res.Code)
		}

		// レスポンスのボディのテスト
		assert.Equal(t, c.ExpectedRespMsg, res.Body.String())

		// 実行後テーブル件数取得
		after := getCountUsers(db)
		diff := after - before

		// レコード追加されていないことの確認
		if diff != 0 {
			t.Errorf(test.ErrMsgNotMatchD, 1, diff)
		}

	}

}
