package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"gorm.io/gorm"
)

type Server interface {
	Run() error
}
type server struct {
	port    string
	handler http.Handler
}

func NewServer(port string, connection *gorm.DB) *server {
	return &server{port: port, handler: NewHandler(connection)}
}

func NewHandler(connection *gorm.DB) *chi.Mux {
	compornents := NewCompornents(connection)

	router := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	router.Use(cors.Handler)

	router.Post("/witticism", func(response http.ResponseWriter, request *http.Request) {
		compornents.WitticismContoroller.AddWitticism(response, request)
	})

	return router
}

func (server *server) Run() error {
	return http.ListenAndServe(server.port, server.handler)
}
