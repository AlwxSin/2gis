package rest

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"applicationDesignTest/internal/models"
)

type Server struct {
	mux *chi.Mux

	DB models.DB
}

func (s *Server) ListenAndServe(port int) error {
	httpServer := http.Server{
		Addr:              fmt.Sprintf("0.0.0.0:%d", port),
		Handler:           s.mux,
		ReadHeaderTimeout: time.Second * 5,
	}

	return httpServer.ListenAndServe()
}

type ServerOptions struct {
	DB models.DB
}

func NewServer(opts *ServerOptions) *Server {
	s := &Server{
		mux: chi.NewRouter(),
		DB:  opts.DB,
	}

	s.mux.Use(middleware.Logger)

	s.mux.Post("/orders", s.createOrder)

	return s
}
