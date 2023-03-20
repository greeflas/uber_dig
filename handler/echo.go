package handler

import (
	"io"
	"net/http"
)

type EchoHandler struct{}

func NewEchoHandler() *EchoHandler {
	return &EchoHandler{}
}
func (h *EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := io.Copy(w, r.Body); err != nil {
		panic(err)
	}
}

func (h *EchoHandler) Pattern() string {
	return "/echo"
}
