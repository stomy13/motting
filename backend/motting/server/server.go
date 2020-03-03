package server

import (
	"net/http"

	"github.com/MasatoTokuse/motting/motting/dbaccess"
	"github.com/go-chi/chi"
)

type Serve interface {
	RunServer(handler http.Handler, port string) error
}
type server struct{}

var Server *server

func NewServer() *server {
	return Server
}

func NewHandler(conargs *dbaccess.ConnectArgs) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("Hello Golang"))
	})

	r.Get("/api/v1/word", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("Hello API"))
	})

	r.Get("/test", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("Hello Test"))
	})

	return r
}

func (*server) RunServer(handler http.Handler, port string) error {
	return http.ListenAndServe(port, handler)
}
