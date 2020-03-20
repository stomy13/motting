package api

import (
	"fmt"
	"net/http"

	"github.com/MasatoTokuse/motting/motting/dbaccess"
	jsonwrapper "github.com/MasatoTokuse/motting/motting/json"
	"github.com/MasatoTokuse/motting/motting/model"
)

func UsersPushToGET(w http.ResponseWriter, r *http.Request) {

	// リクエストから値を受けとる
	values := r.URL.Query()

	db := dbaccess.ConnectGorm()
	defer db.Close()

	pt := []model.PushTime{}
	db.Where("push_at = ?", values.Get("pushAt")).Find(&pt)

	pushTimeJSON, err := jsonwrapper.MarshalString(&pt)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	fmt.Fprint(w, pushTimeJSON)
}
