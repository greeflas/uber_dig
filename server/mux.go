package server

import (
	"github.com/greeflas/uber_dig/handler"
	"net/http"
)

func NewServerMux(echo *handler.EchoHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/echo", echo)

	return mux
}
