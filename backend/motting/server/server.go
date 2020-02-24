package server

import (
	"net/http"

	"github.com/MasatoTokuse/motting/motting/dbaccess"
	"github.com/go-chi/chi"
)

type Serve interface {
	RunServer(port string, conargs *dbaccess.ConnectArgs) error
}
type server struct{}

var Server *server

func NewServer() *server {
	return Server
}

func (*server) RunServer(port string, conargs *dbaccess.ConnectArgs) error {
	r := chi.NewRouter()
	r.Get("/api/v1/word", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("Hello Golang"))
	})

	return http.ListenAndServe(port, r)
}
