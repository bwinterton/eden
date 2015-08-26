package eden

import "encoding/json"

type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func (c *Context) Respond(code int, res Response) {
	c.WriteJSON(code, res)
}

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
