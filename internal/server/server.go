package server

import (
	"go-http/internal/handler"
	"net/http"
)

type Server struct {
	mux *http.ServeMux
}

func New() *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.Home)
	return &Server{mux: mux}
}

func (s *Server) Start(addr string) error {
	return http.ListenAndServe(addr, s.mux)
}
