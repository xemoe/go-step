//
// app.go
//
package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
)

type App struct {
	Router *mux.Router
}

type UserJSON struct {
	Username string `json:"username"`
}

//
// Public
//
func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

//
// Privates
//
func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/hi", a.getHelloWorld).Methods("GET")
	a.Router.HandleFunc("/simple", a.getSimple).Methods("GET")
}

//
// Curl example
// 	`curl -d '{"username":"xemoe"}' \
//		-H "Content-Type: application/json" \
//		-X GET http://localhost:3000/hi`
//
func (a *App) getHelloWorld(w http.ResponseWriter, req *http.Request) {

	/**
	const jstream = `{"Username": "foo"}`
	decoder := json.NewDecoder(strings.NewReader(jstream))

	defer req.Body.Close()
	body, _ := ioutil.ReadAll(req.Body)
	bodylen := len(body)

	if req.Body != nil && bodylen > 0 {
	}
	**/

	decoder := json.NewDecoder(req.Body)
	jsondata := UserJSON{}
	err := decoder.Decode(&jsondata)

	helloto := ""
	//
	// Fallback
	//
	if err != nil || jsondata.Username == "" {
		helloto = "world (fallback)"
	} else {
		helloto = jsondata.Username
	}

	respm := fmt.Sprintf("Hello %s", helloto)
	respj := map[string]string{}
	respj["message"] = respm

	respondWithJSON(w, http.StatusOK, respj)
}

//
// Curl example
//  `curl http://localhost:3000/simple`
//
func (a *App) getSimple(w http.ResponseWriter, req *http.Request) {

	const jstream = `{"Username": "xemoe"}`

	if req.Body != nil {
		fmt.Println(req.Body)
	}

	decoder := json.NewDecoder(strings.NewReader(jstream))
	jsondata := UserJSON{}
	err := decoder.Decode(&jsondata)

	helloto := ""
	//
	// Fallback
	//
	if err != nil || jsondata.Username == "" {
		helloto = "world"
	} else {
		helloto = jsondata.Username
	}

	respm := fmt.Sprintf("Hello %s", helloto)
	respj := map[string]string{}
	respj["message"] = respm

	respondWithJSON(w, http.StatusOK, respj)
}

//
// Responses
//
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
