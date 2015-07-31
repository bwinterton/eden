package eden

import "testing"

func TestRegister(t *testing.T) {
	r := New()

	testFunc := func(c *Context) {}

	r.Register("GET", "/", testFunc)

	handlerFunc, params, redirect := r.router.Lookup("GET", "/")

	if handlerFunc == nil {
		t.Error("Expected the handler to be a function, instead got nil")
	}

	if len(params) != 0 {
		t.Errorf("Expected params length to be 0 instead got: %v", len(params))
	}

	if redirect != false {
		t.Errorf("Expected route redirect to be true, instead got: %v", redirect)
	}

}

func TestGET(t *testing.T) {
	method := "GET"
	path := "/"
	testFunc := func(c *Context) {}

	r := New()
	r.GET(path, testFunc)
	handler, _, _ := r.router.Lookup(method, path)

	if handler == nil {
		t.Error("Path was not properly registered")
	}
}

func TestPOST(t *testing.T) {
	method := "POST"
	path := "/"
	testFunc := func(c *Context) {}

	r := New()
	r.POST(path, testFunc)
	handler, _, _ := r.router.Lookup(method, path)

	if handler == nil {
		t.Error("Path was not properly registered")
	}
}

func TestDELETE(t *testing.T) {
	method := "DELETE"
	path := "/"
	testFunc := func(c *Context) {}

	r := New()
	r.DELETE(path, testFunc)
	handler, _, _ := r.router.Lookup(method, path)

	if handler == nil {
		t.Error("Path was not properly registered")
	}
}

func TestPUT(t *testing.T) {
	method := "PUT"
	path := "/"
	testFunc := func(c *Context) {}

	r := New()
	r.PUT(path, testFunc)
	handler, _, _ := r.router.Lookup(method, path)

	if handler == nil {
		t.Error("Path was not properly registered")
	}
}
