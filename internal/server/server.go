package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	server *http.Server
}

func NewServer(addr string, handler *Handler) *Server {
	return &Server{
		server: &http.Server{
			Addr:    addr,
			Handler: registerRoutes(handler),
		},
	}
}

func (s *Server) Run() error {

	return s.server.ListenAndServe()
}

func registerRoutes(h *Handler) *chi.Mux {
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
