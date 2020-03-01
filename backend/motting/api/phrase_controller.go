package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/MasatoTokuse/motting/motting/dbaccess"
)

// TODO:クエリで検索できるようにする
func PhraseGET(w http.ResponseWriter, r *http.Request) {

	db := dbaccess.ConnectGorm()
	defer db.Close()
	db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&Phrase{})
	phrases := &[]Phrase{}
	db.Find(phrases)

	phrasesJSON, _ := json.Marshal(phrases)
	fmt.Println(*phrases)

	fmt.Fprint(w, string(phrasesJSON))
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
	phrase := &Phrase{
		UserID: values.Get("userid"),
		Text:   values.Get("text"),
		Author: values.Get("author"),
	}

	// Insert
	db := dbaccess.ConnectGorm()
	defer db.Close()
	db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&Phrase{})
	db.Create(phrase)

	fmt.Fprint(w, "ok")
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
	phrase := &Phrase{}
	phrase.ID = uint(id)

	// Delete
	db := dbaccess.ConnectGorm()
	defer db.Close()
	db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&Phrase{})
	db.Delete(phrase)

	fmt.Fprint(w, "ok")
}
