package main

import (
	"context"
	"github.com/greeflas/uber_dig/app"
	"log"
)

func main() {
	ctx := context.Background()

	if err := app.Run(ctx, app.EnvProd); err != nil {
		log.Fatal(err)
	}
}
