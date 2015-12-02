package eden

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Context is the main struct that is used to pass information between
// different levels of the middleware
type Context struct {
	Response   http.ResponseWriter
	Request    *http.Request
	Params     httprouter.Params
	Data       map[string]interface{}
	Status     int
	middleware []Middleware
	index      int // used to keep track of which level of the middleware we are currently on.
}

// NewContext populates a context struct with the given data and returns it
func NewContext(w http.ResponseWriter, r *http.Request, p httprouter.Params, mid []Middleware) *Context {
	c := Context{}
	c.Response = w
	c.Request = r
	c.Params = p
	c.Data = nil
	c.middleware = mid
	c.index = -1

	return &c
}

// Next allows the context to continue on to the next level of the middleware.
// Note that each level of the middleware can call c.Next() whenever they would
// like and then the middleware will go on to finish the chain before completing
// the code past the c.Next() call
func (c *Context) Next() {
	for c.index < len(c.middleware)-1 {
		c.index++                // increment right from the get go to get the next level
		c.middleware[c.index](c) // call the next level of middleware
	}
}

/////////////////////
// Error Handling: //
/////////////////////

// Abort stops the context from executing any more of the middleware layers
// however, the already called middleware will continue to execute.
func (c *Context) abort() {
	c.index = len(c.middleware) // if c.index == len(c.middleware) then the for loop in next will not execute further
}

// Fail aborts the context middleware execution and then writes the given
// code and message to the response.
func (c *Context) Fail(code int, m string) {
	c.abort()

	e := DefaultResponse{"ERROR", m}
	c.Respond(code, e)
}
