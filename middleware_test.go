package eden

import "testing"

func TestUse(t *testing.T) {
	r := New()

	testFunc := func(c *Context) {}

	r.Use(testFunc)

	if len(r.middleware) != 1 {
		t.Errorf("Expected router.middleware to be of lenght 1 instead got: %v", len(r.middleware))
	}
}
