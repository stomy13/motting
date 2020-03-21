package session

import (
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
	session.Options.HttpOnly = true
	// session.Options.Secure = true
	// session.Options.SameSite = http.SameSiteStrictMode

	return session.Save(r, w)
}

func DeleteSession(w http.ResponseWriter, r *http.Request) error {

	session, err := store.Get(r, SessionName)
	if err != nil {
		return err
	}

	session.Options.MaxAge = -1
	return session.Save(r, w)
}

func GetUserID(r *http.Request) (interface{}, error) {

	session, err := store.Get(r, SessionName)
	if err != nil {
		return 0, err
	}

	return session.Values["userID"], nil
}
