package handler

import (
	"fmt"
	"io"
	"net/http"
)

type HelloHandler struct{}

func NewHelloHandler() *HelloHandler {
	return &HelloHandler{}
}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)

		return
	}

	if _, err := fmt.Fprintf(w, "Hello, %s\n", body); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)

		return
	}
}

func (h *HelloHandler) Pattern() string {
	return "/hello"
}
