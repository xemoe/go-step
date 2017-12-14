package main

import (
	"fmt"
	"github.com/xemoe/go-step/example/web-jwt-auth/routes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var a routes.App

//
// Setup
//
func TestMain(m *testing.M) {
	a = routes.App{}
	a.Initialize()
	code := m.Run()
	os.Exit(code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	a.Router.ServeHTTP(w, req)

	return w
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

//
// @testobject GET /simple
// @expected JSON string
//
func TestGetSimple(t *testing.T) {
	req, _ := http.NewRequest("GET", "/simple", nil)
	w := executeRequest(req)
	checkResponseCode(t, http.StatusOK, w.Code)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	expected := `{"message":"Hello xemoe"}`
	actual := string(body)
	if expected != actual {
		t.Errorf("Expected response body %s. Got %s\n", expected, actual)
	}

	printdebug := false
	if printdebug {
		fmt.Println(resp.StatusCode)
		fmt.Println(resp.Header.Get("Content-Type"))
		fmt.Println(string(body))
	}
}
