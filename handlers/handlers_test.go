package handlers

import (
	"io"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHome(t *testing.T) {
	// Get the routes
	routes := getRoutes()

	// Create a new server
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	// Get a page
	resp, err := ts.Client().Get(ts.URL)
	if err != nil {
		t.Log(err)
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("for home page, expected status 200 but got %d", resp.StatusCode)
	}

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(bodyText), "awesome") {
		adl.TakeScreenShot(ts.URL+"/", "Hometest", 1500, 1000)
	}
}
