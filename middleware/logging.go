package middleware

import (
	"log"
	"time"

	"github.com/bwinterton/eden"
)

// LogRequest is a logging middleware which logs requests to the server,
// what they are requesting, and how long the request takes to execute
func LogRequest(c *eden.Context) {
	start := time.Now()

	c.Next()

	duration := time.Since(start)
	log.Printf("%s: %s | Request took %f seconds", c.Request.Method, c.Request.RequestURI, duration.Seconds())
}
