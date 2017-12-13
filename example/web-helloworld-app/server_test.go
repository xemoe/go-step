package main

import (
	"github.com/xemoe/go-step/example/web-helloworld-app/api"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var a api.App

//
// Setup
//
func TestMain(m *testing.M) {
	a = api.App{}
	a.Initialize()
	code := m.Run()
	os.Exit(code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestGetSimple(t *testing.T) {
	req, _ := http.NewRequest("GET", "/simple", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
}
