//
// hi_controller.go
//

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

//
// Curl example
//  `curl http://localhost:3000/simple`
//
func GetSimple(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {

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

	RespondWithJSON(w, http.StatusOK, respj)
}
