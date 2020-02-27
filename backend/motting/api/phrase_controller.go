package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/MasatoTokuse/motting/motting/dbaccess"
)

// TODO:デフォは全リスト、クエリで検索できるようにする
func PhraseGET(w http.ResponseWriter, r *http.Request) {

	db := dbaccess.ConnectGorm()
	defer db.Close()
	db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&Phrase{})
	phrase := &Phrase{}
	db.First(phrase)

	phraseJSON, _ := json.Marshal(phrase)
	fmt.Println(*phrase)

	fmt.Fprint(w, string(phraseJSON))
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
