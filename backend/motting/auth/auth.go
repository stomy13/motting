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
	db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&model.User{})

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

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	// リクエストから値を受けとる
	values, err := util.ParseBody(&r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// setup DB
	db := dbaccess.ConnectGorm()
	defer db.Close()

	// is email exists?
	var user model.User
	db.Model(&model.User{}).Where("email = ?", values.Get("email")).Find(&user)
	if user.ID == 0 {
		fmt.Fprint(w, "this email is not exists")
		return
	}

	// Compare
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(values.Get("password")))
	if err != nil {
		fmt.Fprint(w, "Incorrect email or password")
		return
	}

	// Create session
	err = session.NewSession(w, r, user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "ok")

}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	err := session.DeleteSession(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "ok")

}
