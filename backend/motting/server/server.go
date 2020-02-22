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

type Serve interface {
	RunServer(port string, conargs *ConnectArgs) error
}
type server struct{}

var Server *server

func NewServer() *server {
	return Server
}

func (*server) RunServer(port string, conargs *ConnectArgs) error {
	r := chi.NewRouter()
	r.Get("/api/v1/word", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("Hello Golang"))
	})

	return http.ListenAndServe(port, r)
}
