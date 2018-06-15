package app

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

type handler struct {
}

// NewHandler is a method
func NewHandler() *handler {
	return &handler{}
}

// Ping get healtch check
// @Summary get healtch check
// @Description get health check
// @Accept
// @Produce application/json
// @Router /ping [get]
func (h *handler) Ping(w http.ResponseWriter, r *http.Request) {
	var m = map[string]bool{"pong": true}

	str, err := json.Marshal(m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(str)
}

// Env show environment variables
// @Summary show environment variables
// @Description show environment variables
// @Accept
// @Produce application/json
// @Router /env [get]
func (h *handler) Env(w http.ResponseWriter, r *http.Request) {
	var env = make(map[string]string)
	for _, r := range os.Environ() {
		var a = strings.Split(r, "=")
		env[a[0]] = a[1]
	}

	str, err := json.Marshal(env)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(str)
}

// Header show request header from a client
// @Summary show request header from a client
// @Description show request header from a client
// @Accept
// @Produce application/json
// @Router /header [get]
func (h *handler) Header(w http.ResponseWriter, r *http.Request) {
	str, err := json.Marshal(r.Header)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(str)
}
