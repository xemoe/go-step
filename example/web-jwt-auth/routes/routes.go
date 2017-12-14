//
// routes.go
//

package routes

import (
	"github.com/codegangsta/negroni"
	"github.com/xemoe/go-step/example/web-jwt-auth/controllers"
)

//
// Routes register
//
func (a *App) initializeRoutes() {

	//
	// @routes GET /hi
	//
	a.Router.PathPrefix("/hi").Handler(negroni.New(
		negroni.HandlerFunc(controllers.GetHi),
	)).Methods("GET")

	//
	// @routes GET /simple
	//
	a.Router.PathPrefix("/simple").Handler(negroni.New(
		// negroni.HandlerFunc(authentication.RequireTokenAuthentication),
		negroni.HandlerFunc(controllers.GetSimple),
	)).Methods("GET")

}
