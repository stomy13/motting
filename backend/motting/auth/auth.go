package auth

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MasatoTokuse/motting/motting/dbaccess"
	"github.com/MasatoTokuse/motting/motting/model"
	"github.com/MasatoTokuse/motting/motting/util"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {

	// リクエストから値を受けとる
	values, err := util.ParseBody(&r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	user := &model.User{
		Email:    values.Get("email"),
		Password: values.Get("password"),
	}

	err = user.Validate()
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	// setup DB
	db := dbaccess.ConnectGorm()
	defer db.Close()
	db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&model.User{})

	// email is not already used?
	var count int
	db.Model(&model.User{}).Where("email = ?", user.Email).Count(&count)
	if count > 0 {
		fmt.Fprint(w, "this email is already used")
		return
	}

	// Hash password
	// Insert
	db.Create(user)

	// create session ID
	// return session ID
	fmt.Fprint(w, "ok")
}
