package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/artifactsauce/go-webapi-scaffold/app"
	httpSwagger "github.com/swaggo/http-swagger"
)

// Server is a struct
type Server struct {
	router *chi.Mux
}

// New is a method
func New() *Server {
	return &Server{
		router: chi.NewRouter(),
	}
}

// Router is a method
func (s *Server) Router() {
	r := s.router
	r.Get("/swagger/*", httpSwagger.WrapHandler)
	r.Get("/ping", handler.Ping)
	r.Get("/env", handler.Env)
	r.Get("/header", handler.Header)
}

// Middleware is a method
func (s *Server) Middleware() {
	r := s.router
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.CloseNotify)
	r.Use(middleware.Timeout(time.Second * 60))
}

// @title a scaffold of Web API using go-chi/chi
// @version 0.1
// @description a scaffold of Web API using go-chi/chi.
// @license.name MIT
func main() {
	s := New()
	s.Middleware()
	s.Router()
	log.Fatal(http.ListenAndServe(":8080", s.router))
}
