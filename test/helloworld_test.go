package test

import (
	"io"
	"net/http"
	"testing"
)

func Test_HelloWorld(t *testing.T) {
	expected := "Hello, world!"
	resp, err := http.Get("http://localhost:8080")

	if err != nil {
		t.Fatalf("Failed to GET: %s", err)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		t.Fatalf("Failed to read body: %s", err)
	}

	bodyString := string(body)
	if bodyString != expected {
		t.Fatalf("Unexpected response\nExpected: %s\nActual: %s", expected, bodyString)
	}
}
