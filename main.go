package main

import (
	"log"
	"os"

	"github.com/kataras/iris"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	iris.Get("/", func(c *iris.Context) {
		c.HTML("<b> Iris </b>")
	})

	iris.Listen(port)
}
