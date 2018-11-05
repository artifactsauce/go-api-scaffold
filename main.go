package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/artifactsauce/go-webapi-scaffold/app"
)

var port = flag.Int("port", 8001, "Listening port")

// @title a scaffold of Web API using go-chi/chi
// @version 0.1
// @description a scaffold of Web API using go-chi/chi.
// @license.name MIT
func main() {
	flag.Parse()
	s := app.New()
	s.Middleware()
	s.Routes()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), s.Router))
}
