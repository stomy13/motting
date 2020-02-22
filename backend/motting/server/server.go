package server

import (
	"net/http"

	"github.com/go-chi/chi"
)

type ConnectArgs struct {
	Address  string
	Port     string
	DBName   string
	User     string
	Password string
}

func RunServer(port string, conargs *ConnectArgs) error {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("Hello Golang"))
	})
	http.ListenAndServe(port, r)

	return nil
}
