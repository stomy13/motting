package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/MasatoTokuse/motting/motting/dbaccess"
	"github.com/MasatoTokuse/motting/motting/model"
	"github.com/MasatoTokuse/motting/motting/util"
)

func PushTimeGET(w http.ResponseWriter, r *http.Request) {

	// リクエストから値を受けとる
	values, err := util.ParseBody(&r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	db := dbaccess.ConnectGorm()
	defer db.Close()
	db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&model.PushTime{})
	pt := &model.PushTime{}
	db.Where("user_id = ?", values.Get("userid")).First(pt)

	// At first request, insert initial record
	if pt.ID == 0 {
		initialPushTime := &model.PushTime{
			UserID: values.Get("userid"),
			PushAt: "10:00",
		}
		db.Create(initialPushTime)
		pt = initialPushTime
	}

	pushTimeJSON, _ := json.Marshal(pt)
	fmt.Fprint(w, string(pushTimeJSON))
}

func PushTimePATCH(w http.ResponseWriter, r *http.Request) {

	// リクエストから値を受けとる
	values, err := util.ParseBody(&r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// TODO:引数チェック

	db := dbaccess.ConnectGorm()
	defer db.Close()
	db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&model.PushTime{})

	pt := &model.PushTime{}
	db.Where("user_id = ?", values.Get("userid")).Find(pt)
	if pt.ID == 0 {
		fmt.Fprint(w, "ng")
		return
	}

	pt.PushAt = values.Get("pushAt")
	db.Save(pt)

	pushTimeJSON, _ := json.Marshal(pt)
	fmt.Fprint(w, string(pushTimeJSON))
}
