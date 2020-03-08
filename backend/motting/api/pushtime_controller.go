package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MasatoTokuse/motting/motting/dbaccess"
	"github.com/MasatoTokuse/motting/motting/model"
)

func PushTimeGET(w http.ResponseWriter, r *http.Request) {

	db := dbaccess.ConnectGorm()
	defer db.Close()
	db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&model.PushTime{})
	pt := &model.PushTime{}
	db.First(pt)

	// At first request, insert initial record
	if pt.ID == 0 {
		initialPushTime := &model.PushTime{
			UserID: "whitebox",
			PushAt: "10:00",
		}
		db.Create(initialPushTime)
		pt = initialPushTime
	}

	pushTimeJSON, _ := json.Marshal(pt)
	fmt.Println(*pt)

	fmt.Fprint(w, string(pushTimeJSON))
}
