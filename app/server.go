package app

import (
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

type server struct {
	Router  *chi.Mux
	Handler *handler
}

// New is a method
func New() *server {
	return &server{
		Router:  chi.NewRouter(),
		Handler: NewHandler(),
	}
}

func (s *server) Routes() {
	r := s.Router
	h := s.Handler
	r.Get("/swagger/*", httpSwagger.WrapHandler)
	r.Get("/ping", h.Ping)
	r.Get("/env", h.Env)
	r.Get("/header", h.Header)
}

func (s *server) Middleware() {
	r := s.Router
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.CloseNotify)
	r.Use(middleware.Timeout(time.Second * 60))
}
