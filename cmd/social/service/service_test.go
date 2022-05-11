package service

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
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
	user1 := `{"name":"Aleksander","age":33,"friends":[]}`
	user2 := `{"name":"Pavel","age":34,"friends":[]}`
	w1 := httptest.NewRecorder()
	r1 := httptest.NewRequest(http.MethodPost, "/create", bytes.NewBufferString(user1))
	s := &Server{}

	s.Create(w1, r1)

	if w1.Code != http.StatusCreated {
		t.Logf("got status %v; expected %v", w1.Code, http.StatusCreated)
		t.Fail()
	}
	f := strings.Contains(w1.Body.String(), "ID:")

	if f != true {
		t.Logf("got body %v; expected %v", w1.Body.String(), f)
		t.Fail()
	}

	idSlice := strings.Split(w1.Body.String(), ":")
	idString := idSlice[1]
	idInt, err := strconv.Atoi(idString)
	if err != nil {
		fmt.Println(err)
	}
	if idInt <= 0 {
		t.Logf("got body %v; expected %v", w1.Body.String(), f)
		t.Fail()
	}
	w2 := httptest.NewRecorder()
	r2 := httptest.NewRequest(http.MethodPost, "/create", bytes.NewBufferString(user2))

	s.Create(w2, r2)

	if w2.Code != http.StatusCreated {
		t.Logf("got status %v; expected %v", w2.Code, http.StatusCreated)
		t.Fail()
	}
	f2 := strings.Contains(w2.Body.String(), "ID:")

	if f2 != true {
		t.Logf("got body %v; expected %v", w2.Body.String(), f)
		t.Fail()
	}

	idSlice2 := strings.Split(w2.Body.String(), ":")
	idString2 := idSlice2[1]
	idInt2, err := strconv.Atoi(idString2)
	if err != nil {
		fmt.Println(err)
	}
	if idInt2 <= 0 {
		t.Logf("got body %v; expected %v", w2.Body.String(), f)
		t.Fail()
	}

	if idInt == idInt2 && idInt <= 0 && idInt2 <= 0 {
		t.Logf("got body %v; expected %v", w2.Body.String(), f)
		t.Fail()
	}
}

func TestServer_Make_friends_Error(t *testing.T) {
	st := `{"source_id":10,"target_id":hh}`
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/make_friends", bytes.NewBufferString(st))

	s := &Server{}

	s.Make_friends(w, r)
	if w.Code != http.StatusInternalServerError {
		t.Fail()
	}
	if w.Body.String() != "invalid character 'h' looking for beginning of value" {
		t.Fail()
	}

}

func TestServer_Make_friends_Ok(t *testing.T) {
	st := `{"source_id":10,"target_id":14}`
	st1 := "Aleksander and Pavel friends now"
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/make_friends", bytes.NewBufferString(st))

	s := &Server{}

	s.Make_friends(w, r)

	if w.Code != http.StatusOK {
		t.Logf("got status %v; expected %v", w.Code, http.StatusOK)
		t.Fail()
	}

	if st1 != w.Body.String() {
		t.Fail()
	}

}

func TestServer_Delet_Error(t *testing.T) {
	st := `{"target_id":dd}`
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodDelete, "/delete", bytes.NewBufferString(st))

	s := &Server{}

	s.Delet(w, r)
	if w.Code != http.StatusInternalServerError {
		t.Fail()
	}
	if w.Body.String() != "invalid character 'd' looking for beginning of value" {
		t.Fail()
	}

}

func TestServer_Delet_OK(t *testing.T) {
	st := `{"target_id":5}`
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodDelete, "/delete", bytes.NewBufferString(st))

	s := &Server{}

	s.Delet(w, r)
	t.Log(w.Body.String())
	t.Log(w.Code)
	if w.Code != http.StatusOK {
		t.Fail()
	}
	if w.Body.String() != "Aleksander Delet" {
		t.Fail()
	}

}
