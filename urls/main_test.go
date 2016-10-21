package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"fmt"
	"time"
)

func TestItExists(t *testing.T) {

	_, err := Concatenator("http://some.url")

	if (err == nil) {
		t.Errorf("Expected something out of the function")
	}
}

func TestItFetchesMultipleUrls(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, client")
	}))
	defer ts.Close()

	expectedOutput := "Hello, client"
	expectedOutput += expectedOutput

	output, err := Concatenator(ts.URL, ts.URL)
	if err != nil {
		t.Errorf("Did not expect an error")
	}

	if output != expectedOutput {
		t.Errorf("Expected '%s', got '%s'", expectedOutput, output)
	}
}

func TestItFetchesAUrl(t *testing.T) {
	expectedOutput := "Hello, client"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, expectedOutput)
	}))
	defer ts.Close()

	output, err := Concatenator(ts.URL)
	if err != nil {
		t.Errorf("Did not expect an error")
	}

	if output != expectedOutput {
		t.Errorf("Expected '%s', got '%s'", expectedOutput, output)
	}
}

func TestItErrorsForHttpFails(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Internet is broken", http.StatusBadRequest)
	}))
	defer ts.Close()

	_, err := Concatenator(ts.URL)
	if err == nil {
		t.Errorf("Expected an error back, since the server did a 500 error")
	}
}

func BenchmarkManyFetches(b *testing.B) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(200 * time.Millisecond)
		fmt.Fprint(w, "AAAAAA")
	}))
	defer ts.Close()

	urls := []string{}
	for i := 0; i < 100; i++ {
		urls = append(urls, ts.URL)
	}

	for i := 0; i < b.N; i++ {
		concatMany(b, urls...)
	}
}

func concatMany(b *testing.B, urls ...string) {
	_, err := Concatenator(urls...)
	if err != nil {

		b.Errorf("Did not expect an error: %s", err)
	}
}

