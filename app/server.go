package app

import (
	"sync"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	httpSwagger "github.com/swaggo/http-swagger"
)

var (
	instance *server
	once     sync.Once
)

// Server is struct
type server struct {
	Router  *chi.Mux
	Handler *handler
}

// GetInstance is a method
func GetInstance() *server {
	once.Do(func() {
		instance = &server{
			Router:  chi.NewRouter(),
			Handler: NewHandler(),
		}
	})
	return instance
}

func (s *server) Middleware() {
	r := s.Router
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.CloseNotify)
	r.Use(middleware.Timeout(time.Second * 60))
	r.Use(middleware.SetHeader("Content-Type", "application/json"))
	r.Use(render.SetContentType(render.ContentTypeJSON))
}

func (s *server) Routes() {
	r := s.Router
	h := s.Handler
	r.Get("/swagger/*", httpSwagger.WrapHandler)
	r.Get("/ping", h.Ping)
	r.Get("/env", h.Env)
	r.Get("/header", h.Header)
}

// func (s *Server) Say() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		text := r.FormValue("text")
// 		if text == "" {
// 			http.Error(w, "Bad Request", http.StatusBadRequest)
// 			return
// 		}

// 		s.ghClient.Notify(text)
// 		// err := s.client.Notify(text)

// 		str, err := json.Marshal(map[string]string{"text": text})
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		w.Header().Set("Content-Type", "application/json")
// 		w.Write(str)
// 	}
// }
