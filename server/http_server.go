package server

import (
	"context"
	"log"
	"net"
	"net/http"
)

type HTTPServer struct {
	srv      *http.Server
	listener net.Listener
}

func NewHTTPServer(listener net.Listener, mux *http.ServeMux) *HTTPServer {
	return &HTTPServer{
		srv: &http.Server{
			Handler: mux,
		},
		listener: listener,
	}
}

func (s *HTTPServer) Addr() string {
	return s.listener.Addr().String()
}

func (s *HTTPServer) Start() {
	go func() {
		if err := s.srv.Serve(s.listener); err != nil && err != http.ErrServerClosed {
			log.Panic(err)
		}
	}()
}

func (s *HTTPServer) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
