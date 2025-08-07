package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
	port   string
}

func New(port string) *Server {
	s := &Server{
		router: mux.NewRouter(),
		port:   port,
	}
	return s
}

// AddRoutes allows the specific service (menu, order, etc.)
// to register its unique routes.
func (s *Server) AddRoutes(registerRoutes func(r *mux.Router)) {
	registerRoutes(s.router)
}

func (s *Server) Start() {
	fmt.Printf("Starting server on port %s...\n", s.port)
	if err := http.ListenAndServe(s.port, s.router); err != nil {
		log.Fatalf("Could not start server on port %s: %s", s.port, err)
	}
}
