//
// main.go
//
package main

import (
	"github.com/xemoe/go-step/example/web-helloworld-app/api"
)

func main() {
	a := api.App{}
	a.Initialize()
	a.Run(":3000")
}
