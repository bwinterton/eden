package eden

import (
	"reflect"
	"testing"

	"github.com/julienschmidt/httprouter"
)

// TestNew tests the New() function
func TestNew(t *testing.T) {
	r := New()

	if reflect.TypeOf(r) != reflect.TypeOf(&Router{}) {
		t.Errorf("Expected router to be of type Router instead got: %v", reflect.TypeOf(r))
	}

	if reflect.TypeOf(r.router) != reflect.TypeOf(&httprouter.Router{}) {
		t.Errorf("router.router type is: %v expected: httprouter.Router", reflect.TypeOf(r.router))
	}

	if r.middleware != nil {
		t.Errorf("Expected router.middleware to be nil instead got: %v", r.middleware)
	}
}
