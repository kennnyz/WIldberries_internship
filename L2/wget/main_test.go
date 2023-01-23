package main

import (
	"testing"
)

func TestWget(t *testing.T) {
	err := downloadWebsite("http://example.com")
	if err != nil {
		t.Error("Error:", err)
	}
	err = downloadWebsite("notfound.com")
	if err != nil {
		t.Error("Error:", err)
	}
}
