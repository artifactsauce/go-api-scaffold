package main

import (
	"log"
	"net/http"

	"github.com/artifactsauce/go-webapi-scaffold/app"
)

// @title a scaffold of Web API using go-chi/chi
// @version 0.1
// @description a scaffold of Web API using go-chi/chi.
// @license.name MIT
func main() {
	s := app.New()
	s.Middleware()
	s.Routes()
	log.Fatal(http.ListenAndServe(":8080", s.Router))
}
