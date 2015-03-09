package eden

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// New returns a new *router to be used to configure the endpoints
func New() *Router {
	r := Router{}
	r.router = httprouter.New()
	r.middleware = nil

	return &r
}

// ServeHTTP calls the httprouter.ServeHTTP and implements router as an http.Handler
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(w, req)
}

// Run runs the server listening on the given address
func (r *Router) Run(address string) error {
	if err := http.ListenAndServe(address, r); err != nil {
		return err
	}
	return nil
}

// RunTLS runs a TLS server listening on the given address with the given cert and key
func (r *Router) RunTLS(address string, cert string, key string) error {
	if err := http.ListenAndServeTLS(address, cert, key, r); err != nil {
		return err
	}
	return nil
}

/*
 * TODO: So now there is a lot to be done here, I have the context pretty much
 * working I think, but there is more to do there. The context also needs to
 * be able to be created from the beginning. There should also proabably be an error
 * handling middleware for logging the issues. As well as for logging
 * response times.
 */
