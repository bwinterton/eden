package eden

import "encoding/json"

// WriteJSON writes the given code and data to the client
func (c *Context) WriteJSON(code int, data interface{}) {
	b, err := json.Marshal(data)
	if err != nil {
		c.Response.WriteHeader(500)
		return
	}
	c.Response.WriteHeader(code)
	c.Response.Write(b)
}
