package handlers

import (
	"fmt"
	"net/http"
)

// NewHello returns a new hello
func NewHello() http.Handler {
	return new(Hello)
}

// Hello Service
type Hello struct {
	calls int
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.calls++
	s := fmt.Sprintf("Hello World %v times\n", h.calls)
	rw.Write([]byte(s))
}
