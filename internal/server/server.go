package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	server *http.Server
}

func NewServer(addr string) *Server {
	return &Server{
		server: &http.Server{
			Addr:    addr,
			Handler: registerRoutes(newHandler()),
		},
	}
}

func (s *Server) Run() error {
	log.Printf("Starting server at %s\n", s.server.Addr)

	return s.server.ListenAndServe()
}

func registerRoutes(h *handler) *chi.Mux {
	r := chi.NewRouter()

	// common middlewares here
	// r.Use(...)

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/status/{id}", h.handleTaskStatus)
		r.Get("/download/{id}", h.handleDownload)
		r.Post("/upload", h.handleUpload)
	})

	return r
}
