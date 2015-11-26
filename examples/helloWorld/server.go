package main

import (
	"github.com/bwinterton/eden"
	"github.com/bwinterton/eden/middleware"
)

func main() {

	r := eden.New()

	r.Use(middleware.LogRequest)

	r.GET("/hello", func(c *eden.Context) {
		c.Respond(200, eden.Response{Status: "OK", Data: "Hello World!"})
	})

	r.Run(":8080")

}
