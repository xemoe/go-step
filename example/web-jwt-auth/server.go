//
// main.go
//
package main

import (
	"github.com/xemoe/go-step/example/web-jwt-auth/routes"
)

func main() {
	a := routes.App{}
	a.Initialize()
	a.Run(":3000")
}
