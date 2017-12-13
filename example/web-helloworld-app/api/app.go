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
}

//
// Curl example
// 	`curl -d '{"username":"xemoe"}' \
//		-H "Content-Type: application/json" \
//		-X GET http://localhost:3000/hi`
//
func (a *App) getHelloWorld(w http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)
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
