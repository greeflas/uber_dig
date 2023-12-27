package app

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

type Env int

const (
	EnvProd Env = iota
	EnvTest
)

func Run(ctx context.Context, env Env) error {
	var c *dig.Container

	switch env {
	case EnvProd:
		c = dig.New()
	case EnvTest:
		c = dig.New(dig.DryRun(true))
	}

	if err := provideDependencies(c); err != nil {
		return err
	}

	if err := invoke(ctx, c); err != nil {
		return err
	}

	return nil
}

func provideDependencies(c *dig.Container) error {
	err := c.Provide(func() (net.Listener, error) {
		return server.NewTCPListener(":8080")
	})
	if err != nil {
		return err
	}

	if err := c.Provide(handler.NewEchoHandler); err != nil {
		return err
	}

	if err := c.Provide(handler.NewHelloHandler); err != nil {
		return err
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
	if err != nil {
		return err
	}

	if err := c.Provide(server.NewServerMux); err != nil {
		return err
	}

	if err := c.Provide(server.NewHTTPServer); err != nil {
		return err
	}

	return nil
}

func invoke(ctx context.Context, c *dig.Container) error {
	err := c.Invoke(func(srv *server.HTTPServer) {
		log.Printf("Starting HTTP server %s", srv.Addr())
		srv.Start()
	})
	if err != nil {
		return err
	}

	err = c.Invoke(func(srv *server.HTTPServer) error {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, os.Interrupt)

		sig := <-sigs
		log.Printf("Got signal '%s' - shutting down HTTP server...", sig)

		return srv.Shutdown(ctx)
	})
	if err != nil {
		return err
	}

	return nil
}
