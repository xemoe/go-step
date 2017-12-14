//
// simple_controller.go
//

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//
// Curl example
// 	`curl -d '{"username":"xemoe"}' \
//		-H "Content-Type: application/json" \
//		-X GET http://localhost:3000/hi`
//
func GetHi(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {

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

	RespondWithJSON(w, http.StatusOK, respj)
}
