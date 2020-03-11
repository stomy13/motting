package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/MasatoTokuse/motting/motting/dbaccess"
	jsonwrapper "github.com/MasatoTokuse/motting/motting/json"
	"github.com/MasatoTokuse/motting/motting/model"
	"github.com/jinzhu/gorm"
)

// TODO:クエリで検索できるようにする
func PhraseGET(w http.ResponseWriter, r *http.Request) {

	db := dbaccess.ConnectGorm()
	defer db.Close()
	db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&model.Phrase{})

	responsePhrases(w, db)
}

func PhrasePOST(w http.ResponseWriter, r *http.Request) {

	// リクエストから値を受けとる
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	values, err := url.ParseQuery(string(bytes))
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

	responsePhrases(w, db)
}

func PhraseDELETE(w http.ResponseWriter, r *http.Request) {

	// リクエストから値を受けとる
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	values, err := url.ParseQuery(string(bytes))
	if err != nil {
		log.Fatalln(err)
	}
	id, err := strconv.Atoi(values.Get("id"))
	if err != nil {
		log.Fatalln(err)
	}
	phrase := &model.Phrase{}
	phrase.ID = uint(id)

	// Delete
	db := dbaccess.ConnectGorm()
	defer db.Close()
	db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&model.Phrase{})
	db.Delete(phrase)

	responsePhrases(w, db)
}

func PhrasePATCH(w http.ResponseWriter, r *http.Request) {

	// リクエストから値を受けとる
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	values, err := url.ParseQuery(string(bytes))
	if err != nil {
		log.Fatalln(err)
	}

	// TODO:引数チェック

	// Update
	db := dbaccess.ConnectGorm()
	defer db.Close()
	db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&model.Phrase{})

	phrase := &model.Phrase{}
	db.Where("id = ?", values.Get("id")).Find(phrase)
	if phrase.ID == 0 {
		fmt.Fprint(w, "ng")
		return
	}

	phrase.UserID = values.Get("userid")
	phrase.Text = values.Get("text")
	phrase.Author = values.Get("author")
	db.Save(phrase)

	responsePhrases(w, db)
}

func responsePhrases(w http.ResponseWriter, db *gorm.DB) {
	phrases := dbaccess.PhrasesAll(db)
	phrasesJSON, err := jsonwrapper.MarshalString(phrases)
	if err != nil {
		fmt.Fprint(w, err)
	}
	fmt.Fprint(w, phrasesJSON)
}
