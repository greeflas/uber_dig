package server

import (
	"go.uber.org/dig"
	"net/http"
)

// Route is a http.Handler that knows the mux pattern
// under which it will be registered.
type Route interface {
	http.Handler

	// Pattern reports the path at which this is registered.
	Pattern() string
}

type MuxParams struct {
	dig.In

	Route1 Route `name:"echo"`
	Route2 Route `name:"hello"`
}

func NewServerMux(p MuxParams) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle(p.Route1.Pattern(), p.Route1)
	mux.Handle(p.Route2.Pattern(), p.Route2)

	return mux
}
