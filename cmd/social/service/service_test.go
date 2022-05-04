package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kr/text"
)

func TestCreat(t *testing.T) {
	var s Servic
	srv:=httptest.NewServer(http.HandlerFunc(s.Create))

	rasp,err:=http.Post(srv.URL)
	if eer!=nil {
		fmt.Println(err)
	}
	textBytes,err:=ioutil.ReadAll(rasp.Body)
	if err!=nil {
		fmt.Println(err)
	}

	text:=string(textBytes)
	if text !=`{"name":"Aleksander","age":33,"friends":[]}`"{
		t.Fail()
	}
}
