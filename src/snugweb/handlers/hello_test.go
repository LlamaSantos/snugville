package handlers_test

import "testing"
import "net/http/httptest"

import "snugweb/handlers"
import "github.com/stretchr/testify/assert"

func TestHello(t *testing.T) {

	// Create our Handler we want to test
	handler := handlers.NewHello()

	// Create our mock request to serve.
	r := httptest.NewRequest("GET", "/", nil)

	// Create the response
	rw := httptest.NewRecorder()

	// Issue the request
	handler.ServeHTTP(rw, r)

	// Assertions for all the things
	assert.Equal(t, rw.Code, 200, "Code must be a 200")
}
