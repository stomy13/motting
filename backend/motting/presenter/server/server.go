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
	// TODO: コントローラーの生成とハンドラーマッピング
	return &server{port: port, handler: NewHandler()}
}

func NewHandler() *chi.Mux {
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

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Motting!"))
	})

	return router
}

func (server *server) Run() error {
	return http.ListenAndServe(server.port, server.handler)
}
