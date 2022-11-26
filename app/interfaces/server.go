package interfaces

import (
	"api_template/app/configs"

	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	Router *chi.Mux
	db     *sqlx.DB
}

func NewServer() *Server {
	return &Server{
		Router: chi.NewRouter(),
	}
}

func (s *Server) Init() error {
	conn, err := configs.Init()
	if err != nil {
		return fmt.Errorf("failed db init. %s", err)
	}
	s.db = conn
	s.Route()
	return nil
}

func (s *Server) Route() {

	s.Router.Use(cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}).Handler)

	s.Router.Use(middleware.Logger)
	//接続確認
	s.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

}
