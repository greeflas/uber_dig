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

type RouteResult struct {
	dig.Out

	Routes []Route `group:"routes,flatten"`
}

type MuxParams struct {
	dig.In

	Routes []Route `group:"routes"`
}

func NewServerMux(p MuxParams) *http.ServeMux {
	mux := http.NewServeMux()
	for _, route := range p.Routes {
		mux.Handle(route.Pattern(), route)
	}

	return mux
}
