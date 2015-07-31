package eden

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Struct to represent the router
type Router struct {
	// The httprouter Router
	router *httprouter.Router

	// Slice that holds the middleware
	middleware []Middleware
}

// Register registeres a given method, path, and middleware to the httprouter for handling
func (r *Router) Register(method string, path string, mid ...Middleware) {
	// Combine all the middleware passed in with the global middleware
	allMiddleware := append(r.middleware, mid...)

	// The handler which combines all of the middleware together
	handler := func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		c := NewContext(w, r, p, allMiddleware)
		c.Next()
	}

	r.router.Handle(method, path, handler)
}

// GET is an alias for router.Register("GET", path, middleware)
func (r *Router) GET(path string, mid ...Middleware) {
	r.Register("GET", path, mid...)
}

// POST is an alias for router.Register("POST", path, middleware)
func (r *Router) POST(path string, mid ...Middleware) {
	r.Register("POST", path, mid...)
}

// DELETE is an alias for router.Register("DELETE", path, middleware)
func (r *Router) DELETE(path string, mid ...Middleware) {
	r.Register("DELETE", path, mid...)
}

// UPDATE is an alias for router.Register("UPDATE", path, middleware)
func (r *Router) PUT(path string, mid ...Middleware) {
	r.Register("PUT", path, mid...)
}
