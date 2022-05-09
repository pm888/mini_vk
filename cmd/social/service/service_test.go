package service

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_Create_Error(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/create", bytes.NewBufferString(`asd`))
	s := &Server{}

	s.Create(w, r)

	if w.Code != http.StatusInternalServerError {
		t.Fail()
	}
	if w.Body.String() != "invalid character 'a' looking for beginning of value" {
		t.Fail()
	}
}

func TestServer_Create_Ok(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/create", bytes.NewBufferString(`{"name": "bunin", "age": 42}`))
	s := &Server{}

	s.Create(w, r)

	if w.Code != http.StatusInternalServerError {
		t.Fail()
	}
	if w.Body.String() != "invalid character 'a' looking for beginning of value" {
		t.Fail()
	}
}
