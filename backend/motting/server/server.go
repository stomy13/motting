package server

import (
	"net/http"

	"github.com/MasatoTokuse/motting/motting/api"
	"github.com/MasatoTokuse/motting/motting/dbaccess"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
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

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

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
