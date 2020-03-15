package dbaccess

import (
	"strconv"
	"testing"

	"github.com/MasatoTokuse/motting/motting/model"
	"github.com/MasatoTokuse/motting/motting/test"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestQueryPhrases(t *testing.T) {

	// テストデータ準備
	db := ConnectGormInTest()
	defer db.Close()
	testdata := prepareTestDataPhrase(db)

	param := &ParamPhrase{}
	actual := *QueryPhrases(db, param)
	assertPhrases(t, *testdata, actual)

	param = &ParamPhrase{
		UserID: "whitebox",
	}
	actual = *QueryPhrases(db, param)
	assertPhrases(t, (*testdata)[:5], actual)

	param = &ParamPhrase{
		UserID: "whitebox",
		Text:   "3",
	}
	actual = *QueryPhrases(db, param)
	assertPhrases(t, (*testdata)[2:3], actual)

	param = &ParamPhrase{
		UserID: "blackbox",
		Author: "0",
	}
	actual = *QueryPhrases(db, param)
	assertPhrases(t, (*testdata)[9:10], actual)

	param = &ParamPhrase{
		ID: "9",
	}
	actual = *QueryPhrases(db, param)
	assertPhrases(t, (*testdata)[8:9], actual)

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

func cleanupTestDataPhrase(db *gorm.DB) {
	db.Exec("TRUNCATE TABLE phrases;")
}

func assertPhrases(t *testing.T, expected []model.Phrase, actual []model.Phrase) {
	// 件数が一致すること
	assert.Equal(t, len(expected), len(actual), test.ErrMsgNotMatchD, len(expected), len(actual))

	// 各フィールドが一致すること
	for i, act := range actual {
		assert.Equal(t, expected[i].ID, act.ID, test.ErrMsgNotMatchD, expected[i].ID, act.ID)
		assert.Equal(t, expected[i].UserID, act.UserID, test.ErrMsgNotMatchD, expected[i].UserID, act.UserID)
		assert.Equal(t, expected[i].Text, act.Text, test.ErrMsgNotMatchD, expected[i].Text, act.Text)
		assert.Equal(t, expected[i].Author, act.Author, test.ErrMsgNotMatchD, expected[i].Author, act.Author)
	}
}
