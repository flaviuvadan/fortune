package main

import (
	"log"
	"net/http"
	"os/exec"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

const port = "8080"

// server holds the associated server commands and serves endpoints
type server struct {
	router *chi.Mux
}

// NewServer creates and returns a new fortune server
func NewServer() *server {
	return &server{router: chi.NewRouter()}
}

// setupRoutes creates and sets up the routes available on this server
func (s *server) setupRoutes() {
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.RedirectSlashes)
	s.router.HandleFunc("/", s.root())
	s.router.HandleFunc("/cookie", s.cookie())
	s.router.HandleFunc("/offensive-cookie", s.offensiveCookie())
}

// root returns some server docs
func (s *server) root() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello! Hit me at:\n" +
			"\t/cookie - to obtain a fortune cookie\n" +
			"\t/offensive-cookie - to obtain a potentially offensive fortune cookie"))
	}
}

// cookie returns a Unix fortune cookie
func (s *server) cookie() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cmd := exec.Command("fortune")
		out, err := cmd.Output()
		if err != nil {
			log.Printf("Error getting cookie: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed to obtain fortune cookie."))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(out)
	}
}

// offensiveCookie returns a potentially offensive fortune cookie
func (s *server) offensiveCookie() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cmd := exec.Command("fortune", "-o")
		out, err := cmd.Output()
		if err != nil {
			log.Printf("Error getting offensive cookie: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed to obtain a potentially offensive fortune cookie."))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(out)
	}
}

func main() {
	s := NewServer()
	s.setupRoutes()
	log.Printf("Starting server on port: %v\n", port)
	http.ListenAndServe(":"+port, s.router)
}
