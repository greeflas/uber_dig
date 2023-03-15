package main

import (
	"go.uber.org/dig"
	"log"
)

func main() {
	c := dig.New()

	err := c.Invoke(func() {
		log.Println("Hello, World!")
	})
	if err != nil {
		log.Fatal(err)
	}
}
