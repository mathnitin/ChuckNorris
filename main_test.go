package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
  Test the basic functionality.
  The status code should be 200.
*/
func TestRouter(t *testing.T) {
	// Instantiate the router
	r := newRouter()

	// Create a new server using the "httptest" libraries `NewServer` method
	// Documentation : https://golang.org/pkg/net/http/httptest/#NewServer
	mockServer := httptest.NewServer(r)

	// The mock server we created runs a server and exposes its location in the
	// URL attribute
	// We make a GET request to the "/" route we defined in the router
	resp, err := http.Get(mockServer.URL + "/")

	// Handle any unexpected error
	if err != nil {
		t.Fatal(err)
	}

	// We want our status to be 200 (ok)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be ok, got %d", resp.StatusCode)
	}
}

/*
  Test incorrect operation.
  The status code should be 405.
*/
func TestRouterForIncorrectOperation(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)

	resp, err := http.Post(mockServer.URL+"/", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	// We want our status to be 405
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Status should be 404, got %d", resp.StatusCode)
	}
}

/*
  Test incorrect URL.
  The status code should be 501.
*/
func TestRouterForIncorrectURL(t *testing.T) {
	// Instantiate the router
	r := newRouter()

	// Create a new server using the "httptest" libraries `NewServer` method
	// Documentation : https://golang.org/pkg/net/http/httptest/#NewServer
	mockServer := httptest.NewServer(r)

	// The mock server we created runs a server and exposes its location in the
	// URL attribute
	// We make a GET request to the "/" route we defined in the router
	resp, err := http.Get(mockServer.URL + "/foo")

	// Handle any unexpected error
	if err != nil {
		t.Fatal(err)
	}

	// We want our status to be 501
	if resp.StatusCode != http.StatusNotImplemented {
		t.Errorf("Status should be 404, got %d", resp.StatusCode)
	}
}
