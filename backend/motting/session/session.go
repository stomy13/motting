package session

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewFilesystemStore("./", []byte("secure-key"))

const SessionName = "SID"

func NewSession(w http.ResponseWriter, r *http.Request, userID uint) error {

	session, err := store.New(r, SessionName)
	if err != nil {
		return err
	}

	session.Values["userID"] = userID
	fmt.Println(session.Values["userID"])

	err = session.Save(r, w)
	if err != nil {
		return err
	}

	return nil
}

func GetUserID(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	session, err := store.Get(r, SessionName)
	if err != nil {
		return 0, err
	}

	return session.Values["userID"], nil
}
