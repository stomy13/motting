package auth

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MasatoTokuse/motting/motting/dbaccess"
	"github.com/MasatoTokuse/motting/motting/model"
	"github.com/MasatoTokuse/motting/motting/session"
	"github.com/MasatoTokuse/motting/motting/util"
	"golang.org/x/crypto/bcrypt"
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

	// email is not already used?
	var count int
	db.Model(&model.User{}).Where("email = ?", user.Email).Count(&count)
	if count > 0 {
		fmt.Fprint(w, "this email is already used")
		return
	}

	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	// fmt.Println("パスワード: ", user.Password)
	// fmt.Println("ハッシュ化されたパスワード", hash)
	user.Password = string(hash)
	// fmt.Println("コンバート後のパスワード: ", user.Password)

	// Insert
	db.Create(user)

	// Create session
	err = session.NewSession(w, r, user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "ok")
}
