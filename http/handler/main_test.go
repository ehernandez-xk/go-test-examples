package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNameHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:9090/name/pepe", nil)
	if err != nil {
		t.Fatal("Could not create request", err.Error())
	}

	rec := httptest.NewRecorder()

	nameHandler(rec, req)
	if rec.Code != http.StatusOK {
		t.Error("exected 200 and got:", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "Hi pepe") {
		t.Error("expected 'Hi pepe' and got:", rec.Body.String())
	}
}
