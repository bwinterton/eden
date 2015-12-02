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
	log.Printf("%s: %s | Status: %d | Request took %f miliseconds", c.Request.Method, c.Request.RequestURI, c.Status, duration.Seconds()*100)
}
