package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(mainHandler)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != 302 {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestRouter(t *testing.T) {
	r := newRouter()

	mockServer := httptest.NewServer(r)

	resp, err := http.Get(mockServer.URL + "/")

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be ok, got %d", resp.StatusCode)
	}

	// In the next few lines, the response body is read, and converted to a string
	/* defer resp.Body.Close()
	// read the body into a bunch of bytes (b)
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	// convert the bytes to a string
	respString := string(b)
	expected := "Hello World!"

	// We want our response to match the one defined in our handler.
	// If it does happen to be "Hello world!", then it confirms, that the
	// route is correct
	if respString != expected {
		t.Errorf("Response should be %s, got %s", expected, respString)
	} */

}

func TestRouterForNonExistentRoute(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)
	resp, err := http.Post(mockServer.URL+"/regions", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Status should be 405, got %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	respString := string(b)
	expected := ""

	if respString != expected {
		t.Errorf("Response should be %s, got %s", expected, respString)
	}

}

func TestStaticFileServer(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)

	resp, err := http.Get(mockServer.URL + "/a/")

	if err != nil {
		t.Fatal(err)
	}

	// We want our status to be 200 (ok)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be 200, got %d", resp.StatusCode)
	}

	// It isn't wise to test the entire content of the HTML file.
	// Instead, we test that the content-type header is "text/html; charset=utf-8"
	// so that we know that an html file has been served
	contentType := resp.Header.Get("Content-Type")
	expectedContentType := "text/html; charset=utf-8"

	if expectedContentType != contentType {
		t.Errorf("Wrong content type, expected %s, got %s", expectedContentType, contentType)
	}

}
