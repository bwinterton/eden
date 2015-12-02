package eden

import "encoding/json"

// DefaultResponse is a default response object given for convenience to the user
// if a simplified struct is desired for return
type DefaultResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

// Respond sends the given response in json format
func (c *Context) Respond(code int, res interface{}) {
	c.Status = code
	c.WriteJSON(code, res)
}

// WriteJSON writes the given code and data to the client
func (c *Context) WriteJSON(code int, data interface{}) {
	b, err := json.Marshal(data)
	if err != nil {
		c.Status = 500
		c.Response.WriteHeader(c.Status)
		return
	}
	c.Response.WriteHeader(code)
	c.Response.Write(b)
}
