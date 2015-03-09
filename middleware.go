package eden

// The Middleware type is any function that a *Context
type Middleware func(*Context)

// Use appends the given middleware to the router's middleware.
func (r *Router) Use(m ...Middleware) {
	r.middleware = append(r.middleware, m...)
}
