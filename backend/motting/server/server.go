package server

import (
	"net/http"

	"github.com/MasatoTokuse/motting/motting/api"
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

	const urlPushtime = "/api/v1/pushtime"
	r.Get(urlPushtime, api.PushTimeGET)
	r.Patch(urlPushtime, api.PushTimePATCH)

	return r
}

func (*server) RunServer(handler http.Handler, port string) error {
	return http.ListenAndServe(port, handler)
}
