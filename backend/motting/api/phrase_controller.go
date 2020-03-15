package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MasatoTokuse/motting/motting/dbaccess"
	jsonwrapper "github.com/MasatoTokuse/motting/motting/json"
	"github.com/MasatoTokuse/motting/motting/model"
	"github.com/MasatoTokuse/motting/motting/util"
	"github.com/jinzhu/gorm"
)

func PhraseGET(w http.ResponseWriter, r *http.Request) {

	// リクエストから値を受けとる
	values, err := util.ParseBody(&r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// valuesチェック
	param := &dbaccess.ParamPhrase{
		ID:     values.Get("id"),
		UserID: values.Get("userid"),
		Text:   values.Get("text"),
		Author: values.Get("author"),
	}

	db := dbaccess.ConnectGorm()
	defer db.Close()
	db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&model.Phrase{})

	responsePhrases(w, db, param)
}

func PhrasePOST(w http.ResponseWriter, r *http.Request) {

	// リクエストから値を受けとる
	values, err := util.ParseBody(&r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	phrase := &model.Phrase{
		UserID: values.Get("userid"),
		Text:   values.Get("text"),
		Author: values.Get("author"),
	}

	// Insert
	db := dbaccess.ConnectGorm()
	defer db.Close()
	db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&model.Phrase{})
	db.Create(phrase)

	param := &dbaccess.ParamPhrase{
		UserID: values.Get("userid"),
	}
	responsePhrases(w, db, param)
}

func PhraseDELETE(w http.ResponseWriter, r *http.Request) {

	// リクエストから値を受けとる
	values, err := util.ParseBody(&r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// Delete
	db := dbaccess.ConnectGorm()
	defer db.Close()
	db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&model.Phrase{})

	phrase := &model.Phrase{}
	db.Where("id = ?", values.Get("id")).Where("user_id = ?", values.Get("userid")).Find(phrase)
	if phrase.ID == 0 {
		fmt.Fprint(w, "ng")
		return
	}
	db.Delete(phrase)

	param := &dbaccess.ParamPhrase{
		UserID: values.Get("userid"),
	}
	responsePhrases(w, db, param)
}

func PhrasePATCH(w http.ResponseWriter, r *http.Request) {

	// リクエストから値を受けとる
	values, err := util.ParseBody(&r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// TODO:引数チェック

	// Update
	db := dbaccess.ConnectGorm()
	defer db.Close()
	db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&model.Phrase{})

	phrase := &model.Phrase{}
	db.Where("id = ?", values.Get("id")).Where("user_id = ?", values.Get("userid")).Find(phrase)
	if phrase.ID == 0 {
		fmt.Fprint(w, "ng")
		return
	}

	phrase.Text = values.Get("text")
	phrase.Author = values.Get("author")
	db.Save(phrase)

	param := &dbaccess.ParamPhrase{
		UserID: values.Get("userid"),
	}
	responsePhrases(w, db, param)
}

func responsePhrases(w http.ResponseWriter, db *gorm.DB, param *dbaccess.ParamPhrase) {
	phrases := dbaccess.QueryPhrases(db, param)
	phrasesJSON, err := jsonwrapper.MarshalString(phrases)
	if err != nil {
		fmt.Fprint(w, err)
	}
	fmt.Fprint(w, phrasesJSON)
}
