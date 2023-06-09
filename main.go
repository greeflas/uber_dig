package main

import (
	"context"
	"github.com/greeflas/uber_dig/handler"
	"github.com/greeflas/uber_dig/server"
	"go.uber.org/dig"
	"log"
	"net"
	"os"
	"os/signal"
)

func main() {
	ctx := context.Background()
	c := dig.New()

	err := c.Provide(func() (net.Listener, error) {
		return server.NewTCPListener(":8080")
	})
	if err != nil {
		log.Fatal(err)
	}

	err = c.Provide(handler.NewEchoHandler)
	if err != nil {
		log.Fatal(err)
	}

	err = c.Provide(handler.NewHelloHandler)
	if err != nil {
		log.Fatal(err)
	}

	err = c.Provide(func(
		echo *handler.EchoHandler,
		hello *handler.HelloHandler,
	) server.RouteResult {
		return server.RouteResult{
			Routes: []server.Route{
				echo,
				hello,
			},
		}
	})

	err = c.Provide(server.NewServerMux)
	if err != nil {
		log.Fatal(err)
	}

	err = c.Provide(server.NewHTTPServer)
	if err != nil {
		log.Fatal(err)
	}

	err = c.Invoke(func(srv *server.HTTPServer) {
		log.Printf("Starting HTTP server %s", srv.Addr())
		srv.Start()
	})
	if err != nil {
		log.Fatal(err)
	}

	err = c.Invoke(func(srv *server.HTTPServer) error {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, os.Interrupt)

		sig := <-sigs
		log.Printf("Got signal '%s' - shutting down HTTP server...", sig)

		return srv.Shutdown(ctx)
	})
	if err != nil {
		log.Fatal(err)
	}
}
